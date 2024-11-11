package functions

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func Search_Bar(w http.ResponseWriter, r *http.Request) {
	var Search_Artist []Artist
	if r.Method != "GET" {
		ErrorHandler(w, r, 405)
		return
	}
	er, err := Fitch_Global(w, r, Url_Artists), Fitch_Global(w, r, Url_Locations)
	if er != nil || err != nil {
		ErrorHandler(w, r, 500)
		return
	}
	Search := r.FormValue("search")
	low_Search := strings.ToLower(Search)
	if len(Search)==0{
		ErrorHandler(w,r,400)
		return

	}
	found := map[int]bool{}
	for _, artist := range Artists {
		if strings.Contains(strings.ToLower(artist.Name), low_Search) && !found[artist.Id] {
			Search_Artist = append(Search_Artist, artist)
			found[artist.Id] = true
		}
		for _ , member := range artist.Members{
			if strings.Contains(strings.ToLower(member),low_Search)&& !found[artist.Id]{
				Search_Artist=append(Search_Artist, artist)
				found[artist.Id]=true
			}
			if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)),low_Search)&& !found[artist.Id]{
				Search_Artist=append(Search_Artist, artist)
				found[artist.Id]=true
			}

		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), low_Search)&& !found[artist.Id]{
			Search_Artist=append(Search_Artist, artist)
			found[artist.Id]=true
		}
		for _ , location :=range Locations.Index{
			for _, loc := range location.Locatins{
				if strings.Contains(strings.ToLower(loc), low_Search) && !found[artist.Id]{
					if artist.Id!=location.Id{
						continue
					}
					Search_Artist=append(Search_Artist, artist)
					found[artist.Id]=true
				}

		
			}
		}
	}
	if len(Search_Artist)==0{
		ErrorHandler(w,r,404)
		return
	}
	tmpl, err := template.ParseFiles("Template/Search_Bar.html")
	if err!= nil {
		ErrorHandler(w,r,500)
		return
	}
	err=tmpl.Execute(w,Search_Artist)
	if err!=nil {
		ErrorHandler(w,r,500)
		return
	}
}
