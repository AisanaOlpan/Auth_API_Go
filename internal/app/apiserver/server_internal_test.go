package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AisanaOlpan/GoProject/internal/app/model"
	"github.com/AisanaOlpan/GoProject/internal/app/store/teststore"
	"github.com/go-playground/assert/v2"
	"github.com/gorilla/sessions"
)

func TestServer_HandleUsersCreate(t *testing.T) {

	s := newServer(teststore.New(), sessions.NewCookieStore([]byte("secret")))
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "ais@mail.ru",
				"password": "password",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name:         "invalid",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, rec.Code)
		})
	}
	// rec := httptest.NewRecorder()
	// req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	// s := newServer(teststore.New())
	// s.ServeHTTP(rec, req)
	// assert.Equal(t, rec.Code, http.StatusOK)
	// fmt.Print(rec.Code)
}

func TestServer_HandleUsersFindByEmail(t *testing.T) {
	u := model.TestUser(t) //когда запускаем тест, тестовая база бывает пустым поэтому сначала создаем юзера
	store := teststore.New()
	store.User().Create(u)
	sessionStrore := sessions.NewCookieStore([]byte("secret"))
	s := newServer(store, sessionStrore)
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "ais@mail.ru",
				"password": "password",
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid",
			payload:      "invalid",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)
			s.ServeHTTP(rec, req)
			assert.Equal(t, rec.Code, rec.Code)
		})
	}
}
