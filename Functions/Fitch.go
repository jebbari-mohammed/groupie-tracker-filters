package functions

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)
type Search struct{
	Name string `json:"name"`
	Members []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Location string `json:"locations"` 

}

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}
type Location struct {
	Id       int      `json:"id"`
	Locatins []string `json:"locations"`
}
type Date struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}
type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var (
	Url_Artists   = "https://groupietrackers.herokuapp.com/api/artists"
	Url_Locations = "https://groupietrackers.herokuapp.com/api/locations"
	Url_dates     = "https://groupietrackers.herokuapp.com/api/dates"
	Url_Relations = "https://groupietrackers.herokuapp.com/api/relation"
	Artists       []Artist
	Locations     struct {
		Index []Location `json:"index"`
	}
)

var Dates struct {
	Index []Date `json:"index"`
}
var Relations struct {
	Index []Relation `json:"index"`
}

func Fitch_Global(w http.ResponseWriter, r *http.Request, pattern string) error {
	URL := ""

	switch pattern {
	case Url_Artists:
		URL = Url_Artists

	case Url_Locations:
		URL = Url_Locations

	case Url_dates:
		URL = Url_dates

	case Url_Relations:
		URL = Url_Relations

	}
	resp, error := http.Get(URL)
	if error != nil {
		ErrorHandler(w, r, 500)
		return error
	}
	defer resp.Body.Close()

	body, error := io.ReadAll(resp.Body)
	if error != nil {
		log.Fatal("Error Reading !!!")
	}
	if pattern == Url_Artists {
		err := json.Unmarshal(body, &Artists)
		if err != nil {
			log.Fatalf("Erreur de désérialisation : %v", err)
		}
	} else if pattern == Url_Locations {
		err := json.Unmarshal(body, &Locations)
		if err != nil {
			log.Fatalf("Erreur de désérialisation : %v", err)
		}
	} else if pattern == Url_dates {
		err := json.Unmarshal(body, &Dates)
		if err != nil {
			log.Fatalf("Erreur de désérialisation : %v", err)
		}
	} else {
		err := json.Unmarshal(body, &Relations)
		if err != nil {
			log.Fatalf("Erreur de désérialisation : %v", err)
		}

	}
	return nil
}

//func Fitch_Location(w http.ResponseWriter, r *http.Request) error {
//resp, error := http.Get("https://groupietrackers.herokuapp.com/api/locations")
//if error != nil {
//ErrorHandler(w, r, 400)
//return error
//}
//defer resp.Body.Close()
//body, error := io.ReadAll(resp.Body)
//if error != nil {
//	log.Fatal("Error Reading !!!")
//}
//err := json.Unmarshal(body, &Locations)
//if err != nil {
//	log.Fatalf("Erreur de désérialisation : %v", err)
//}

//return nil
//}

//func Fitch_Date(w http.ResponseWriter, r *http.Request) error {
//resp, error := http.Get("https://groupietrackers.herokuapp.com/api/dates")
//if error != nil {
//ErrorHandler(w, r, 400)
//return error
//}
///defer resp.Body.Close()
///body, error := io.ReadAll(resp.Body)
//if error != nil {
//	log.Fatal("Error Reading !!!")
//}

//return nil
//}

//func Fitch_Relation(w http.ResponseWriter, r *http.Request) error {
//resp, error := http.Get("https://groupietrackers.herokuapp.com/api/relation")
//if error != nil {
//	ErrorHandler(w, r, 400)
//	return error
//}
///defer resp.Body.Close()
//body, error := io.ReadAll(resp.Body)
//if error != nil {
//	log.Fatal("Error Reading !!!")
///}

//return nil
//}
