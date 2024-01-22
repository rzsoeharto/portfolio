package auth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	r := gin.New()

	r.POST("/login", Login)

	reqBody := `{ "username": "testuser", "password": "testpassword" }`
	req := httptest.NewRequest("POST", "/login", strings.NewReader(reqBody))

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseBody map[string]interface{}

	err := json.Unmarshal(w.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "testuser", responseBody["username"])
	assert.Contains(t, responseBody, "name")
}
