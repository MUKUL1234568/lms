package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"library-management-api/config"
	"library-management-api/models"
	"library-management-api/routes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	// "github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

// Test Main (Runs Once Before All Tests)
func TestMain(m *testing.M) {
	fmt.Println("🚀 Setting up test environment...")

	// Initialize In-Memory DB (Single Connection for All Tests)
	var err error
	testDB, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to test database")
	}
	testDB.AutoMigrate(&models.Library{}, &models.User{}, &models.Book{}, &models.RequestEvent{}, &models.IssueRegistry{})
	config.DB = testDB

	// Setup Router Once

	// Run Tests
	code := m.Run()

	// Cleanup (Optional)
	fmt.Println("🛑 Cleaning up test environment...")

	os.Exit(code)
}

func setupLibraryRouter() *gin.Engine {
	r := gin.Default()
	routes.LibraryRoutes(r)
	return r
}

func setupBookRouter() *gin.Engine {
	r := gin.Default()
	routes.BookRoutes(r)
	return r
}

func setupUserRouter() *gin.Engine {
	r := gin.Default()
	routes.UserRoutes(r)
	return r
}

func SetUpauthroutes() *gin.Engine {
	r := gin.Default()
	routes.AuthRoutes(r)
	return r
}

func SetIssueroutes() *gin.Engine {
	r := gin.Default()
	routes.IssueRegistryRoutes(r)
	return r
}

func SetrequestsRoutes() *gin.Engine {
	r := gin.Default()
	routes.RequestRoutes(r)
	return r
}

// Function to retrieve token with custom credentials
func getAuthToken(email, password string) string {
	// Prepare login request
	router := SetUpauthroutes()
	loginBody := map[string]string{
		"email":    email,
		"password": password,
	}
	body, _ := json.Marshal(loginBody)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Debugging: Print the raw response body
	fmt.Println("🔍 Login Response:", rr.Body.String())

	// Check if the response is not 200 OK
	if rr.Code != http.StatusOK {
		fmt.Printf("❌ Login failed! Status Code: %d\n", rr.Code)
		return ""
	}

	// Try to parse the response
	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	if err != nil {
		fmt.Println("❌ Error decoding JSON response:", err)
		return ""
	}

	// Check if token exists
	token, exists := response["token"]
	if !exists || token == "" {
		fmt.Println("❌ Token not found in response")
		return ""
	}

	fmt.Println("✅ Token Retrieved:", token)
	return token
}
