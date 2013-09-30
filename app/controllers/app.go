package controllers

import ( 
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
)

type App struct {
	*revel.Controller
	lazyboy.CouchDBController
}

func (c App) Index() revel.Result {
	temp := c.DBName
	return c.Render(temp)
}
