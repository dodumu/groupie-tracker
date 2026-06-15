package handler

var (
	ArtistsCache   []Artist
	LocationsCache []Location
	DatesCache     []Date
	RelationsCache []Relation
)

func LoadArtists() error {
	artists, err := GetArtists()
	if err != nil {
		return err
	}

	ArtistsCache = artists
	return nil
}

func LoadLocations() error {
	locations, err := GetLocations()
	if err != nil {
		return err
	}

	LocationsCache = locations
	return nil
}

func LoadDates() error {
	dates, err := GetDate()
	if err != nil {
		return err
	}

	DatesCache = dates
	return nil
}

func LoadRelations() error {
	relations, err := GetRelations()
	if err != nil {
		return err
	}

	RelationsCache = relations
	return nil
}

func LoadData() error {
	if err := LoadArtists(); err != nil {
		return err
	}

	if err := LoadLocations(); err != nil {
		return err
	}

	if err := LoadDates(); err != nil {
		return err
	}

	if err := LoadRelations(); err != nil {
		return err
	}

	return nil
}
