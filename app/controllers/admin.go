package controllers

import (
	//"code.google.com/p/go.crypto/bcrypt"
	//"fmt"
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
	"github.com/taddevries/revelBlog/app/models"
	"github.com/taddevries/revelBlog/app/routes"
	"strings"
	"time"
)

type Admin struct {
	App
}

func (c Admin) CheckUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(routes.App.Index())
	}
	return nil
}

func (c Admin) GetNewPost() revel.Result {

	return c.Render()
}

func (c Admin) PostNewPost(post models.Post) revel.Result {
	//validate everything
	post.Validate(c.Validation)

	// Handle errors
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Error("Missing Things!")
		return c.Redirect(Admin.GetNewPost)
	}

	t := time.Now()
	user := c.getUser(c.Session["Id"])

	//setup id based on header but only the leters and numbers
	charOnly := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z':
			return r
		case r >= 'a' && r <= 'z':
			return r
		case r >= '0' && r <= '9':
			return r
		case r == '-':
			return r
		}
		return -1
	}

	post.Id = (strings.Map(charOnly, strings.ToLower(strings.Join(strings.Fields(post.Header), "-"))))
	post.Date = []int{t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second()}
	post.Author = user.DisplayName
	post.Type = "post"

	lazyboy.Database.Insert(post)

	return c.Redirect(routes.App.Index())
}
