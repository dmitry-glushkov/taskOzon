package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
	server := NewServer()
	server.store.Set("1", "val", 0)

	b := &bytes.Buffer{}
	payload := map[string]interface{}{
		"key": "1",
	}
	json.NewEncoder(b).Encode(payload)
	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	c.Request, _ = http.NewRequest("POST", "/get", b)
	server.ServeHTTP(resp, c.Request)
	var result string
	json.Unmarshal(resp.Body.Bytes(), &result)
	assert.Equal(t, "val", result)
}

func TestServer_GetAllKeysHandler(t *testing.T) {
	server := NewServer()
	server.store.Set("1", "val", 0)
	server.store.Set("2", "val", 0)

	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)
	c.Request, _ = http.NewRequest("GET", "/getallkeys", nil)
	server.ServeHTTP(resp, c.Request)
	var result []string
	json.Unmarshal(resp.Body.Bytes(), &result)
	assert.Equal(t, []string{"1", "2"}, result)
}
