package controllers

import (
	"strings"
	"time"

	goravelhttp "github.com/goravel/framework/contracts/http"

	"goravel/app/facades"
)

type AdminController struct{}

func NewAdminController() *AdminController { return &AdminController{} }

func (r *AdminController) LoginPage(ctx goravelhttp.Context) goravelhttp.Response {
	if ctx.Request().Session().Has("admin_authenticated") {
		return ctx.Response().Redirect(goravelhttp.StatusFound, "/admin/dashboard")
	}
	errorMessage, _ := ctx.Request().Session().Pull("admin_error", "").(string)
	return ctx.Response().View().Make("admin/login.tmpl", map[string]any{"error": errorMessage})
}

func (r *AdminController) Login(ctx goravelhttp.Context) goravelhttp.Response {
	username := strings.TrimSpace(ctx.Request().Input("username"))
	password := ctx.Request().Input("password")
	validUsername := facades.Config().GetString("admin.username")
	validPassword := facades.Config().GetString("admin.password")

	if username == "" || password == "" || username != validUsername || password != validPassword {
		ctx.Request().Session().Flash("admin_error", "Invalid username or password.")
		return ctx.Response().Redirect(goravelhttp.StatusFound, "/admin")
	}

	ctx.Request().Session().Put("admin_authenticated", true)
	ctx.Request().Session().Put("admin_username", username)
	ctx.Request().Session().Forget("admin_error")
	return ctx.Response().Redirect(goravelhttp.StatusFound, "/admin/dashboard")
}

func (r *AdminController) Dashboard(ctx goravelhttp.Context) goravelhttp.Response {
	if !ctx.Request().Session().Has("admin_authenticated") {
		return ctx.Response().Redirect(goravelhttp.StatusFound, "/admin")
	}

	enquiries, err := readEnquiries()
	if err != nil {
		ctx.Request().Session().Flash("admin_error", "Could not load enquiries.")
		return ctx.Response().Redirect(goravelhttp.StatusFound, "/admin")
	}
	today := 0
	todayPrefix := time.Now().Format("2006-01-02")
	for _, enquiry := range enquiries {
		if enquiry.CreatedAt != nil && enquiry.CreatedAt.Format("2006-01-02") == todayPrefix {
			today++
		}
	}

	return ctx.Response().View().Make("admin/dashboard.tmpl", map[string]any{
		"enquiries": enquiries,
		"count":     len(enquiries),
		"today":     today,
	})
}

func (r *AdminController) Logout(ctx goravelhttp.Context) goravelhttp.Response {
	_ = ctx.Request().Session().Invalidate()
	return ctx.Response().Redirect(goravelhttp.StatusFound, "/admin")
}
