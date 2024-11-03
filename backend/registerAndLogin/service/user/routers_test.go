package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	types "internal/pkg"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestUserNewHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)

	t.Run("should faild if the user payload is invalid", func(t *testing.T) {
		payload := &types.RegisterUserPayload{
			Username: "user",
			Email:    "invalid",
			Password: "1213313",
		}
		marshelled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshelled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status code %v, got %v", http.StatusBadRequest, rr.Code)
		}
	})

	t.Run("should correctly register a new user", func(t *testing.T) {
		payload := &types.RegisterUserPayload{
			Username: "user",
			Email:    "sedfcwefwe@lcledc.com",
			Password: "13213141",
		}
		marshalled, _ := json.Marshal(payload)

		req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()

		router.HandleFunc("/register", handler.handleRegister)
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Errorf("expected status code %v, got %v", http.StatusCreated, rr.Code)
			t.Logf("response body: %v", rr.Body.String())
		}
	})
}

type mockUserStore struct{}

func (m *mockUserStore) GetUserByEmail(email string) (*types.User, error) {
	return nil, fmt.Errorf("user not found")
}

func (m *mockUserStore) CreateUser(user types.User) error {
	return nil
}

func (m *mockUserStore) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
