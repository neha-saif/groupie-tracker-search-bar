package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"	
)


type Data struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Artist       string   `json:"name"`
	Members      []string `json:"members"`
	AlbumYear    int      `json:"creationDate"`
	Album1       string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	RelUrl       string   `json:"relations"`
}

type Urelles struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func LoadData(Url string) ([]Data, error) {
	response, err := http.Get(Url)
	if err != nil {
		fmt.Println("Error getting response from url")
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error cannot store get resposne in body")
	}

	var character []Data
	err = json.Unmarshal(body, &character)
	if err != nil {
		fmt.Print("Error storing body in address of character")
	}

	return character, nil
}

func LoadUrelles(Url string) ([]Urelles, error) {
	response, err := http.Get(Url)
	if err != nil {
		fmt.Println("Error getting response from url")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("error reading response body")
	}

	// Unmarshal the data into a map of Urelles to extract index
	var data map[string][]Urelles
	errr := json.Unmarshal(body, &data)
	if errr != nil { 
		fmt.Print("Error storing body in address of character(unnmarshal issue)")
	}

	// Extract index from the data map which is the type urelles struct
	// ok is the boolean whhich indicates which checks if index exists
	urelles, ok := data["index"]
	if !ok {
	 fmt.Println("error: key 'index' not found in API response")
	}

	return urelles, nil
}
