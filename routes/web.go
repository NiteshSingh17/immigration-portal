package routes

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/support"

	"goravel/app/facades"
	"goravel/app/http/controllers"
)

func Web() {
	facades.Route().Get("/", func(ctx http.Context) http.Response {
		success, _ := ctx.Request().Session().Pull("enquiry_success", "").(string)
		errorMessage, _ := ctx.Request().Session().Pull("enquiry_error", "").(string)
		return ctx.Response().View().Make("welcome.tmpl", map[string]any{
			"version":         support.Version,
			"enquiry_success": success,
			"enquiry_error":   errorMessage,
		})
	})

	facades.Route().Static("public", "./public")

	enquiryController := controllers.NewEnquiryController()
	facades.Route().Post("/contact/enquiries", enquiryController.Store)

	adminController := controllers.NewAdminController()
	facades.Route().Get("/admin", adminController.LoginPage)
	facades.Route().Post("/admin/login", adminController.Login)
	facades.Route().Get("/admin/dashboard", adminController.Dashboard)
	facades.Route().Post("/admin/logout", adminController.Logout)

	userController := controllers.NewUserController()
	facades.Route().Get("/users", userController.Index)

	facades.Route().Get("/privacy", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("privacy.tmpl")
	})

	facades.Route().Get("/terms", func(ctx http.Context) http.Response {
		return ctx.Response().View().Make("terms.tmpl")
	})
}
