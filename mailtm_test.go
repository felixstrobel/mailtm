package mailtm

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNew(t *testing.T) {
	client := New()

	if client.baseUrl != DefaultBaseURL {
		t.Errorf("expected baseUrl %s, got %s", DefaultBaseURL, client.baseUrl)
	}
	if client.http == nil {
		t.Error("expected http client to be initialized, got nil")
	}
}

func TestWithBaseURL(t *testing.T) {
	customURL := "https://custom.mail.tm"
	client := New(WithBaseURL(customURL))

	if client.baseUrl != customURL {
		t.Errorf("expected baseUrl %s, got %s", customURL, client.baseUrl)
	}
}

func TestWithHttpClient(t *testing.T) {
	mockHttpClient := &http.Client{}
	client := New(WithHttpClient(mockHttpClient))

	if client.http != mockHttpClient {
		t.Error("expected custom HTTP client, got different instance")
	}
}

func TestAuthenticatedRequest_MissingToken(t *testing.T) {
	client := New()

	req, _ := http.NewRequest("GET", "https://example.com", nil)
	err := client.authenticatedRequest(req, nil)

	if err == nil || err.Error() != "missing authentication" {
		t.Errorf("expected missing authentication error, got %v", err)
	}
}

func TestAuthenticatedRequest_Success(t *testing.T) {
	mockResponse := map[string]string{"key": "value"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "Bearer test-token" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))
	client.token = "test-token"

	req, _ := http.NewRequest("GET", server.URL, nil)
	var result map[string]string
	err := client.authenticatedRequest(req, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result["key"] != "value" {
		t.Errorf("expected key 'value', got %s", result["key"])
	}
}

func TestRequest_Success(t *testing.T) {
	mockResponse := map[string]string{"key": "value"}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResponse)
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	req, _ := http.NewRequest("GET", server.URL, nil)
	var result map[string]string
	err := client.request(req, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result["key"] != "value" {
		t.Errorf("expected key 'value', got %s", result["key"])
	}
}

func TestRequest_WithNilResult(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a valid JSON response even though we're not expecting to unmarshal it
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"key": "value"}`))
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	req, _ := http.NewRequest("GET", server.URL, nil)

	// Passing nil as the result, we should not get an error.
	err := client.request(req, nil)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRequest_WithEmptyStructResult(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a valid JSON response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"key": "value"}`))
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	req, _ := http.NewRequest("GET", server.URL, nil)

	// Passing an empty struct as the result
	var result struct{} // Empty struct
	err := client.request(req, &result)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
