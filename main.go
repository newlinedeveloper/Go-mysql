package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	http.Handler("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}