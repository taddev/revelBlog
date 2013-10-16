package controllers

import (
	"code.google.com/p/go.crypto/bcrypt"
	//"fmt"
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
	"github.com/taddevries/revelBlog/app/models"
	"strings"
	"time"
)

type Admin struct {
	App
}

func (c Admin) CheckUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Index)
	}
	return nil
}

func (c Admin) GetNewPost() revel.Result {
	categories := models.ViewCategory{}
	opts := make(map[string]interface{})
	lazyboy.Database.Query("_design/blog/_view/category", opts, &categories)

	return c.Render(categories)
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
	post.Author = user.Id
	post.Type = "post"

	lazyboy.Database.Insert(post)

	return c.Redirect(App.Index)
}

func (c Admin) GetSettings() revel.Result {
	return c.Render()
}

func (c Admin) PostSettings(user models.User, confirmPassword, oldPassword string) revel.Result {
	//get the current settings from the database
	tempuser := c.getUser(user.Id)

	if oldPassword != "" {
		err := bcrypt.CompareHashAndPassword([]byte(tempuser.Password), []byte(oldPassword))
		if err != nil {
			c.Flash.Error("Old Password Incorrect!")
			return c.Redirect(Admin.GetSettings)
		}
	}

	//validate everything
	user.Validate(c.Validation, confirmPassword)

	// Handle errors
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		c.Flash.Error("Missing Things!")
		return c.Redirect(Admin.GetSettings)
	}

	//if there is a new password, encrypt it and store it
	//otherwise get the current password so it can be reinserted
	if user.Password != "" {
		bcryptPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(bcryptPassword)
	} else {
		user.Password = tempuser.Password
	}

	//perform the update, if there is an error we're probably out of
	//sync so send an error message to the page
	_, err := lazyboy.Database.Edit(&user)
	if err != nil {
		// edit failed, out of sync?
		c.FlashParams()
		c.Flash.Error("Edit Failed, out of sync?")
		return c.Redirect(Admin.GetSettings)
	}

	//If we got this far everything has worked so we'll set a success
	//message and send the page back to settings to display the
	//updated information
	c.Flash.Success("Edit Successful")
	return c.Redirect(Admin.GetSettings)
}
