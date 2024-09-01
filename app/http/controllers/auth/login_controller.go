package auth

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/constants"

	authRequests "goravel/app/http/requests/auth"
	authServices "goravel/app/services/auth"
	"goravel/app/utils"
)

type LoginController struct {
	userService authServices.User
	hashService authServices.Hash
}

func NewLoginController(userService authServices.User, hashService authServices.Hash) *LoginController {
	return &LoginController{
		userService: userService,
		hashService: hashService,
	}
}

func (r *LoginController) Store(ctx http.Context) http.Response {
	var loginRequest authRequests.LoginRequest
	errors, err := ctx.Request().ValidateRequest(&loginRequest)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrUnprocessableEntity).SetErrors(err.Error()).Build(ctx)
	}

	if errors != nil && len(errors.All()) > 0 {
		return utils.NewJsonResponse().SetMessage(constants.ErrBadRequest).SetErrors(errors.All()).Build(ctx)
	}

	user, err := r.userService.FindByEmail(loginRequest.Email)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrUserNotFound).SetErrors(err.Error()).Build(ctx)
	}

	if !r.hashService.Check(loginRequest.Password, user.Password) {
		return utils.NewJsonResponse().SetMessage(constants.ErrInvalidCredentials).Build(ctx)
	}

	token, err := r.userService.GenerateToken(ctx, user)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrInternalServer).SetErrors(err.Error()).Build(ctx)
	}

	return utils.NewJsonResponse().SetMessage(constants.SuccessLoginSuccessful).SetData(map[string]string{
		"token": token,
	}).Build(ctx)
}

func (r *LoginController) Destroy(ctx http.Context) http.Response {
	if err := r.userService.DestroyToken(ctx); err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrInternalServer).SetErrors(err.Error()).Build(ctx)
	}

	return utils.NewJsonResponse().SetMessage(constants.SuccessLogoutSuccessful).Build(ctx)
}
