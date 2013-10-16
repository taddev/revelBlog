package models

import (
	"github.com/robfig/revel"
)

type ViewUser struct {
	TotalRows uint64 `json:"total_rows"`
	Offset    uint64
	Rows      []ViewUserRow
}

type ViewUserRow struct {
	Id    string
	Key   string
	Value []string
}

type User struct {
	Id          string `json:"_id"`
	Rev         string `json:"_rev"`
	Username    string
	DisplayName string
	Password    string
}

func (user *User) Validate(v *revel.Validation, confirmPassword string) {
	v.Required(user.Username).Message("Username Reuired")
	v.MinSize(user.Username, 3)
	v.Required(user.DisplayName).Message("Display Name Required")
	v.Required(confirmPassword == user.Password).Message("The passwords do not match")
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
