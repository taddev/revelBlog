package controllers

import (
	"github.com/revel/revel"
)

func init() {
	revel.InterceptMethod(Admin.CheckUser, revel.BEFORE)
	revel.InterceptMethod(App.AddUser, revel.BEFORE)
}
