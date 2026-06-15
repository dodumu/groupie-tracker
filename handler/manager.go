package handler

import (
	"fmt"
	"strings"
)

func GetLocationByID(id int) (Location, error) {

	for _, location := range LocationsCache {
		if location.ID == id {
			return location, nil
		}
	}
	return Location{}, fmt.Errorf("location with ID %d not found", id)
}

func GetDatesByID(id int) (Date, error) {

	for _, date := range DatesCache {
		if date.ID == id {
			return date, nil
		}
	}
	return Date{}, fmt.Errorf("date with ID %d not found", id)
}

func GetRelationByID(id int) (Relation, error) {

	for _, relation := range RelationsCache {
		if relation.ID == id {
			return relation, nil
		}
	}
	return Relation{}, fmt.Errorf("relation with ID %d not found", id)
}

func GetArtistsByID(id int) (Artist, error) {

	for _, artist := range ArtistsCache {
		if artist.ID == id {
			return artist, nil
		}
	}
	return Artist{}, fmt.Errorf("artist with ID %d not found", id)
}

func GetArtistPage(id int) (ArtistPage, error) {
	artist, err := GetArtistsByID(id)
	if err != nil {
		return ArtistPage{}, err
	}

	location, err := GetLocationByID(id)
	if err != nil {
		return ArtistPage{}, err
	}
	date, err := GetDatesByID(id)
	if err != nil {
		return ArtistPage{}, err
	}
	relation, err := GetRelationByID(id)
	if err != nil {
		return ArtistPage{}, err
	}
	target := ArtistPage{
		Artist:    artist,
		Locations: location.Location,
		Dates:     date.Dates,
		Relations: relation.DatesLocation,
	}
	return target, nil
}

func SearchUsers(name string) (Artist, error) {
	band, err := GetArtists()
	if err != nil {
		return Artist{}, err
	}

	for _, singer := range band {
		if strings.ToLower(singer.Name) == name || strings.Contains(strings.ToLower(singer.Name), name) {
			return singer, nil
		}
	}
	return Artist{}, fmt.Errorf("Artist not found")
}
