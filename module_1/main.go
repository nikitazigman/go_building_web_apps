package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	PORT   = ":8080"
	DBHost = "127.0.0.1"
	DBPort = 5432
	DBUser = "postgres"
	DBPass = "postgres"
	DBName = "postgres"
)

var database *sql.DB

type Page struct {
	Title   string
	Content string
	Date    string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageID := vars["id"]
	thisPage := Page{}
	fmt.Println(pageID)
	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE id = $1", pageID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)
	if err != nil {
		log.Println("Couldn't get page")
		http.Error(w, http.StatusText(404), http.StatusNotFound)
	}
	fmt.Fprintln(w, thisPage.Title+" "+thisPage.Content+" "+thisPage.Date)
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBHost, DBPort, DBUser, DBPass, DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	database = db
	defer database.Close()
	err = database.Ping()
	if err != nil {
		panic(err)
	}
	rtr := mux.NewRouter()
	rtr.HandleFunc("/pages/{id:[0-9]+}", ServePage)
	http.Handle("/", rtr)
	http.ListenAndServe(PORT, nil)
}
