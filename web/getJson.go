package groupie

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var Groups = make([]Groupie, 0)

func ParseData() {
	log.SetFlags(log.Lshortfile)
	url := "https://groupietrackers.herokuapp.com/api/artists"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("1235")

	err = json.Unmarshal(jsonData, &Groups)
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("dfjhsdfsdf")
	getLocation()
	getRelations()
	getConcertDates()
}

func getLocation() {
	type t struct {
		Index []Location `json:"index"`
	}

	temp := t{
		Index: make([]Location, 0),
	}

	url := "https://groupietrackers.herokuapp.com/api/locations"

	idData, err := getData(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(idData, &temp)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(Groups) != len(temp.Index) {
		log.Fatal("error: lengths of locations and groups are different")
		return
	}
	for id := range Groups {
		Groups[id].Locations = temp.Index[id]
	}
}

func getRelations() {
	type t struct {
		Index []Relation `json:"index"`
	}

	temp := t{
		Index: make([]Relation, 0),
	}

	url := "https://groupietrackers.herokuapp.com/api/relation"

	idData, err := getData(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = json.Unmarshal(idData, &temp)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(Groups) != len(temp.Index) {
		log.Fatal("error: lengths of locations and groups are different")
		return
	}
	for id := range Groups {
		Groups[id].Relations = temp.Index[id]
	}
}

func getConcertDates() {
	type t struct {
		Index []Date `json:"index"`
	}

	temp := t{
		Index: make([]Date, 0),
	}

	url := "https://groupietrackers.herokuapp.com/api/dates"

	idData, err := getData(url)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(idData, &temp)
	if err != nil {
		log.Fatal(err)
		return
	}
	if len(Groups) != len(temp.Index) {
		log.Fatal("error: lengths of locations and groups are different")
		return
	}
	for id := range Groups {
		Groups[id].ConcertDates = temp.Index[id]
	}

}

func getData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	idData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return idData, nil
}
