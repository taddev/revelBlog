package models

//import

type Post struct {
	Id     string `json:"_id"`
	Rev    string `json:"_rev"`
	Header string //`json:"header"`
	Body   string //`json:"body"`
	Date   []int  //`json:"date"`
	Author string //`json:"author"`
}
