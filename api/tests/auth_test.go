// package tests

// import (
// 	"bytes"
// 	"encoding/json"

// 	// "library-management-api/config"
// 	// "library-management-api/routes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	// "github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	// "gorm.io/driver/sqlite"
// 	// "gorm.io/gorm"

// 	"fmt"
// 	// "golang.org/x/crypto/bcrypt"
// 	// "library-management-api/models"
// 	// "library-management-api/services"
// )

// func TestLogin_Success(t *testing.T) {
// 	r := SetUpauthroutes()

// 	requestBody := map[string]string{
// 		"email":    "mukul@example.com", // Replace with actual test user
// 		"password": "mukul123",          // Replace with actual password
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "Login successful")

// 	fmt.Println("✅ TestLogin_Success Passed") // Explicit success message
// }

// func TestLogin_InvalidPassword(t *testing.T) {
// 	r := SetUpauthroutes()

// 	requestBody := map[string]string{
// 		"email":    "mukul@example.com", // Use actual test user email
// 		"password": "wrongpassword",     // Incorrect password
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusUnauthorized, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "invalid password")

// 	fmt.Println("✅ TestLogin_InvalidPassword Passed") // Explicit success message
// }

// func TestLogin_InvalidEmail(t *testing.T) {
// 	r := SetUpauthroutes()
// 	// RegisterUser handles user registration

// 	requestBody := map[string]string{
// 		"email":    "nonexistent@example.com", // An email not in test.db
// 		"password": "somepassword",
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusUnauthorized, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "invalid credentials")

// 	fmt.Println("✅ TestLogin_InvalidEmail Passed") // Explicit success message
// }

// // // go test -coverprofile=coverage ./...
package tests
