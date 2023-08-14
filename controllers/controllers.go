package controllers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("view/index.html")

	if err != nil {
		panic(err)
	}

	temp.Execute(w, r)
}

func SignUp(w http.ResponseWriter, r http.Request) {

}
