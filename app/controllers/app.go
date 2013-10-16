package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
	"github.com/taddevries/revelBlog/app/models"
)

type App struct {
	*revel.Controller
}

func (c App) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c App) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if id, ok := c.Session["Id"]; ok {
		return c.getUser(id)
	}
	return nil
}

func (c App) getUser(id string) *models.User {
	user := models.User{}
	lazyboy.Database.Retrieve(id, &user)

	return &user
}

func (c App) Index() revel.Result {
	result := models.ViewSummary{}
	opts := make(map[string]interface{})
	opts["limit"] = 10
	opts["descending"] = true

	lazyboy.Database.Query("_design/blog/_view/summary", opts, &result)
	rowCount := len(result.Rows)
	return c.Render(result, rowCount)
}

func (c App) GetPost(id string) revel.Result {
	result := models.Post{}
	lazyboy.Database.Retrieve(id, &result)
	return c.Render(result)
}

func (c App) GetLogin() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
		return c.Redirect(App.Index)
	}
	return c.Render()
}

func (c App) PostLogin(user models.User) revel.Result {
	result := models.ViewUser{}
	opts := make(map[string]interface{})
	opts["limit"] = 1
	opts["key"] = user.Username

	lazyboy.Database.Query("_design/admin/_view/user", opts, &result)

	if len(result.Rows) != 0 {
		err := bcrypt.CompareHashAndPassword([]byte(result.Rows[0].Value[0]), []byte(user.Password))
		if err == nil {
			c.Session["Id"] = result.Rows[0].Id
			c.Flash.Success("Login Success!")
			return c.Redirect(App.Index)
		}
	}

	c.Flash.Out["Username"] = user.Username
	c.Flash.Error("Login Failed")
	return c.Redirect(App.Index)
}

func (c App) GetLogout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	c.Flash.Success("Logout Successful!")
	return c.Redirect(App.Index)
}

func (c App) GetCategory(category string) revel.Result {
	return c.Render()
}

/*
func (c App) AddUser() revel.Result {
	bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte("demo"), bcrypt.DefaultCost)
	stringPassword := string(bcryptPassword)
	demoUser := models.User{Id:"", Rev:"", Username:"demo", DisplayName:"Demo User", Password:stringPassword}

	lazyboy.Database.Insert(demoUser)
	c.Flash.Success("Demo User Added")
	return c.Redirect(routes.App.Index())
}
*/
