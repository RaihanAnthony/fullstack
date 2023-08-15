package models

import "time"

type user struct {
	ID            string    `gorm:"primaryKey;autoIncrement;type:VARCHAR(199)"`
	First_Name    *string   `gorm:"column:first_name;type:VARCHAR(100)"`
	Last_Name     *string   `gorm:"column:last_name;type:VARCHAR(100)"`
	Email         *string   `gorm:"column:email;type:VARCHAR(100)"`
	Phone         *string   `gorm:"column:Phone;type:VARCHAR(100)"`
	Token         *string   `gorm:"column:token;type:VARCHAR(100)"`
	Refresh_Token *string   `gorm:"column:refresh_token;type:VARCHAR(100)"`
	Created_Token time.Time `json:"created_token" gorm:"column:created_token;type:TIMESTAMP"`
	Created_At    time.Time `json:"created_at" gorm:"column:created_at;type:TIMESTAMP"`
	Updated_At    time.Time `json:"update_at" gorm:"column:update_at;type:TIMESTAMP"`
	User_ID       *string   `json:"user_id" gorm:"column:user_id;type:VARCHAR(100)"`
}
