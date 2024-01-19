package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Inti DB
	initDB()
	code := m.Run()
	// Close connections after test ends
	os.Exit(code)
}

// unittest to check insert users
func TestCreateUserHandler(t *testing.T) {
	clearDatabase()

	// Create a test HTTP request
	req, err := http.NewRequest("POST", "/createUser", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Input
	q := req.URL.Query()
	q.Add("name", "John Doe")
	q.Add("email", "john.doe@example.com")
	req.URL.RawQuery = q.Encode()

	// Create a test HTTP response recorder
	w := httptest.NewRecorder()

	// Call the API handler
	createUserHandler(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Query the database to verify the user was created
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

// Unit test method to check update user details
func TestUpdateUserHandler(t *testing.T) {
	clearDatabase()

	// Create a test HTTP request
	req, err := http.NewRequest("PUT", "/updateUser?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Set form values for the request
	q := req.URL.Query()
	q.Add("name", "Updated Name")
	q.Add("email", "updated.email@example.com")
	req.URL.RawQuery = q.Encode()

	// Create a test HTTP response recorder
	w := httptest.NewRecorder()

	// Call the API handler
	updateUserHandler(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Optionally, you can assert the response body or perform additional checks
	// ...

	// Query the database to verify the user was updated
	var name, email string
	err = db.QueryRow("SELECT name, email FROM users WHERE id = 1").Scan(&name, &email)
	assert.NoError(t, err)
	assert.Equal(t, "Updated Name", name)
	assert.Equal(t, "updated.email@example.com", email)
}

func clearDatabase() {
	// Use a testing database or a transaction to isolate tests and avoid interference
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		panic(err)
	}
}
