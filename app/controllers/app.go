package controllers

import (
	//"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
	"github.com/taddevries/revelBlog/app/models"
	"github.com/taddevries/revelBlog/app/routes"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	result := models.View{}
	opts := make(map[string]interface{})
	opts["limit"] = 10
	opts["descending"] = true

	lazyboy.Database.Query("_design/blog/_view/header", opts, &result)
	return c.Render(result)
}

func (c App) GetPost(id string) revel.Result {
	result := models.Post{}

	lazyboy.Database.Retrieve(id, &result)

	//t := time.Date(result.Date[0], result.Date[1], result.Date[2], result.Date[3], result.Date[4], result.Date[5], 0, time.MDT)

	postTime := fmt.Sprintf("%02d:%02d", result.Date[3], result.Date[4])
	postDate := fmt.Sprintf("%d/%02d/%02d", result.Date[0], result.Date[1], result.Date[2])

	return c.Render(result, postTime, postDate)
}

func (c App) NewPost() revel.Result {
	c.checkUser()
	return c.Render()
}

func (c App) SaveNewPost(post models.Post) revel.Result {

	return c.Render(post)
}

func (c App) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	/*
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	*/
	return nil
}

func (c App) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Login First")
		return c.Redirect(routes.App.Index())
	}
	return nil
}

func (c App) GetLogin() revel.Result {
	return c.Render()
}

func (c App) PostLogin(username, password string) revel.Result {
	/*
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(routes.App.Index())
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	*/
	return c.Render(routes.App.Index())
}