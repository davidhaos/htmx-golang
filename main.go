package main

import (
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {

	// mysql
	// replace with implementation of DBConnection interface
	// high level, stable components such as this, should use interfaces
	// and abstract classes. review wellnesswave construction
	// this should be in a different module that constructs the object.

	// your policy/business logic should be so decoupled from devices/details/db it doesn't even know what they are
	// different module constructs db object because main package doesn't know what will be inserted
	mysqlConn := &MySQLConnection{dsn: "user:password@tcp(localhost:3306)/dbname"}

	// Open connection
	if err := mysqlConn.Open(); err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}
	defer mysqlConn.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8080",
		// Enforce timeouts
		//WriteTimeout: 15 * time.Second,
		//ReadTimeout: 15 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		fmt.Println("error starting server: ", err)
	}
	//	fmt.Println("Hello")
}
