package controllers

import (
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
	"github.com/taddevries/revelBlog/app/models"
	//"github.com/taddevries/revelBlog/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	result := models.View{}
	opts := make(map[string]interface{})
	opts["limit"] = 10

	lazyboy.Database.Query("_design/blog/_view/header", opts, &result)
	return c.Render(result)
}

func (c App) GetPost(id string) revel.Result {
	result := models.Post{}

	lazyboy.Database.Retrieve(id, &result)
	return c.Render(result)
}
