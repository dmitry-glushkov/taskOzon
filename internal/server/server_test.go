package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_SetHandler(t *testing.T) {
	server := NewServer()
	testCases := []struct {
		name         string
		payload      interface{}
		expectedCode int
	}{
		{
			name: "valid without ttl",
			payload: map[string]interface{}{
				"key":   "1",
				"value": "valid value",
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "valid with ttl",
			payload: map[string]interface{}{
				"key":   "2",
				"value": "val",
				"ttl":   1,
			},
			expectedCode: http.StatusCreated,
		},
		{
			name: "invalid value",
			payload: map[string]interface{}{
				"key": "3",
			},
			expectedCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/set", b)
			server.ServeHTTP(rec, req)
			assert.Equal(t, tc.expectedCode, rec.Code)
		})
	}
}

func TestServer_GetHandler(t *testing.T) {

}

func TestServer_GetAllKeysHandler(t *testing.T) {
	// server := NewServer()
	// server.store.Set("1", "val", 0)
	// server.store.Set("2", "val", 0)

	// var result []string
	// b := &bytes.Buffer{}
	// rec := httptest.NewRecorder()
	// req, _ := http.NewRequest(http.MethodGet, "/getallkeys", b)
	// json.NewDecoder(b).Decode(result)
	// server.ServeHTTP(rec, req)
	// assert.Equal(t, []string{"1", "2"}, result)
}

func TestServer_DeleteHandler(t *testing.T) {

}
