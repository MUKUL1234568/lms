// package tests

// import (
// 	"bytes"
// 	"encoding/json"
// 	// "library-management-api/controllers"
// 	"library-management-api/models"
// 	// "library-management-api/services"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	//"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// )

// func TestAddBook(t *testing.T) {

// 	r := setupBookRouter()
// 	token := getAuthToken("mukulowner@example.com", "mukul123")

// 	book := models.Book{
// 		ISBN:        "123456789",
// 		Title:       "Test Book",
// 		Authors:     "John Doe",
// 		Publisher:   "Test Publisher",
// 		Version:     "1st",
// 		TotalCopies: 5,
// 	}

// 	jsonValue, _ := json.Marshal(book)
// 	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(jsonValue))
// 	req.Header.Set("Content-Type", "application/json")
// 	req.Header.Set("Authorization", "Bearer "+token)

// 	rr := httptest.NewRecorder()
// 	r.ServeHTTP(rr, req)

// 	assert.Equal(t, http.StatusCreated, rr.Code)
// 	// var response map[string]interface{}
// 	// json.Unmarshal(w.Body.Bytes(), &response)
// 	// assert.Equal(t, "Book added successfully", response["message"])
// }

// // func TestGetBooksByLibrary(t *testing.T) {
// // 	setupTestDB()
// // 	r := setupTestRouter()

// // 	req, _ := http.NewRequest("GET", "/books", nil)
// // 	w := httptest.NewRecorder()
// // 	r.ServeHTTP(w, req)

// // 	assert.Equal(t, http.StatusOK, w.Code)
// // }

// // func TestGetBookByISBN(t *testing.T) {
// // 	setupTestDB()
// // 	r := setupTestRouter()

// // 	req, _ := http.NewRequest("GET", "/books/123456789", nil)
// // 	w := httptest.NewRecorder()
// // 	r.ServeHTTP(w, req)

// // 	assert.Equal(t, http.StatusOK, w.Code)
// // }

// // func TestUpdateBook(t *testing.T) {
// // 	setupTestDB()
// // 	r := setupTestRouter()

// // 	updatedBook := models.Book{
// // 		Title: "Updated Title",
// // 	}
// // 	jsonValue, _ := json.Marshal(updatedBook)
// // 	req, _ := http.NewRequest("PUT", "/books/123456789", bytes.NewBuffer(jsonValue))
// // 	req.Header.Set("Content-Type", "application/json")

// // 	w := httptest.NewRecorder()
// // 	r.ServeHTTP(w, req)

// // 	assert.Equal(t, http.StatusOK, w.Code)
// // }

// // func TestDeleteBook(t *testing.T) {
// // 	setupTestDB()
// // 	r := setupTestRouter()

// // 	req, _ := http.NewRequest("DELETE", "/books/123456789", nil)
// // 	w := httptest.NewRecorder()
// // 	r.ServeHTTP(w, req)

// // 	assert.Equal(t, http.StatusOK, w.Code)
// // }
package tests
