package main

import (
	"groupie/handler"
	"net/http"
)

func main() {
	// aritst, _ := handler.GetRelationByID(2)
	// fmt.Println(aritst)

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/artist/", handler.ArtistHandler)
	http.ListenAndServe(":8090", nil)
}
