package entities

import "time"

type User struct {
	ID            string
	First_Name    string `validate:"required,min=4,max=20" label:"First Name"`
	Last_Name     string `validate:"required,min=4,max=20" label:"Last Name"`
	Email         string `validate:"required,email,isunique=users-email"`
	Password      string `validate:"required,gte=6"`
	Phone         string `validate:"required,isunique=users-phone"`
	Token         string
	Refresh_Token string
	Created_Token time.Time
	Created_At    time.Time
	Updated_At    time.Time
	Cpassword     string `validate:"required,eqfield=Password" label:"konfirmasi password"`
}
