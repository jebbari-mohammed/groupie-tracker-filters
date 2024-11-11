package functions

import (
	"html/template"
	"net/http"

	
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, 405)
		return
	}

	if r.URL.Path != "/" {
		ErrorHandler(w, r, 404)
		return
	}
	Error, errorr:= Fitch_Global(w, r , Url_Artists), Fitch_Global(w,r,Url_Locations)
	if Error != nil || errorr!=nil {
		ErrorHandler(w,r,500)
		return
	}
	Data_Result:=struct{
		Artists [] Artist
		Location []Location


	}{
		Artists: Artists,
		Location: Locations.Index,


	}
	tmpl, error := template.ParseFiles("Template/index.html")
	if error != nil {
		ErrorHandler(w, r, 500)
		return
	}

	error =tmpl.Execute(w, Data_Result)
	if error!=nil{
		ErrorHandler(w,r,500)
		return
	}
}
