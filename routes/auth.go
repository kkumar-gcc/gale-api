package routes

import (
	"github.com/goravel/framework/contracts/route"
	"github.com/goravel/framework/facades"

	"goravel/app/http/controllers/auth"
	authServices "goravel/app/services/auth"
)

func Auth() {
	userService := authServices.NewUserImpl()
	hashService := authServices.NewHashImpl()
	mailService := authServices.NewMailImpl()
	passwordResetService := authServices.NewPasswordResetImpl()

	loginController := auth.NewLoginController(userService, hashService)
	registerController := auth.NewRegisterController(userService, hashService, mailService)
	newPasswordController := auth.NewPasswordController(userService, passwordResetService)
	forgotPasswordController := auth.NewForgotPasswordController(userService, passwordResetService, mailService)
	verifyEmailController := auth.NewVerifyEmailController()

	facades.Route().Prefix("auth").Group(func(router route.Router) {
		router.Middleware().Post("/login", loginController.Store)
		router.Post("/register", registerController.Store)
		router.Post("/forgot-password", forgotPasswordController.Store)
		router.Post("/reset-password", newPasswordController.Store)
		router.Get("/verify-email/{id}/{hash}", verifyEmailController.Store)
		router.Post("/logout", loginController.Destroy)
	})
}
