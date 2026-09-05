package auth_test

import (
	"MODULE_PATH/internal/auth"
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	headers := http.Header{}
	bearer := "Bearer TOKEN_STRING"
	headers.Set("Authorization", bearer)
	expected := "TOKEN_STRING"
	result, err := auth.GetBearerToken(headers)
	if err != nil {
		t.Fatalf("Can't get bearer token %v", err)
	}
	if result != expected {
		t.Fatalf("got %q, want %q", result, expected)
	}
}
