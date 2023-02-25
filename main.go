package main

import (
	"catalog/cmd/api"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/Manga", api.CreateManga).Methods("POST")
	router.HandleFunc("/Mangas", api.SearchManga).Methods("GET")
	router.HandleFunc("/Manga/{id}", api.DeleteManga).Methods("DELETE")
	router.HandleFunc("/Manga/{id}", api.UpdateManga).Methods("PUT")
	router.HandleFunc("/Manga/{id}", api.GetMangaById).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
