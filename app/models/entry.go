package models

//import

type Entry struct {
	Id    string `json:"_id"`
	Rev   string `json:"_rev"` // useful only for Retrieve and Edit
	Entry int    `json:"entry"`
	Name  string `json:"name"`
}
