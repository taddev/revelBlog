package controllers

import (
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
	"github.com/taddevries/revelBlog/app/models"
)

type App struct {
	*revel.Controller
	lazyboy.CouchDBController
}

func (c App) Index() revel.Result {
	//temp := c.DBUrl
	id := "616edbad6650d9d1acb68b3157017c82"
	record := models.Entry{}
	c.Database.Retrieve(id, &record)
	return c.Render(id, record)
}
