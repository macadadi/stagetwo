package model

import "time"

type User struct {
	Id           int64     `json:"id"`
	Full_name    string    `json:"full_name"`
	Created_at   time.Time `json:"created_at"`
	Country_code int       `json:"country_code"`
}
