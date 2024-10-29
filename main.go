package main

import (
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

func main() {

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
