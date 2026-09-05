package auth_test

import (
	"testing"
	"time"

	"MODULE_PATH/internal/auth"
	"github.com/google/uuid"
)

func TestValidJWT(t *testing.T) {
	userID := uuid.New()
	secret := "my-secret"

	token, err := auth.MakeJWT(userID, secret, time.Hour)
	if err != nil {
		t.Fatalf("MakeJWT failed: %v", err)
	}
	gotID, err := auth.ValidateJWT(token, secret)
	if err != nil {
		t.Fatalf("ValidateJWT failed: %v", err)
	}

	if gotID != userID {
		t.Errorf("got %v, want %v", gotID, userID)
	}
}
