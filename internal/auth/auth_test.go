package auth

import (
	"net/http"
	"testing"
)

func TestAPIKey_Success(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey secret-token")
	token, err := GetAPIKey(h)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if token != "secret-token" {
		t.Fatalf("expected token 'secret-token', got %q", token)
	}
}

func TestAPIKey_MissingHeader(t *testing.T) {
	h := http.Header{}
	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestAPIKey_EmptyAuthorizationValue(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "")

	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
}

func TestGetAPIKey_WrongScheme(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "WrongScheme abc")

	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "malformed authorization header" {
		t.Fatalf("expected 'malformed authorization header' error, got %v", err)
	}
}

func TestGetAPIKey_MissingToken(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey")

	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatal("expected error, got nil")
	}

	if err.Error() != "malformed authorization header" {
		t.Fatalf("expected 'malformed authorization header' error, got %v", err)
	}
}

