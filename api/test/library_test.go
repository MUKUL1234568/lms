package routes_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"library-management-api/config"
	"library-management-api/models"
	"library-management-api/routes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Mock Database Setup
func initTestDB() {
	var err error
	config.DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	config.DB.AutoMigrate(&models.Library{}, &models.User{}, &models.Book{})
}

// Setup Router for Testing
func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	initTestDB()
	r := gin.Default()
	routes.LibraryRoutes(r)
	return r
}

// Test Creating a Library
func TestLibraryRoutes_CreateLibrary(t *testing.T) {
	r := setupRouter()
	requestBody := map[string]string{
		"library_name":   "alpha",
		"owner_name":     "mukul",
		"owner_email":    "mukul@example.com",
		"owner_password": "mukul123",
		"owner_contact":  "9876543210",
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Validate Response
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "Library created successfully")

	// Log success message
	fmt.Println("✅ TestLibraryRoutes_CreateLibrary Passed")
}
func TestLibraryRoutes_CreateLibrarywithregistredmail(t *testing.T) {
	r := setupRouter()
	requestBody := map[string]string{
		"library_name":   "alpha",
		"owner_name":     "mukul",
		"owner_email":    "mukul@example.com",
		"owner_password": "mukul123",
		"owner_contact":  "9876543210",
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Validate Response
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Contains(t, rr.Body.String(), "user with this email already exists")

	// Log success message
	fmt.Println("✅  TestLibraryRoutes_CreateLibrarywithregistredmail")
}

// Test Fetching All Libraries
func TestLibraryRoutes_GetAllLibraries(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/libraries/", nil)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Validate Response
	assert.Equal(t, http.StatusOK, rr.Code)

	// Log success message
	fmt.Println("✅ TestLibraryRoutes_GetAllLibraries Passed")
}

// Test Database State Validation
func TestLibraryRoutes_DatabaseState(t *testing.T) {
	var library models.Library
	err := config.DB.Where("name = ?", "alpha").First(&library).Error
	assert.NoError(t, err, "Library should be saved in DB")

	// Log success message
	fmt.Println("✅ TestLibraryRoutes_DatabaseState Passed")
}

// Test API Performance
func TestLibraryRoutes_Performance(t *testing.T) {
	r := setupRouter()
	req, _ := http.NewRequest("GET", "/libraries/", nil)
	rr := httptest.NewRecorder()

	start := time.Now()
	r.ServeHTTP(rr, req)
	duration := time.Since(start)

	assert.Less(t, duration.Milliseconds(), int64(50), "API should respond within 50ms")

	// Log success message
	fmt.Println("✅ TestLibraryRoutes_Performance Passed")
}
