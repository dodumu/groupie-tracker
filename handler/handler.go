package handler

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func BaseHandler(w http.ResponseWriter, page string, data any) {
	tmpl, err := template.ParseFiles(
		"templates/base.html",
		"templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	artists, err := GetArtists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	page := HomePage{
		Artist: artists,
	}

	BaseHandler(w, "home.html", page)
}

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = strings.TrimPrefix(path, "/artist/")
	if path == "" {
		http.Error(w, "missing artist id", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	target, err := GetArtistPage(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	BaseHandler(w, "artists.html", target)
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("search")
	if query == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	artist, err := SearchUsers(strings.ToLower(query))
	if err != nil {
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}

	// reuse ArtistPage — get the full page for that artist
	target, err := GetArtistPage(artist.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	BaseHandler(w, "artists.html", target)
}
