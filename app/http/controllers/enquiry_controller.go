package controllers

import (
	"strings"

	goravelhttp "github.com/goravel/framework/contracts/http"

	"goravel/app/facades"
	"goravel/app/models"
)

type EnquiryController struct{}

func NewEnquiryController() *EnquiryController { return &EnquiryController{} }

func readEnquiries() ([]models.Enquiry, error) {
	var enquiries []models.Enquiry
	if err := facades.Orm().Query().Model(&models.Enquiry{}).OrderByDesc("created_at").Get(&enquiries); err != nil {
		return nil, err
	}
	return enquiries, nil
}

func (r *EnquiryController) Store(ctx goravelhttp.Context) goravelhttp.Response {
	name := strings.TrimSpace(ctx.Request().Input("name"))
	email := strings.TrimSpace(ctx.Request().Input("email"))
	phone := strings.TrimSpace(ctx.Request().Input("phone"))
	service := strings.TrimSpace(ctx.Request().Input("service"))
	message := strings.TrimSpace(ctx.Request().Input("message"))

	if name == "" || email == "" || phone == "" || message == "" {
		ctx.Request().Session().Flash("enquiry_error", "Please complete all required fields.")
		return ctx.Response().Redirect(goravelhttp.StatusFound, "/#contact")
	}

	enquiry := models.Enquiry{
		Name: name, Email: email, Phone: phone, Service: service, Message: message,
	}
	if err := facades.Orm().Query().Create(&enquiry); err != nil {
		ctx.Request().Session().Flash("enquiry_error", "We couldn't save your enquiry. Please try again.")
		return ctx.Response().Redirect(goravelhttp.StatusFound, "/#contact")
	}

	ctx.Request().Session().Flash("enquiry_success", "Thank you! Our team will contact you shortly.")
	return ctx.Response().Redirect(goravelhttp.StatusFound, "/#contact")
}
