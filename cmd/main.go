package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"hospital-management-system/pkg/config"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()
	dbConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	port := os.Getenv("PORT") 
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, world!")
	})
	fmt.Println("Starting server on port:", port)
	http.ListenAndServe(":"+port, nil)
}