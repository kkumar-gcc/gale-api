package auth

import (
	"errors"

	"github.com/goravel/framework/facades"
	"github.com/goravel/framework/support/carbon"

	"goravel/app/models"
	"goravel/app/utils"
)

type PasswordReset interface {
	GenerateToken(email string) (string, error)
	DestroyToken(*models.PasswordResetToken) error
	FindByEmailAndToken(email, token string) (*models.PasswordResetToken, error)
}

type PasswordResetImpl struct{}

func NewPasswordResetImpl() *PasswordResetImpl {
	return &PasswordResetImpl{}
}

func (s *PasswordResetImpl) GenerateToken(email string) (string, error) {
	token, err := utils.GenerateResetToken()
	if err != nil {
		return "", err
	}

	resetToken := models.PasswordResetToken{
		Email: email,
		Token: token,
	}

	if err := facades.Orm().Query().Create(&resetToken); err != nil {
		return "", err
	}

	return token, nil
}

func (s *PasswordResetImpl) FindByEmailAndToken(email, token string) (*models.PasswordResetToken, error) {
	var resetToken models.PasswordResetToken
	if err := facades.Orm().Query().Where("email = ? AND token = ?", email, token).First(&resetToken); err != nil {
		return nil, err
	}

	if carbon.Now().DiffAbsInHours(resetToken.CreatedAt.Carbon) > 24 {
		return nil, errors.New("token has expired")
	}

	return &resetToken, nil
}

func (s *PasswordResetImpl) DestroyToken(passwordReset *models.PasswordResetToken) error {
	_, err := facades.Orm().Query().Delete(&passwordReset)
	return err
}
