package models

import (
	"github.com/robfig/revel"
)

type ViewHeader struct {
	TotalRows uint64 `json:"total_rows"`
	Offset    uint64
	Rows      []ViewHeaderRow
}

type ViewSummary struct {
	TotalRows uint64 `json:"total_rows`
	Offset    uint64
	Rows      []ViewSummaryRow
}

type ViewHeaderRow struct {
	Id    string //`json:"id"`
	Key   []int  //`json:"key"`
	Value string //`json:"value"`
}

type ViewSummaryRow struct {
	Id    string
	Key   []int
	Value []string
}

type Post struct {
	Id     string `json:"_id"`
	Rev    string `json:"_rev"`
	Header string //`json:"header"`
	Body   string //`json:"body"`
	Date   []int  //`json:"date"`
	Author string //`json:"author"`
	Type   string
}

func (post *Post) Validate(v *revel.Validation) {
	v.Required(post.Header).Message("Header Required")
	v.Required(post.Body).Message("Body Required")
	/*
	v.Required(user.FirstName)
	v.Required(user.LastName)
	v.Required(user.Age)
	v.Range(user.Age, 16, 120)
	v.Required(user.Password)
	v.MinSize(user.Password, 6)
	v.Required(user.PasswordConfirm)
	v.Required(user.PasswordConfirm == user.Password).
		Message("The passwords do not match.")
	v.Required(user.Email)
	v.Email(user.Email)
	v.Required(user.EmailConfirm)
	v.Required(user.EmailConfirm == user.Email).
		Message("The email addresses do not match")
	v.Required(user.TermsOfUse)
	*/
}
