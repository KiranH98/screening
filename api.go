package main

import (
	"fmt"
	"net/http"
)

/* func setupJsonApi() {
	http.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection
		conn := createConnection()
		name := r.FormValue("name")
		email := r.FormValue("email")
		query := "INSERT INTO users (name, email) VALUES (" + name + ", " + email + ")"
		result, err := conn.Exec(query)
		fmt.Println("result ", result, " err ", err.Error())
		w.Write([]byte("Created user successfully!"))
	})
	http.HandleFunc("/updateUser", func(w http.ResponseWriter, r *http.Request) {
		// create mysql connection
		conn := createConnection()
		name := r.FormValue("name")
		email := r.FormValue("email")
		query := "Update users set name=" + name + ", email=" + email + " where id=" + r.FormValue("id")
		result, err := conn.Exec(query)
		fmt.Println("result ", result, " err ", err.Error())
		w.Write([]byte("User updated successfully!"))
	})
} */

func setupJSONAPI() {
	http.HandleFunc("/createUser", createUserHandler)
	http.HandleFunc("/updateUser", updateUserHandler)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	query := "INSERT INTO users (name, email) VALUES (?, ?)"

	_, err := db.Exec(query, name, email)
	handleDBOperationResult(w, err, "Created user successfully!")
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	id := r.FormValue("id")
	query := "UPDATE users SET name=?, email=? WHERE id=?"

	_, err := db.Exec(query, name, email, id)
	handleDBOperationResult(w, err, "User updated successfully!")
}

func handleDBOperationResult(w http.ResponseWriter, err error, successMessage string) {
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(successMessage))
}
