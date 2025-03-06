package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"library-management-api/config"
	"library-management-api/models"
	// "library-management-api/routes"
	"net/http"
	"net/http/httptest"
	// "os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// Test Creating a Library
func TestLibraryRoutes_CreateLibrary(t *testing.T) {
	requestBody := map[string]string{
		"library_name":   "alpha",
		"owner_name":     "mukulowner",
		"owner_email":    "mukulowner@example.com",
		"owner_password": "mukul123",
		"owner_contact":  "9876543210",
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router := setupLibraryRouter()
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
		"owner_email":    "mukulowner@example.com", // Same email as previous test
		"owner_password": "mukul123",
		"owner_contact":  "9876543210",
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router := setupLibraryRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Validate Response: Should return 409 Conflict
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Contains(t, rr.Body.String(), "user with this email already exists")

	fmt.Println("✅ TestLibraryRoutes_CreateLibrarywithRegisteredEmail Passed")
}
func TestLibraryRoutes_CreateLibrarywithRegisteredlibname(t *testing.T) {
	requestBody := map[string]string{
		"library_name":   "alpha",
		"owner_name":     "mukul",
		"owner_email":    "mukulowner1@example.com", // Same email as previous test
		"owner_password": "mukul123",
		"owner_contact":  "9876543210",
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router := setupLibraryRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Validate Response: Should return 409 Conflict
	assert.Equal(t, http.StatusConflict, rr.Code)
	assert.Contains(t, rr.Body.String(), "choose different library name")

	fmt.Println("✅ TestLibraryRoutes_CreateLibrarywithRegisteredEmail Passed")
}

func TestLibraryRoutes_CreateLibrarywithinvalidnumber(t *testing.T) {
	requestBody := map[string]string{
		"library_name":   "alpha2",
		"owner_name":     "mukul",
		"owner_email":    "mukulowner@example.com", // Same email as previous test
		"owner_password": "mukul123",
		"owner_contact":  "987654",
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/libraries/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router := setupLibraryRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Validate Response: Should return 409 Conflict
	assert.Equal(t, http.StatusBadRequest, rr.Code)
	// assert.Contains(t, rr.Body.String(), "user with this email already exists")

	fmt.Println("✅ TestLibraryRoutes_CreateLibrarywithRegisteredEmail Passed")
}

// Test Fetching All Libraries
func TestLibraryRoutes_GetAllLibraries(t *testing.T) {
	req, _ := http.NewRequest("GET", "/libraries/", nil)
	req.Header.Set("Content-Type", "application/json")
	router := setupLibraryRouter()
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
	router := setupLibraryRouter()
	start := time.Now()
	router.ServeHTTP(rr, req)
	duration := time.Since(start)

	assert.Less(t, duration.Milliseconds(), int64(50), "API should respond within 50ms")

	fmt.Println("✅ TestLibraryRoutes_Performance Passed")
}

func TestUserRoutes_RegisterUser(t *testing.T) {
	requestBody := map[string]interface{}{
		"name":           "Mukul",
		"email":          "mukul@example.com",
		"password":       "123",
		"contact_number": "9876543210",
		"lib_id":         1,
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router := setupUserRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	// assert.Contains(t, rr.Body.String(), "User registered successfully")

	fmt.Println("✅ TestUserRoutes_RegisterUser Passed")
}

func TestUserRoutes_RegisterUserwithnotlibrary(t *testing.T) {
	requestBody := map[string]interface{}{
		"name":           "Mukul",
		"email":          "mukul@example.com",
		"password":       "123",
		"contact_number": "9876543210",
		"lib_id":         2,
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router := setupUserRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	// assert.Contains(t, rr.Body.String(), "User registered successfully not regi library")

	fmt.Println("✅ TestUserRoutes_RegisterUser Passed")
}

// Test Duplicate Email Registration
func TestUserRoutes_RegisterDuplicateEmail(t *testing.T) {
	requestBody := map[string]interface{}{
		"name":           "Alice",
		"email":          "mukul@example.com",
		"password":       "securepassword",
		"contact_number": "9876543210",
		"lib_id":         1,
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router := setupUserRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusInternalServerError, rr.Code)
	// assert.Contains(t, rr.Body.String(), "user with this email already exists")

	fmt.Println("✅ TestUserRoutes_RegisterDuplicateEmail Passed")
}

// Test Invalid Contact Number (Less than 10 Digits)
func TestUserRoutes_RegisterInvalidContact(t *testing.T) {
	requestBody := map[string]interface{}{
		"name":           "Bob",
		"email":          "bob@example.com",
		"password":       "password",
		"contact_number": "12345", // Invalid contact number
		"lib_id":         1,
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router := setupUserRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	// assert.Contains(t, rr.Body.String(), "Invalid phone number")

	fmt.Println("✅ TestUserRoutes_RegisterInvalidContact Passed")
}

// Test Invalid Email Format
func TestUserRoutes_RegisterInvalidEmail(t *testing.T) {
	requestBody := map[string]interface{}{
		"name":           "Charlie",
		"email":          "invalid-email",
		"password":       "password",
		"contact_number": "9876543210",
		"lib_id":         1,
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	router := setupUserRouter()
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	// assert.Contains(t, rr.Body.String(), "Invalid email format")

	fmt.Println("✅ TestUserRoutes_RegisterInvalidEmail Passed")
}

func TestUserRoutes_GetProfile(t *testing.T) {
	router := setupUserRouter()

	// Get a valid token using specific credentials
	token := getAuthToken("mukul@example.com", "123")

	req, _ := http.NewRequest("GET", "/user/profile", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token) // Attach token

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// assert.Contains(t, rr.Body.String(), "Test User") // Adjust based on expected response
	fmt.Println(token)
	fmt.Println("✅ TestUserRoutes_GetProfile Passed")
}

func TestUserRoutes_GetallProfile(t *testing.T) {
	router := setupUserRouter()

	// Get a valid token using specific credentials
	token := getAuthToken("mukulowner@example.com", "mukul123")

	req, _ := http.NewRequest("GET", "/user/all", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token) // Attach token

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// assert.Contains(t, rr.Body.String(), "Test User") // Adjust based on expected response
	fmt.Println(token)
	fmt.Println("✅ TestUserRoutes_all user Passed")
}

func TestLogin_Success(t *testing.T) {
	r := SetUpauthroutes()

	requestBody := map[string]string{
		"email":    "mukul@example.com", // Replace with actual test user
		"password": "123",               // Replace with actual password
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// assert.Contains(t, rr.Body.String(), "Login successful")

	fmt.Println("✅ TestLogin_Success Passed") // Explicit success message
}

func TestLogin_InvalidPassword(t *testing.T) {
	r := SetUpauthroutes()

	requestBody := map[string]string{
		"email":    "mukul@example.com", // Use actual test user email
		"password": "wrongpassword",     // Incorrect password
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	// assert.Contains(t, rr.Body.String(), "invalid password")

	fmt.Println("✅ TestLogin_InvalidPassword Passed") // Explicit success message
}

func TestLogin_InvalidEmail(t *testing.T) {
	r := SetUpauthroutes()
	// RegisterUser handles user registration

	requestBody := map[string]string{
		"email":    "nonexistent@example.com", // An email not in test.db
		"password": "somepassword",
	}
	body, _ := json.Marshal(requestBody)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)
	// assert.Contains(t, rr.Body.String(), "invalid credentials")

	fmt.Println("✅ TestLogin_InvalidEmail Passed") // Explicit success message
}

func TestAddBook(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukulowner@example.com", "mukul123")
	book := models.Book{
		ISBN:        "1234567892345",
		Title:       "Test Book",
		Authors:     "John Doe",
		Publisher:   "Test Publisher",
		Version:     "1st",
		TotalCopies: 5,
	}

	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	// var response map[string]interface{}
	// json.Unmarshal(w.Body.Bytes(), &response)
	// assert.Equal(t, "Book added successfully", response["message"])
}

func TestAddBookwithinvalidisbn(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukulowner@example.com", "mukul123")
	book := models.Book{
		ISBN:        "12345678923",
		Title:       "Test Book",
		Authors:     "John Doe",
		Publisher:   "Test Publisher",
		Version:     "1st",
		TotalCopies: 5,
	}

	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

}

func TestAddBookwithwrongowner(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukuloner@example.com", "mukul123")
	book := models.Book{
		ISBN:        "12345678923",
		Title:       "Test Book",
		Authors:     "John Doe",
		Publisher:   "Test Publisher",
		Version:     "1st",
		TotalCopies: 5,
	}

	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

}

func TestAddBookbookalreadyindtabase(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukulowner@example.com", "mukul123")
	book := models.Book{
		ISBN:        "1234567892345",
		Title:       "Test Book",
		Authors:     "John Doe",
		Publisher:   "Test Publisher",
		Version:     "1st",
		TotalCopies: 5,
	}

	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books/", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestGetallbookbylib(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukulowner@example.com", "mukul123")

	req, _ := http.NewRequest("GET", "/books/lib", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestGetallBookbyLibwithunautorizedperson(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("muklowner@example.com", "mukul123")

	req, _ := http.NewRequest("GET", "/books/lib", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

}

func TestGetbookbyisbn(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukul@example.com", "123")

	req, _ := http.NewRequest("GET", "/books/1234567892345", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}
func TestGetbookbyisbnwithwrong(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukul@example.com", "123")

	req, _ := http.NewRequest("GET", "/books/1234562892345", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

}

func TestGetbookbyisbnwithwrongperson(t *testing.T) {

	r := setupBookRouter()
	//token := getAuthToken("mukul@example.com", "123")

	req, _ := http.NewRequest("GET", "/books/1234562892345", nil)
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusUnauthorized, rr.Code)

}

func TestUpdatebook(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukulowner@example.com", "mukul123")
	book := models.Book{

		Version:     "2st",
		TotalCopies: 7,
	}

	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("PATCH", "/books/1234567892345", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestUpdatebookwithwrogcopies(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukulowner@example.com", "mukul123")
	book := models.Book{

		Version:         "2st",
		TotalCopies:     7,
		AvailableCopies: 8,
	}

	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("PATCH", "/books/1234567892345", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

}

func TestUpdatebookwithunauthorized(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukul@example.com", "123")
	book := models.Book{

		Version:     "2st",
		TotalCopies: 7,
	}

	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("PATCH", "/books/1234567892345", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)

}

func TestUpdatebookwithwrongisbn(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukulowner@example.com", "mukul123")
	book := models.Book{

		Version:     "2st",
		TotalCopies: 7,
	}

	jsonValue, _ := json.Marshal(book)
	req, _ := http.NewRequest("PATCH", "/books/1234567892340", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)

}

func TestDeletebook(t *testing.T) {

	r := setupBookRouter()
	token := getAuthToken("mukulowner@example.com", "mukul123")

	req, _ := http.NewRequest("DELETE", "/books/1234567892345", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestGetIssuedBookforOwner(t *testing.T) {

	r := SetIssueroutes()
	token := getAuthToken("mukulowner@example.com", "mukul123")

	req, _ := http.NewRequest("GET", "/issueregistry/admin", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestGetIssuedBookforreader(t *testing.T) {

	r := SetIssueroutes()
	token := getAuthToken("mukul@example.com", "123")

	req, _ := http.NewRequest("GET", "/issueregistry/user", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestAddrequest(t *testing.T) {

	r := SetrequestsRoutes()
	token := getAuthToken("mukul@example.com", "123")

	BBody := map[string]string{
		"request_type": "Issue",
		"isbn":         "1234567892345",
	}
	body, _ := json.Marshal(BBody)

	req, _ := http.NewRequest("POST", "/requests/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

}
func TestGetRequetbyuser(t *testing.T) {

	r := SetrequestsRoutes()
	token := getAuthToken("mukul@example.com", "123")

	req, _ := http.NewRequest("GET", "/requests/user", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestGetRequetbyowner(t *testing.T) {

	r := SetrequestsRoutes()
	token := getAuthToken("mukulowner@example.com", "mukul123")

	req, _ := http.NewRequest("GET", "/requests/allreq", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestApproverequet(t *testing.T) {

	r := SetrequestsRoutes()
	token := getAuthToken("mukulowner@example.com", "mukul123")

	BBody := map[string]bool{
		"approve": true,
	}
	body, _ := json.Marshal(BBody)

	req, _ := http.NewRequest("PUT", "/requests/1/approve", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}

func TestApproverequetwithwronguser(t *testing.T) {

	r := SetrequestsRoutes()
	token := getAuthToken("mukul@example.com", "123")

	BBody := map[string]bool{
		"approve": true,
	}
	body, _ := json.Marshal(BBody)

	req, _ := http.NewRequest("PUT", "/requests/1/approve", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusForbidden, rr.Code)

}

func TestAddrequestt(t *testing.T) {

	r := SetrequestsRoutes()
	token := getAuthToken("mukul@example.com", "123")

	BBody := map[string]string{
		"request_type": "Return",
		"isbn":         "1234567892345",
	}
	body, _ := json.Marshal(BBody)

	req, _ := http.NewRequest("POST", "/requests/", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

}

func TestApproverequett(t *testing.T) {

	r := SetrequestsRoutes()
	token := getAuthToken("mukulowner@example.com", "mukul123")

	BBody := map[string]bool{
		"approve": true,
	}
	body, _ := json.Marshal(BBody)

	req, _ := http.NewRequest("PUT", "/requests/2/approve", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

}
