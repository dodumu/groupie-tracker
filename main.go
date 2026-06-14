package main

import (
	"groupie/handler"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static", http.StripPrefix("/static/", fs))
	// aritst, _ := handler.GetRelationByID(2)
	// fmt.Println(aritst)

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/artist/", handler.ArtistHandler)
	http.ListenAndServe(":8090", nil)
}
