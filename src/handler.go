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

	secretPath := fmt.Sprintf("%s/%s", namespace, app)
	url := fmt.Sprintf("%s/v1/%s", vaultAddress, secretPath)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error when creating request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Header.Set("X-Vault-Token", vaultToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error when processing request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("Response status is not OK")
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

	response := ApiResponse{
		Data: secretResp.Data[secretName],
	}

	_ = json.NewEncoder(w).Encode(response)

}
