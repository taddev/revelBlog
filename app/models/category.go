package models

import (
	"github.com/revel/revel"
)

type ViewCategory struct {
	TotalRows uint64 `json:"total_rows"`
	Offset    uint64
	Rows      []ViewCategoryRow
}

type ViewCategoryRow struct {
	Id    string
	Key   string
	Value string
}

type Category struct {
	Id          string `json:"_id"`
	Rev         string `json:"_rev"`
	Name        string
	Description string
	Type        string
}

func (category *Category) Validate(v *revel.Validation) {
	v.Required(category.Name).Message("Name Required")
	v.Required(category.Description).Message("Description Required")
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
