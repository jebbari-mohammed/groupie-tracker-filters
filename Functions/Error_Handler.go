package functions

import (
	"html/template"
	"net/http"
)

var (
	Message string
	
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int) {
	switch code {
	case 400:
		Message = "Bad Request"
	case 403:
		Message = "Forbidden"
	case 404:
		Message = "NoT Found"
	case 405:
		Message = "Method Not Allowed "
	case 500:
		Message = "Internal Server Error"
	default:
		Message = "Error"
	}
	w.WriteHeader(code)
	tmpl,err:=template.ParseFiles("Template/Error.html")
	if err!=nil {
		w.WriteHeader(500)
		http.ServeFile(w,r,"Template/error.html")
		return
	}
	

	data := struct {
		Message string
		Code    int
	}{
		Message: Message,
		Code:    code,
	}

	err=tmpl.Execute(w, data)
	if err!=nil{
		w.WriteHeader(500)
		http.ServeFile(w,r,"Template/error.html")
		return
	}
}
