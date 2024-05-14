package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/HiroAcocoro/meal-planner-app-server/internal/types"
)

func TestCreateJwt(t *testing.T) {
	secret := []byte("secret")

	token, err := CreateJwt(secret, "test")
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}

	if token == "" {
		t.Error("expected token to be not empty")
	}
}

func TestParseJwt(t *testing.T) {
	secret := []byte("secret")
	mockToken, err := CreateJwt(secret, "test")
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}
	formatMockToken := fmt.Sprintf("Bearer %s", mockToken)

	request, _ := http.NewRequest("GET", "/", nil)
	request.Header.Set("Authorization", formatMockToken)

	token, err := ParseJwt(request)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if token == nil || !token.Valid {
		t.Errorf("Expected valid token, got %v", token)
	}

	// Test with invalid token
	request.Header.Set("Authorization", "invalid_token")
	_, err = ParseJwt(request)
	if err == nil || !strings.Contains(err.Error(), "invalid token") {
		t.Errorf("Expected error with message 'invalid token', got %v", err)
	}
}

func TestGetUserIdByToken(t *testing.T) {
	secret := []byte("secret")
	expectedUserId := "testId"
	mockToken, err := CreateJwt(secret, expectedUserId)
	if err != nil {
		t.Errorf("error creating JWT: %v", err)
	}
	formatMockToken := fmt.Sprintf("Bearer %s", mockToken)

	request, _ := http.NewRequest("GET", "/", nil)
	request.Header.Set("Authorization", formatMockToken)

	token, err := ParseJwt(request)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if token == nil || !token.Valid {
		t.Errorf("Expected valid token, got %v", token)
	}

	userId := GetUserIdByToken(token)
	if userId != expectedUserId {
		t.Errorf("Expected userId '%s', got '%s'", expectedUserId, userId)
	}
}

func TestGetUserIdFromContext(t *testing.T) {
  expectedUserId := "testUserId"
  ctx := context.WithValue(context.Background(), types.UserKey, expectedUserId)

  userId := GetUserIdFromContext(ctx)
  
  if userId != expectedUserId {
    t.Errorf("Expected userId '%s', got '%s'", expectedUserId, userId)
  }
}
