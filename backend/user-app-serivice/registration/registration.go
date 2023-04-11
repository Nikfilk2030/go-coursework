package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// func main, port := "3003"
func main() {
	port := "3003"

	// Connect to the PostgreSQL database
	db, err := sql.Open("postgres", "host=localhost port=3001 user=admin password=root dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Check the connection to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var user User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&user)
		if err != nil {
			panic(err)
		}

		fmt.Println(user.Login, user.Password) // tests

		// Hash the user's password before saving it
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			panic(err)
		}

		// Insert user data into the "users" table
		sqlStatement := `
			INSERT INTO users (login, password)
			VALUES ($1, $2)`
		_, err = db.Exec(sqlStatement, user.Login, hashedPassword)
		if err != nil {
			panic(err)
		}

		fmt.Println(user.Login, hashedPassword)
		w.Header().Set("Content-Type", "application/json")
	})

	// program listens on port
	http.ListenAndServe(":"+port, nil)
}
