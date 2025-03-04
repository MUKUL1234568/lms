// package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	// "library-management-api/config"
// 	// "library-management-api/models"
// 	// "library-management-api/routes"
// 	"net/http"
// 	"net/http/httptest"
// 	// "os"
// 	"testing"

// 	// "github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	// "gorm.io/driver/sqlite"
// 	// "gorm.io/gorm"
// )

// // Mock Database Setup
// // var testDB *gorm.DB
// // var router *gin.Engine

// // Test Main (Runs Once Before All Tests)
// // func TestMain(m *testing.M) {
// // 	fmt.Println("🚀 Setting up test environment...")

// // 	// Initialize In-Memory DB (Single Connection for All Tests)
// // 	var err error
// // 	testDB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
// // 	if err != nil {
// // 		panic("Failed to connect to test database")
// // 	}
// // 	testDB.AutoMigrate(&models.Library{}, &models.User{}, &models.Book{})
// // 	config.DB = testDB

// // 	// Setup Router Once
// // 	router = gin.Default()
// // 	routes.UserRoutes(router)
// // 	routes.LibraryRoutes(router)

// // 	// Create a library before adding users
// // 	createLibrary()

// // 	// Run Tests
// // 	code := m.Run()

// // 	// Cleanup (Optional)
// // 	fmt.Println("🛑 Cleaning up test environment...")

// // 	os.Exit(code)
// // }

// // Helper function to create a library
// // func createLibrary() {
// // 	requestBody := map[string]string{
// // 		"library_name":   "Test Library",
// // 		"owner_name":     "John Doe",
// // 		"owner_email":    "owner@example.com",
// // 		"owner_password": "password123",
// // 		"owner_contact":  "9876543210",
// // 	}
// // 	body, _ := json.Marshal(requestBody)

// // 	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
// // 	req.Header.Set("Content-Type", "application/json")
// // 	router := setupUserRouter()
// // 	rr := httptest.NewRecorder()
// // 	router.ServeHTTP(rr, req)

// // 	assert.Equal(nil, http.StatusCreated, rr.Code)
// // 	fmt.Println("✅ Library created successfully for user registration tests")
// // }

// // Test Successful User Registration
// func TestUserRoutes_RegisterUser(t *testing.T) {
// 	requestBody := map[string]interface{}{
// 		"name":           "Mukul",
// 		"email":          "mukul@example.com",
// 		"password":       "123",
// 		"contact_number": "9876543210",
// 		"lib_id":         1,
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	router := setupUserRouter()
// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusCreated, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "User registered successfully")

// 	fmt.Println("✅ TestUserRoutes_RegisterUser Passed")
// }

// func TestUserRoutes_RegisterUserwithnotlibrary(t *testing.T) {
// 	requestBody := map[string]interface{}{
// 		"name":           "Mukul",
// 		"email":          "mukul@example.com",
// 		"password":       "123",
// 		"contact_number": "9876543210",
// 		"lib_id":         2,
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	router := setupUserRouter()
// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusBadRequest, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "User registered successfully")

// 	fmt.Println("✅ TestUserRoutes_RegisterUser Passed")
// }

// // Test Duplicate Email Registration
// func TestUserRoutes_RegisterDuplicateEmail(t *testing.T) {
// 	requestBody := map[string]interface{}{
// 		"name":           "Alice",
// 		"email":          "mukul@example.com",
// 		"password":       "securepassword",
// 		"contact_number": "9876543210",
// 		"lib_id":         1,
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	router := setupUserRouter()
// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusInternalServerError, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "user with this email already exists")

// 	fmt.Println("✅ TestUserRoutes_RegisterDuplicateEmail Passed")
// }

// // Test Invalid Contact Number (Less than 10 Digits)
// func TestUserRoutes_RegisterInvalidContact(t *testing.T) {
// 	requestBody := map[string]interface{}{
// 		"name":           "Bob",
// 		"email":          "bob@example.com",
// 		"password":       "password",
// 		"contact_number": "12345", // Invalid contact number
// 		"lib_id":         1,
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	router := setupUserRouter()
// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusBadRequest, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "Invalid phone number")

// 	fmt.Println("✅ TestUserRoutes_RegisterInvalidContact Passed")
// }

// // Test Invalid Email Format
// func TestUserRoutes_RegisterInvalidEmail(t *testing.T) {
// 	requestBody := map[string]interface{}{
// 		"name":           "Charlie",
// 		"email":          "invalid-email",
// 		"password":       "password",
// 		"contact_number": "9876543210",
// 		"lib_id":         1,
// 	}
// 	body, _ := json.Marshal(requestBody)

// 	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	router := setupUserRouter()
// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusBadRequest, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "Invalid email format")

// 	fmt.Println("✅ TestUserRoutes_RegisterInvalidEmail Passed")
// }

// func TestUserRoutes_GetProfile(t *testing.T) {
// 	router := setupUserRouter()

// 	// Get a valid token using specific credentials
// 	token := getAuthToken("mukul@example.com", "123")

// 	req, _ := http.NewRequest("GET", "/user/profile", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+token) // Attach token

// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "Test User") // Adjust based on expected response
// 	fmt.Println(token)
// 	fmt.Println("✅ TestUserRoutes_GetProfile Passed")
// }

// func TestUserRoutes_GetallProfile(t *testing.T) {
// 	router := setupUserRouter()

// 	// Get a valid token using specific credentials
// 	token := getAuthToken("mukulowner@example.com", "mukul123")

// 	req, _ := http.NewRequest("GET", "/user/all", nil)
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+token) // Attach token

// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	// assert.Contains(t, rr.Body.String(), "Test User") // Adjust based on expected response
// 	fmt.Println(token)
// 	fmt.Println("✅ TestUserRoutes_GetProfile Passed")
// }

// func TestUserRoutes_makeAdmin(t *testing.T) {
// 	router := setupUserRouter()

// 	// Get a valid token using specific credentials
// 	token := getAuthToken("mukulowner@example.com", "mukul123")

// 	// Only send the role in the body (not the ID)
// 	BBody := map[string]string{
// 		"role": "Admin",
// 	}
// 	body, _ := json.Marshal(BBody)

// 	// Use the user ID in the URL, not in the body
// 	req, _ := http.NewRequest("PUT", "/role/2", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+token) // Attach token

// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

// 	// Check if the response status code is OK (200)
// 	assert.Equal(t, http.StatusOK, rr.Code)

// 	// Optionally, check the response body for confirmation of success
// 	// assert.Contains(t, rr.Body.String(), "Admin") // Adjust based on expected response

// 	fmt.Println(token)
// 	fmt.Println("✅ TestUserRoutes_makeAdmin Passed")
// }

// func TestUserRoutes_makeaddmin1(t *testing.T) {
// 	router := setupUserRouter()

// 	// Get a valid token using specific credentials
// 	token := getAuthToken("mukulowner@example.com", "mukul123")
// 	BBody := map[string]string{
// 		"role": "Owner",
// 	}
// 	body, _ := json.Marshal(BBody)

// 	req, _ := http.NewRequest("PUT", "/role/2", bytes.NewBuffer(body))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+token) // Attach token

// 	rr := httptest.NewRecorder()
// 	router.ServeHTTP(rr, req)

//		assert.Equal(t, http.StatusOK, rr.Code)
//		// assert.Contains(t, rr.Body.String(), "Test User") // Adjust based on expected response
//		fmt.Println(token)
//		fmt.Println("✅ TestUserRoutes_GetProfile Passed")
//	}
package tests
