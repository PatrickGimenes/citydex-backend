package model

import "time"

type User struct {
	Id        string    `json:"id_user"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Home_city string    `json:"home_city"`
	Create_At time.Time `json:"create_At"`
	Update_At time.Time `json:"update_At"`
}

type UserResponse struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Home_city string    `json:"home_city"`
	Create_At time.Time `json:"create_At"`
	Update_At time.Time `json:"update_At"`
}
