package app

import (
	"github.com/robfig/revel"
	"github.com/taddevries/lazyboy"
	"github.com/taddevries/revelBlog/app/models"
	"time"
)

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.ActionInvoker,           // Invoke the action.
	}

	//timeFormat provides a means to configure the way time is displayed
	//this is configured as a custom template function so the way time
	//is displayed can be changed just by changing this function's output
	//each post stores it's time as a slice [yyyy,mm,dd,hh,mm,ss]
	//this function takes that format and converts it to a native Go Time object
	//the timezone information can also be set in the app.conf file or left to
	//it's default of UTC
	revel.TemplateFuncs["timeFormat"] = func(date []int) string {
		const layout = "_2 Oct 2006"
		z, _ := time.LoadLocation(revel.Config.StringDefault("timezone", "UTC"))
		t := time.Date(date[0], time.Month(date[1]), date[2], date[3], date[4], date[5], 0, z)
		return t.Format(layout)
	}

	revel.TemplateFuncs["getAuthor"] = func(userId string) string {
		user := models.User{}
		lazyboy.Database.Retrieve(userId, &user)

		return user.DisplayName
	}

	revel.TemplateFuncs["getCategory"] = func(categoryId string) string {
		category := models.Category{}
		lazyboy.Database.Retrieve(categoryId, &category)
		return category.Description
	}

	revel.OnAppStart(lazyboy.AppInit)
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}
