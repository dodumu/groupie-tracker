package main

import (
	"groupie/handler"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static", http.StripPrefix("/static/", fs))
	// aritst, _ := handler.GetRelationByID(2)
	// fmt.Println(aritst)

	err := handler.LoadData()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/artist/", handler.ArtistHandler)
	http.HandleFunc("/search", handler.SearchHandler)
	http.ListenAndServe(":8090", nil)
}
