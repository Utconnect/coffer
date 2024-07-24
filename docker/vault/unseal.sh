#!/bin/bash

export VAULT_ADDR='http://0.0.0.0:5200'

STATUS=$(vault status 2>/dev/null)
INIT_OUTPUT_FILE="vault_init_output.txt"

if [ $? -ne 0 ]; then
    echo "Error: Unable to connect to Vault"
    exit 1
fi

INITIALIZED=$(echo "$STATUS" | grep 'Initialized' | awk '{print $2}')
SEALED=$(echo "$STATUS" | grep 'Sealed' | awk '{print $2}')

if [ "$INITIALIZED" = "false" ]; then
    echo "Vault is not initialized. Initializing..."
    vault operator init > "$INIT_OUTPUT_FILE"
    if [ $? -eq 0 ]; then
        echo "Vault has been initialized. Output saved to $INIT_OUTPUT_FILE"
        echo "Please secure this file as it contains sensitive information."
    else
        echo "Error: Failed to initialize Vault"
        exit 1
    fi
fi

sleep .25
if [ "$SEALED" = "true" ] && [ -f "$INIT_OUTPUT_FILE" ]; then
    echo "Vault is sealed. Unsealing using keys from $INIT_OUTPUT_FILE..."
    UNSEAL_KEYS=$(grep "Unseal Key " "$INIT_OUTPUT_FILE" | awk '{print $4}')

    echo "$UNSEAL_KEYS" | head -n 3 | while read -r key; do
        sleep .25
        vault operator unseal "$key" > /dev/null
    done

    echo "Vault has been unsealed."
elif [ "$SEALED" = "false" ]; then
    echo "Vault is already unsealed."
else
    echo "Vault is sealed but no init file found. Unable to unseal automatically."
    exit 1
fi

# Login to Vault using the root token
sleep .25
if [ -f "$INIT_OUTPUT_FILE" ]; then
    echo "Looking up for logged in token."

    if vault token lookup > /dev/null 2>&1; then
        echo "Vault is already logged in."
    else
        echo "Logging in to Vault..."
        ROOT_TOKEN=$(grep "Initial Root Token: " "$INIT_OUTPUT_FILE" | awk '{print $4}')
        vault login "$ROOT_TOKEN"
        echo "Successfully logged in to Vault."
        echo "Enabling KV secrets engine at be/ namespace..."
        sleep 2
    fi

    if vault secrets list | grep '^be/' > /dev/null; then
        echo "KV secrets engine is already enabled at be/ namespace."
    else
        vault secrets enable -path=be/ kv-v2
        if [ $? -eq 0 ]; then
            echo "KV secrets engine enabled successfully at be/ namespace."
        else
            echo "Failed to enable KV secrets engine at be/ namespace."
        fi
    fi

    sleep 1
    echo "Adding default secrets"

    # Iterate through all environment variables
    env | while IFS='=' read -r name value; do
        # Check if the variable name starts with VAULT_KV_
        if [[ $name == VAULT_KV_* ]]; then
            # Remove VAULT_KV_ prefix
            name_without_prefix=${name#VAULT_KV_}

            # Convert to lowercase
            namespace=$(echo "$name_without_prefix" | cut -d "_" -f 1 | tr '[:upper:]' '[:lower:]')
            app=$(echo "$name_without_prefix" | cut -d "_" -f 2 | tr '[:upper:]' '[:lower:]')
            env_name=$(echo "$name_without_prefix" | cut -d "_" -f 3-)

            # Construct the Vault namespace and key
            vault_namespace="${namespace}/${app}"
            vault_key="${env_name}"

            # Add the secret to Vault
            if vault kv patch "${vault_namespace}" "${vault_key}=${value}" > /dev/null 2>&1; then
                echo "Secret ${vault_namespace}/${vault_key} added successfully to Vault."
            elif vault kv put "${vault_namespace}" "${vault_key}=${value}" > /dev/null 2>&1; then
                echo "Secret ${vault_namespace}/${vault_key} added successfully to Vault."
            else
                echo "Failed to add secret ${vault_namespace}/${vault_key} to Vault."
            fi
        fi
    done
else
    echo "Init file not found. Unable to log in automatically."
    exit 1
fi