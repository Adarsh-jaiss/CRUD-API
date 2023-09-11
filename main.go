package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// "os/user"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// opening a DB
	err := OpenDB()
	if err != nil {
		log.Printf("Database not found %v\n", err)
	}

	//closing a db
	defer DB.Close()

	CreateTable()
}

func CreateTable() {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS USERS(id SERIAL PRIMARY KEY NOT NULL, name TEXT NOT NULL,email TEXT )")
	if err != nil {
		log.Fatal(err)
	}

	// seeting up a router
	router := mux.NewRouter()
	router.HandleFunc("/users", getAllUsers(DB)).Methods("GET")
	router.HandleFunc("/users/{id}", getUsers(DB)).Methods("GET")
	router.HandleFunc("/users", createUser(DB)).Methods("POST")
	router.HandleFunc("/users/{id}", updateUser(DB)).Methods("PUT")
	router.HandleFunc("/users/{id}", DeleteUser(DB)).Methods("DELETE")

	//START SERVER
	fmt.Println("starting a new server at port 10000")
	log.Fatal(http.ListenAndServe(":10000", middleware(router)))
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getAllUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		users := []User{}
		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(users)
	}
}

func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)

		vars := mux.Vars(r)
		id := vars["id"]

		err := db.QueryRow("SELECT * FROM users WHERE id= $1 ", id).Scan(&u, id, &u.Name, &u.Email)
		if err != nil {
			// TODO error handaling
			panic(err)
		}

		json.NewEncoder(w).Encode(u)

	}

}

func createUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)

		var id int
		err := db.QueryRow("INSERT INTO USERS (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&id)
		if err != nil {
			log.Fatal(err)
		}

		u.ID = id // Assign the inserted ID to the User struct
		json.NewEncoder(w).Encode(u)
	}
}

func updateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User
		json.NewDecoder(r.Body).Decode(&u)
		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Exec("UPDATE users SET name = $1,email = $2 WHERE id = $3", u.Name, u.Email, id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(u)
	}
}

func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		_, err := db.Query("DELETE FROM users WHERE id= $1", id)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode("USER Deleted")

	}
}
