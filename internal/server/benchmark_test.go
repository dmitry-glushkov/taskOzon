package server

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func BenchmarkServer_Set(b *testing.B) {
	server := NewServer(100)
	for i := 0; i < b.N; i++ {
		resp := httptest.NewRecorder()
		gin.SetMode(gin.TestMode)
		c, _ := gin.CreateTestContext(resp)
		b := &bytes.Buffer{}
		payload := map[string]interface{}{
			"key":   "1",
			"value": "val",
		}
		json.NewEncoder(b).Encode(payload)
		c.Request, _ = http.NewRequest(http.MethodPost, "/set", b)
		server.ServeHTTP(resp, c.Request)
	}
}
