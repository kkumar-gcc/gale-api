package auth

import (
	"github.com/goravel/framework/contracts/http"

	"goravel/app/constants"
	authServices "goravel/app/services/auth"
	"goravel/app/utils"
)

type PasswordController struct {
	userService          authServices.User
	passwordResetService authServices.PasswordReset
	hashService          authServices.Hash
}

func NewPasswordController(user authServices.User, passwordReset authServices.PasswordReset, hashService authServices.Hash) *PasswordController {
	return &PasswordController{
		userService:          user,
		passwordResetService: passwordReset,
		hashService:          hashService,
	}
}

func (r *PasswordController) Store(ctx http.Context) http.Response {
	token := ctx.Request().Input("token")
	if token == "" {
		return utils.NewJsonResponse().SetMessage(constants.ErrBadRequest).SetErrors("Token is required").Build(ctx)
	}

	password := ctx.Request().Input("password")
	if password == "" {
		return utils.NewJsonResponse().SetMessage(constants.ErrBadRequest).SetErrors("Password is required").Build(ctx)
	}

	email := ctx.Request().Input("email")
	if email == "" {
		return utils.NewJsonResponse().SetMessage(constants.ErrBadRequest).SetErrors("Email is required").Build(ctx)
	}

	passwordReset, err := r.passwordResetService.FindByEmailAndToken(email, token)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrPasswordResetToken).SetErrors(err.Error()).Build(ctx)
	}

	user, err := r.userService.FindByEmail(passwordReset.Email)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrUserNotFound).SetErrors(err.Error()).Build(ctx)
	}

	user.Password, err = r.hashService.Make(password)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrHashingPassword).SetErrors(err.Error()).Build(ctx)
	}

	if err := r.userService.Save(user); err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrInternalServer).SetErrors(err.Error()).Build(ctx)
	}

	if err := r.passwordResetService.DestroyToken(passwordReset); err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrInternalServer).SetErrors(err.Error()).Build(ctx)
	}

	return utils.NewJsonResponse().SetMessage(constants.SuccessPasswordReset).Build(ctx)
}
