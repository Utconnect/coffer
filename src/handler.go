package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
)

var vaultAddress = os.Getenv("VAULT_ADDRESS")
var vaultToken = os.Getenv("VAULT_TOKEN")

type SecretResponse struct {
	Data SecretResponseData `json:"data"`
}

type SecretResponseData struct {
	Data map[string]string `json:"data"`
}

type ApiResponse struct {
	Data string `json:"data"`
}

func getSecret(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	app := vars["app"]
	namespace := vars["namespace"]
	secretName := vars["secretName"]

	secretPath := fmt.Sprintf("%s/data/%s", namespace, app)
	url := fmt.Sprintf("%s/v1/%s", vaultAddress, secretPath)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-Vault-Token", vaultToken)

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("Response status is not OK")
		http.Error(w, fmt.Sprintf("Vault returned status code %d", resp.StatusCode), resp.StatusCode)
		return
	}

	body, _ := io.ReadAll(resp.Body)

	var secretResp SecretResponse
	_ = json.Unmarshal(body, &secretResp)

	response := ApiResponse{
		Data: secretResp.Data.Data[secretName],
	}

	_ = json.NewEncoder(w).Encode(response)
}
