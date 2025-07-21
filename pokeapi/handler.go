package pokeapi

import (
	"encoding/json"
	"net/http"
)

func GetLocationAreas(url string) (locations []Location, next string, previous string, err error) {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", "", err
	}

	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return nil, "", "", err
	}

	var location_area LocationArea
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&location_area); err != nil {
		return nil, "", "", err
	}

	for _, result := range location_area.Results {
		locations = append(locations, result)
	}

	if location_area.Next == nil {
		next = ""
	} else {
		next = string(*location_area.Next)
	}

	if location_area.Previous == nil {
		previous = ""
	} else {
		previous = string(*location_area.Previous)
	}
	return locations, next, previous, nil
}
