package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func mockVaultServer() *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v1/namespace/app" {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		if r.Header.Get("X-Vault-Token") != "test-token" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		response := SecretResponse{
			Data: map[string]string{"secretName": "secretValue"},
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(response)
	})
	return httptest.NewServer(handler)
}

func getHandledRouter(method string, template string, f func(http.ResponseWriter, *http.Request)) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(template, f).Methods(method)
	return router
}

func serveRequest(router *mux.Router, method string, url string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, url, body)

	// Record the response
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func TestGetSecret_VaultError(t *testing.T) {
	// Set environment variables for testing
	vaultToken = "invalid-test-token"

	// Mock Vault server
	vaultServer := mockVaultServer()
	defer vaultServer.Close()

	// Override vaultAddress to point to the mock server
	vaultAddress = vaultServer.URL

	router := getHandledRouter("GET", "/secret/{namespace}/{app}/{secretName}", getSecret)
	rr := serveRequest(router, "GET", "/secret/namespace/app/secretName", nil)

	expected := "Vault returned status code 401"
	if body := strings.TrimSpace(rr.Body.String()); body != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", body, expected)
	}
}

func TestGetSecret_Success(t *testing.T) {
	// Set environment variables for testing
	vaultToken = "test-token"

	// Mock Vault server
	vaultServer := mockVaultServer()
	defer vaultServer.Close()

	// Override vaultAddress to point to the mock server
	vaultAddress = vaultServer.URL

	router := getHandledRouter("GET", "/secret/{namespace}/{app}/{secretName}", getSecret)
	rr := serveRequest(router, "GET", "/secret/namespace/app/secretName", nil)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body is what we expect
	expected := `{"data":"secretValue"}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
