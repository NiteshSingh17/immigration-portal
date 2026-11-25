package routes

import (
	"github.com/goravel/framework/contracts/http"

	"goravel/app/facades"
	"goravel/app/http/controllers"
)

func Web() {
	pageController := controllers.NewPageController()

	// Static assets routes
	facades.Route().Static("public", "./public")
	facades.Route().Static("fonts", "./public/fonts")
	facades.Route().Static("img", "./public/img")
	facades.Route().Static("css", "./public/css")
	facades.Route().Static("js", "./public/js")
	facades.Route().Static("wp-content/themes/twentytwenty/img", "./public/img")
	facades.Route().Static("wp-content/themes/twentytwenty/css", "./public/css")
	facades.Route().Static("wp-content/themes/twentytwenty/js", "./public/js")

	// Home page
	facades.Route().Get("/", pageController.Home)
	facades.Route().Get("/home", pageController.Home)
	facades.Route().Get("/home/", pageController.Home)

	// Primary target landing page
	facades.Route().Get("/easy-pr-certificate", pageController.EasyPrCertificate)
	facades.Route().Get("/easy-pr-certificate/", pageController.EasyPrCertificate)

	// About
	facades.Route().Get("/about", pageController.About)
	facades.Route().Get("/about/", pageController.About)
	facades.Route().Get("/ab", pageController.About)

	// Contact
	facades.Route().Get("/contact", pageController.Contact)
	facades.Route().Get("/contact/", pageController.Contact)
	facades.Route().Post("/contact", pageController.ContactSubmit)

	// Interactive OINP Points Calculator
	facades.Route().Get("/oinp-calculator", pageController.OinpCalculator)
	facades.Route().Get("/oinp-calculator/", pageController.OinpCalculator)

	// Immigration Services
	facades.Route().Get("/express-entry", pageController.ExpressEntry)
	facades.Route().Get("/express-entry/", pageController.ExpressEntry)

	facades.Route().Get("/student-visa", pageController.StudentVisa)
	facades.Route().Get("/student-visa/", pageController.StudentVisa)

	facades.Route().Get("/work-permit", pageController.WorkPermit)
	facades.Route().Get("/work-permit/", pageController.WorkPermit)

	facades.Route().Get("/visitor-visa", pageController.VisitorVisa)
	facades.Route().Get("/visitor-visa/", pageController.VisitorVisa)

	facades.Route().Get("/spouse-sponsorship", pageController.SpouseSponsorship)
	facades.Route().Get("/spouse-sponsorship/", pageController.SpouseSponsorship)

	facades.Route().Get("/caregiver-lmia", pageController.CaregiverLmia)
	facades.Route().Get("/caregiver-lmia/", pageController.CaregiverLmia)

	facades.Route().Get("/caregiver-permanent-residency", pageController.CaregiverPermanentResidency)
	facades.Route().Get("/caregiver-permanent-residency/", pageController.CaregiverPermanentResidency)
	facades.Route().Get("/caregiver-pr", pageController.CaregiverPermanentResidency)
	facades.Route().Get("/caregiver-pr/", pageController.CaregiverPermanentResidency)

	facades.Route().Get("/citizenship", pageController.Citizenship)
	facades.Route().Get("/citizenship/", pageController.Citizenship)

	facades.Route().Get("/business-investor", pageController.BusinessInvestor)
	facades.Route().Get("/business-investor/", pageController.BusinessInvestor)

	facades.Route().Get("/refugee", pageController.Refugee)
	facades.Route().Get("/refugee/", pageController.Refugee)

	// Additional pages
	facades.Route().Get("/connect", pageController.Connect)
	facades.Route().Get("/connect/", pageController.Connect)

	facades.Route().Get("/terms-conditions", pageController.TermsConditions)
	facades.Route().Get("/terms-conditions/", pageController.TermsConditions)

	facades.Route().Get("/private-policy", pageController.PrivatePolicy)
	facades.Route().Get("/private-policy/", pageController.PrivatePolicy)
	facades.Route().Get("/cookie-policy", pageController.PrivatePolicy)
	facades.Route().Get("/cookie-policy/", pageController.PrivatePolicy)

	facades.Route().Get("/imunify-bot-check", func(ctx http.Context) http.Response {
		return ctx.Response().String(200, "OK")
	})
}
