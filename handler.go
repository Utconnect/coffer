package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var vaultAddress = os.Getenv("VAULT_ADDRESS")
var vaultToken = os.Getenv("VAULT_TOKEN")

// SecretResponse represents the structure of the Vault response
type SecretResponse struct {
	Data map[string]interface{} `json:"data"`
}

func getSecret(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting secret...")
	secretName := r.URL.Query().Get("name")
	if secretName == "" {
		http.Error(w, "Missing 'name' query parameter", http.StatusBadRequest)
		return
	}

	secretPath := fmt.Sprintf("secret/data/%s", secretName)
	url := fmt.Sprintf("%s/v1/%s", vaultAddress, secretPath)
	log.Printf(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("X-Vault-Token", vaultToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Vault returned status code %d", resp.StatusCode), resp.StatusCode)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var secretResp SecretResponse
	err = json.Unmarshal(body, &secretResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(secretResp.Data)
}
