// package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"library-management-api/config"
// 	"library-management-api/models"
// 	"library-management-api/routes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"time"

// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// // Mock Database Setup
// var testDB *gorm.DB

// // Initialize shared test DB
// func initTestDB() {
// 	var err error
// 	if testDB == nil {
// 		testDB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// 		if err != nil {
// 			panic("Failed to connect to test database")
// 		}
// 		testDB.AutoMigrate(&models.Library{}, &models.User{}, &models.Book{})
// 	}
// 	config.DB = testDB
// }

// // Setup Router for Testing
// func setupRouter() *gin.Engine {
// 	// gin.SetMode(gin.TestMode)
// 	initTestDB()
// 	r := gin.Default()
// 	routes.LibraryRoutes(r)
// 	return r
// }

// // Test Creating a Library
// func TestLibraryRoutes_CreateLibrary(t *testing.T) {
// 	r := setupRouter()
// 	requestBody := map[string]string{
// 		"library_name":   "alpha",
// 		"owner_name":     "mukul",
// 		"owner_email":    "mukul@example.com",
// 		"owner_password": "mukul123",
// 		"owner_contact":  "9876543210",
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)

// 	// Validate Response
// 	assert.Equal(t, http.StatusCreated, rr.Code)
// 	assert.Contains(t, rr.Body.String(), "Library created successfully")

// 	// Log success message
// 	fmt.Println("✅ TestLibraryRoutes_CreateLibrary Passed")
// }
// func TestLibraryRoutes_CreateLibrarywithregistredmail(t *testing.T) {
// 	r := setupRouter()
// 	requestBody := map[string]string{
// 		"library_name":   "alpha",
// 		"owner_name":     "mukul",
// 		"owner_email":    "mukul@example.com",
// 		"owner_password": "mukul123",
// 		"owner_contact":  "9876543210",
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)

// 	// Validate Response
// 	assert.Equal(t, http.StatusConflict, rr.Code)
// 	assert.Contains(t, rr.Body.String(), "user with this email already exists")

// 	// Log success message
// 	fmt.Println("✅  TestLibraryRoutes_CreateLibrarywithregistredmail")
// }

// // Test Fetching All Libraries
// func TestLibraryRoutes_GetAllLibraries(t *testing.T) {
// 	r := setupRouter()
// 	req, _ := http.NewRequest("GET", "/libraries/", nil)
// 	req.Header.Set("Content-Type", "application/json")

// 	rr := httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)

// 	// Validate Response
// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	// Log success message
// 	fmt.Println("✅ TestLibraryRoutes_GetAllLibraries Passed")
// }

// // Test Database State Validation
// func TestLibraryRoutes_DatabaseState(t *testing.T) {
// 	var library models.Library
// 	err := config.DB.Where("name = ?", "alpha").First(&library).Error
// 	assert.NoError(t, err, "Library should be saved in DB")

// 	// Log success message
// 	fmt.Println("✅ TestLibraryRoutes_DatabaseState Passed")
// }

// // Test API Performance
// func TestLibraryRoutes_Performance(t *testing.T) {
// 	r := setupRouter()
// 	req, _ := http.NewRequest("GET", "/libraries/", nil)
// 	rr := httptest.NewRecorder()

// 	start := time.Now()
// 	r.ServeHTTP(rr, req)
// 	duration := time.Since(start)

// 	assert.Less(t, duration.Milliseconds(), int64(50), "API should respond within 50ms")

// 	// Log success message
// 	fmt.Println("✅ TestLibraryRoutes_Performance Passed")
// }

package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"library-management-api/config"
	"library-management-api/models"
	"library-management-api/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Mock Database Setup
var testDB *gorm.DB
var router *gin.Engine

// Test Main (Runs Once Before All Tests)
func TestMain(m *testing.M) {
	fmt.Println("🚀 Setting up test environment...")

	// Initialize In-Memory DB (Single Connection for All Tests)
	var err error
	testDB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	testDB.AutoMigrate(&models.Library{}, &models.User{}, &models.Book{})
	config.DB = testDB

	// Setup Router Once
	router = gin.Default()
	routes.LibraryRoutes(router)

	// Run Tests
	code := m.Run()

	// Cleanup (Optional)
	fmt.Println("🛑 Cleaning up test environment...")

	os.Exit(code)
}

// Test Creating a Library
func TestLibraryRoutes_CreateLibrary(t *testing.T) {
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
	router.ServeHTTP(rr, req)

	// Validate Response
	assert.Equal(t, http.StatusCreated, rr.Code)
	assert.Contains(t, rr.Body.String(), "Library created successfully")

	fmt.Println("✅ TestLibraryRoutes_CreateLibrary Passed")
}

// Test Creating a Library with an Already Registered Email
func TestLibraryRoutes_CreateLibrarywithRegisteredEmail(t *testing.T) {
	requestBody := map[string]string{
		"library_name":   "alpha2",
		"owner_name":     "mukul",
		"owner_email":    "mukul@example.com", // Same email as previous test
		"owner_password": "mukul123",
		"owner_contact":  "9876543210",
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Validate Response: Should return 409 Conflict
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Contains(t, rr.Body.String(), "user with this email already exists")

	fmt.Println("✅ TestLibraryRoutes_CreateLibrarywithRegisteredEmail Passed")
}

// Test Fetching All Libraries
func TestLibraryRoutes_GetAllLibraries(t *testing.T) {
	req, _ := http.NewRequest("GET", "/libraries/", nil)
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	fmt.Println("✅ TestLibraryRoutes_GetAllLibraries Passed")
}

// Test Database State Validation
func TestLibraryRoutes_DatabaseState(t *testing.T) {
	var library models.Library
	err := config.DB.Where("name = ?", "alpha").First(&library).Error
	assert.NoError(t, err, "Library should be saved in DB")

	fmt.Println("✅ TestLibraryRoutes_DatabaseState Passed")
}

// Test API Performance
func TestLibraryRoutes_Performance(t *testing.T) {
	req, _ := http.NewRequest("GET", "/libraries/", nil)
	rr := httptest.NewRecorder()

	start := time.Now()
	router.ServeHTTP(rr, req)
	duration := time.Since(start)

	assert.Less(t, duration.Milliseconds(), int64(50), "API should respond within 50ms")

	fmt.Println("✅ TestLibraryRoutes_Performance Passed")
}
