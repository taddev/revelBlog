package models

//import

type User struct {
	Id          string `json:"_id"`
	Rev         string `json:"_rev"`
	Username    string 
	DisplayName string 
	Password    string 
}
