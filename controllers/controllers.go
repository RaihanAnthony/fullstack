package controllers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/Raihanpoke/fullstack/entities"
	"github.com/Raihanpoke/fullstack/libraries"
	"github.com/Raihanpoke/fullstack/models"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("view/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}

var userModel = models.NewUserModel()
var validation = libraries.NewValidation()

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("view/login.html")

		temp.Execute(w, nil)
	}
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("view/signup.html")
		temp.Execute(w, nil)

	} else if r.Method == http.MethodPost {
		// do proses register

		// mengambil inputan form
		r.ParseForm()

		user := entities.User{
			First_Name: r.Form.Get("first_name"),
			Last_Name:  r.Form.Get("last_name"),
			Email:      r.Form.Get("email"),
			Phone:      r.Form.Get("phone"),
			Password:   r.Form.Get("password"),
			Cpassword:  r.Form.Get("cpassword"),
		}

		errorMessages := validation.Struct(user)

		if errorMessages != nil {

			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       user,
			}

			temp, _ := template.ParseFiles("view/signup.html")
			temp.Execute(w, data)
		} else {

			user.ID = uuid.New().String()
			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			user.Password = string(hashPassword)
			user.Created_At = time.Now()
			user.Updated_At = time.Now()

			userModel.Create(user)

			data := map[string]interface{}{
				"pesan": "registrasi berhasil",
			}

			temp, _ := template.ParseFiles("view/signup.html")
			temp.Execute(w, data)
		}

		// errorMessages := validation.struct(user)

	}
}
