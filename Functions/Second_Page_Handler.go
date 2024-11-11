package functions

import (
	
	"html/template"
	"net/http"
	"strconv"
)

type data struct {
	Artist   Artist
	Location []string
	Date     []string
	Relation map[string][]string
}

func Second_Page_Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, 405)
		return
	}
	er, err, errr, errrr := Fitch_Global(w, r,Url_Artists), Fitch_Global(w, r,Url_Locations), Fitch_Global(w, r,Url_dates), Fitch_Global(w, r,Url_Relations)
	if er != nil || err != nil || errr != nil || errrr != nil {
		return
	}
	id := r.FormValue("id")
	idd, _ := strconv.Atoi(id)
	

	i := false
	var Select_Artist Artist
	for _, Artist := range Artists {
		if idd == Artist.Id {

			Select_Artist = Artist
			i = true
			break
		}
	}
	if !i {
		
		ErrorHandler(w,r,404)
		return
	}
	var Select_Location Location
	for _, location := range Locations.Index {
		if idd == location.Id {
			Select_Location = location
			break
		}
	}
	var Select_Date Date
	for _, Date := range Dates.Index {
		if idd == Date.Id {
			Select_Date = Date
			break
		}
	}
	var Select_Relation Relation
	for _, Relation := range Relations.Index {
		if idd == Relation.Id {
			Select_Relation = Relation
			break
		}
	}
	result := data{
		Artist:   Select_Artist,
		Location: Select_Location.Locatins,
		Date:     Select_Date.Dates,
		Relation: Select_Relation.DatesLocations,
	}
	
	tmp, err := template.ParseFiles("Template/artist.html")
	if err != nil {
		ErrorHandler(w,r,500)
		return
	}
	err = tmp.Execute(w, result)
	if err != nil {
		ErrorHandler(w,r,500)
		return
	}
}
