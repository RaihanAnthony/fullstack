package entities

import "time"

type User struct {
	ID            string
	First_Name    string
	Last_Name     string
	Email         string
	Password      string
	Phone         string
	Token         string
	Refresh_Token string
	Created_Token time.Time
	Created_At    time.Time
	Updated_At    time.Time
	User_ID       string
	Cpassword     string
}
