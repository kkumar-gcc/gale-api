package auth

import (
	"github.com/goravel/framework/contracts/http"
	"goravel/app/constants"
	"strconv"

	"goravel/app/http/requests/auth"
	"goravel/app/models"
	authServices "goravel/app/services/auth"
	"goravel/app/utils"
)

type RegisterController struct {
	userService authServices.User
	hashService authServices.Hash
	mailService authServices.Mail
}

func NewRegisterController(userService authServices.User, hashService authServices.Hash, mailService authServices.Mail) *RegisterController {
	return &RegisterController{
		userService: userService,
		hashService: hashService,
		mailService: mailService,
	}
}

func (r *RegisterController) Store(ctx http.Context) http.Response {
	var registerRequest auth.RegisterRequest
	errors, err := ctx.Request().ValidateRequest(&registerRequest)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrUnprocessableEntity).SetErrors(err.Error()).Build(ctx)
	}

	if errors != nil && len(errors.All()) > 0 {
		return utils.NewJsonResponse().SetMessage(constants.ErrBadRequest).SetErrors(errors.All()).Build(ctx)
	}

	if r.userService.Exists(registerRequest.Email) {
		return utils.NewJsonResponse().SetMessage(constants.ErrUserAlreadyExists).Build(ctx)
	}

	hashedPassword, err := r.hashService.Make(registerRequest.Password)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrHashingPassword).SetErrors(err.Error()).Build(ctx)
	}

	user := models.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: hashedPassword,
	}
	if err := r.userService.Create(&user); err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrCreatingUser).SetErrors(err.Error()).Build(ctx)
	}

	hashedEmail, err := r.hashService.Make(user.Email)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrInternalServer).SetErrors(err.Error()).Build(ctx)
	}

	if err := r.mailService.SendVerificationEmail(user.Email, strconv.Itoa(int(user.ID)), hashedEmail); err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrSendingVerificationEmail).SetErrors(err.Error()).Build(ctx)
	}

	token, err := r.userService.GenerateToken(ctx, &user)
	if err != nil {
		return utils.NewJsonResponse().SetMessage(constants.ErrInternalServer).SetErrors(err.Error()).Build(ctx)
	}

	return utils.NewJsonResponse().SetMessage(constants.SuccessUserRegistered).SetData(map[string]string{
		"token": token,
	}).Build(ctx)
}
