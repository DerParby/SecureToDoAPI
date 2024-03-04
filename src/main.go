package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 8080
	user     = "your_username"
	password = "your_password"
	dbname   = "your_database_name"
)

var db *sql.DB

func loginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	var storedPassword string
	err := db.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Fprintf(w, "Benutzer nicht gefunden")
			return
		}
		log.Fatal(err)
	}

	// Überprüfe das eingegebene Passwort mit dem gespeicherten Passwort
	if password == storedPassword {
		fmt.Fprintf(w, "Erfolgreich eingeloggt")
	} else {
		fmt.Fprintf(w, "Falsches Passwort")
	}
}

func main() {
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/login", loginHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
