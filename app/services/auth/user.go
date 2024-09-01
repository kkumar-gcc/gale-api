package auth

import (
	"errors"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"

	"goravel/app/models"
)

type User interface {
	Create(user *models.User) error
	Save(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(id string) (*models.User, error)
	GetUser(ctx http.Context) (*models.User, error)
	Exists(email string) bool
	DestroyToken(ctx http.Context) error
	GenerateToken(ctx http.Context, user *models.User) (string, error)
}

type UserImpl struct{}

func NewUserImpl() *UserImpl {
	return &UserImpl{}
}

func (s *UserImpl) Create(user *models.User) error {
	if err := facades.Orm().Query().Create(user); err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func (s *UserImpl) Save(user *models.User) error {
	if err := facades.Orm().Query().Save(user); err != nil {
		return errors.New("failed to save user")
	}
	return nil
}

func (s *UserImpl) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := facades.Orm().Query().Where("email = ?", email).FirstOrFail(&user); err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (s *UserImpl) FindById(id string) (*models.User, error) {
	var user models.User
	if err := facades.Orm().Query().Where("id = ?", id).FirstOrFail(&user); err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (s *UserImpl) GetUser(ctx http.Context) (*models.User, error) {
	var user models.User
	if err := facades.Auth(ctx).User(&user); err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (s *UserImpl) Exists(email string) bool {
	_, err := s.FindByEmail(email)
	return err == nil
}

func (s *UserImpl) DestroyToken(ctx http.Context) error {
	if err := facades.Auth(ctx).Logout(); err != nil {
		return errors.New("failed to destroy token")
	}
	return nil
}

func (s *UserImpl) GenerateToken(ctx http.Context, user *models.User) (string, error) {
	token, err := facades.Auth(ctx).Login(user)
	if err != nil {
		return "", errors.New("failed to generate token")
	}
	return token, nil
}
