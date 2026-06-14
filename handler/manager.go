package handler

import "fmt"

func GetLocationByID(id int) (Location, error) {

	locations, err := GetLocations()
	if err != nil {
		return Location{}, err
	}
	for _, location := range locations {
		if location.ID == id {
			return location, nil
		}
	}
	return Location{}, fmt.Errorf("location with ID %d not found", id)
}

func GetDatesByID(id int) (Date, error) {
	dates, err := GetDate()
	if err != nil {
		return Date{}, err
	}
	for _, date := range dates {
		if date.ID == id {
			return date, nil
		}
	}
	return Date{}, fmt.Errorf("date with ID %d not found", id)
}

func GetRelationByID(id int) (Relation, error) {
	relations, err := GetRelations()
	if err != nil {
		return Relation{}, err
	}
	for _, relation := range relations {
		if relation.ID == id {
			return relation, nil
		}
	}
	return Relation{}, fmt.Errorf("relation with ID %d not found", id)
}

func GetArtistsByID(id int) (Artist, error) {
	artsits, err := GetArtists()
	if err != nil {
		return Artist{}, nil
	}
	for _, artist := range artsits {
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
