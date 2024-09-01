package auth

import (
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/support/carbon"
	"goravel/app/constants"

	authServices "goravel/app/services/auth"
	"goravel/app/utils"
)

type VerifyEmailController struct {
	userService authServices.User
	hashService authServices.Hash
}

func NewVerifyEmailController() *VerifyEmailController {
	return &VerifyEmailController{
		userService: authServices.NewUserImpl(),
		hashService: authServices.NewHashImpl(),
	}
}

func (r *VerifyEmailController) Store(ctx http.Context) http.Response {
	userId := ctx.Request().Input("id")
	hash := ctx.Request().Input("hash")

	if userId == "" || hash == "" {
		return utils.NewJsonResponse().SetMessage(constants.ErrUnprocessableEntity).SetErrors("User ID and hash are required").Build(ctx)
	}

	user, err := r.userService.FindById(userId)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrUserNotFound).SetErrors(err.Error()).Build(ctx)
	}

	if user.EmailVerifiedAt != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrEmailAlreadyVerified).Build(ctx)
	}

	if !r.hashService.Check(user.Email, hash) {
		return utils.NewJsonResponse().SetMessage(constants.ErrInvalidVerificationLink).Build(ctx)
	}

	user.EmailVerifiedAt = &carbon.DateTime{Carbon: carbon.Now()}

	if err := r.userService.Save(user); err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrEmailVerificationFailed).SetErrors(err.Error()).Build(ctx)
	}

	return utils.NewJsonResponse().SetMessage(constants.SuccessEmailVerified).Build(ctx)
}
