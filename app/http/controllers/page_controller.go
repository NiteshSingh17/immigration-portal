package controllers

import (
	"github.com/goravel/framework/contracts/http"
)

type PageController struct {
}

func NewPageController() *PageController {
	return &PageController{}
}

func (r *PageController) Home(ctx http.Context) http.Response {
	return ctx.Response().View().Make("home.tmpl")
}

func (r *PageController) EasyPrCertificate(ctx http.Context) http.Response {
	return ctx.Response().View().Make("easy-pr-certificate.tmpl")
}

func (r *PageController) About(ctx http.Context) http.Response {
	return ctx.Response().View().Make("about.tmpl")
}

func (r *PageController) Contact(ctx http.Context) http.Response {
	return ctx.Response().View().Make("contact.tmpl")
}

func (r *PageController) ContactSubmit(ctx http.Context) http.Response {
	// Handle form submissions if needed
	name := ctx.Request().Input("name")
	email := ctx.Request().Input("email")
	phone := ctx.Request().Input("phone")
	message := ctx.Request().Input("message")

	_ = name
	_ = email
	_ = phone
	_ = message

	return ctx.Response().Json(http.StatusOK, map[string]any{
		"success": true,
		"message": "Thank you for reaching out! Our team will contact you shortly.",
	})
}

func (r *PageController) OinpCalculator(ctx http.Context) http.Response {
	return ctx.Response().View().Make("oinp-calculator.tmpl")
}

func (r *PageController) ExpressEntry(ctx http.Context) http.Response {
	return ctx.Response().View().Make("express-entry.tmpl")
}

func (r *PageController) StudentVisa(ctx http.Context) http.Response {
	return ctx.Response().View().Make("student-visa.tmpl")
}

func (r *PageController) WorkPermit(ctx http.Context) http.Response {
	return ctx.Response().View().Make("work-permit.tmpl")
}

func (r *PageController) VisitorVisa(ctx http.Context) http.Response {
	return ctx.Response().View().Make("visitor-visa.tmpl")
}

func (r *PageController) SpouseSponsorship(ctx http.Context) http.Response {
	return ctx.Response().View().Make("spouse-sponsorship.tmpl")
}

func (r *PageController) CaregiverLmia(ctx http.Context) http.Response {
	return ctx.Response().View().Make("caregiver-lmia.tmpl")
}

func (r *PageController) CaregiverPermanentResidency(ctx http.Context) http.Response {
	return ctx.Response().View().Make("caregiver-permanent-residency.tmpl")
}

func (r *PageController) Citizenship(ctx http.Context) http.Response {
	return ctx.Response().View().Make("citizenship.tmpl")
}

func (r *PageController) BusinessInvestor(ctx http.Context) http.Response {
	return ctx.Response().View().Make("business-investor.tmpl")
}

func (r *PageController) Refugee(ctx http.Context) http.Response {
	return ctx.Response().View().Make("refugee.tmpl")
}

func (r *PageController) Connect(ctx http.Context) http.Response {
	return ctx.Response().View().Make("connect.tmpl")
}

func (r *PageController) TermsConditions(ctx http.Context) http.Response {
	return ctx.Response().View().Make("terms-conditions.tmpl")
}

func (r *PageController) PrivatePolicy(ctx http.Context) http.Response {
	return ctx.Response().View().Make("private-policy.tmpl")
}
