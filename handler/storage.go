package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	artist := []Artist{}
	err = json.NewDecoder(resp.Body).Decode(&artist)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	if err != nil {
		return artist, err
	}
	defer resp.Body.Close()
	return artist, nil
}

func GetLocations() ([]Location, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	var response LocationResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response.Index, nil
}

func GetDate() ([]Date, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	var response DatesResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response.Index, nil
}

func GetRelations() ([]Relation, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code %d", resp.StatusCode)
	}
	var response RelationResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}
	return response.Index, nil
}
