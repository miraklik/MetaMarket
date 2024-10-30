package jwt_test

import (
	"testing"
	"time"

	"internal/internal/domain/models"
	jwt_test "internal/internal/lib/jwt"

	"github.com/dgrijalva/jwt-go"
)

func TestNewToken_Success(t *testing.T) {
	user := models.User{
		ID:    1,
		Email: "test@example.com",
	}

	app := models.App{
		ID:     1,
		Secret: "testsecret",
	}

	duration := time.Hour

	tokenString, err := jwt_test.NewToken(user, app, duration)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if tokenString == "" {
		t.Fatalf("Expected token string, got empty string")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.Secret), nil
	})

	if err != nil {
		t.Fatalf("Error parsing token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		t.Fatalf("Token is not valid or claims conversion failed")
	}

	if claims["uid"] != float64(user.ID) {
		t.Errorf("Expected user ID %d, got %v", user.ID, claims["uid"])
	}

	if claims["email"] != user.Email {
		t.Errorf("Expected email %s, got %s", user.Email, claims["email"])
	}

	if claims["app_id"] != float64(app.ID) {
		t.Errorf("Expected app ID %d, got %v", app.ID, claims["app_id"])
	}
}

func TestNewToken_InvalidSecret(t *testing.T) {

	user := models.User{
		ID:    1,
		Email: "test@example.com",
	}

	app := models.App{
		ID:     1,
		Secret: "",
	}

	duration := time.Hour

	_, err := jwt_test.NewToken(user, app, duration)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}
}

func TestNewToken_ExpiredToken(t *testing.T) {

	user := models.User{
		ID:    1,
		Email: "test@example.com",
	}

	app := models.App{
		ID:     1,
		Secret: "testsecret",
	}

	duration := -time.Hour

	tokenString, err := jwt_test.NewToken(user, app, duration)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if tokenString == "" {
		t.Fatalf("Expected token string, got empty string")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(app.Secret), nil
	})

	if err != nil {
		t.Fatalf("Error parsing token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || token.Valid {
		t.Fatalf("Expected token to be invalid due to expiration")
	}

	if exp, ok := claims["exp"].(float64); ok {
		if exp > float64(time.Now().Unix()) {
			t.Errorf("Expected token to be expired, exp: %v", exp)
		}
	} else {
		t.Error("Exp claim not found in token")
	}
}
