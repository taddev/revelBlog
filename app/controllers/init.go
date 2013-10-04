package controllers

import (
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
)

func init() {
	revel.OnAppStart(lazyboy.AppInit)
	//revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
	//revel.InterceptMethod(Application.AddUser, revel.BEFORE)
	//revel.InterceptMethod(Hotels.checkUser, revel.BEFORE)
	//revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
	//revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}
