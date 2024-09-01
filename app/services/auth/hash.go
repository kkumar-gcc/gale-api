package auth

import "github.com/goravel/framework/facades"

type Hash interface {
	Make(value string) (string, error)
	Check(value, hash string) bool
}

type HashImpl struct{}

func NewHashImpl() *HashImpl {
	return &HashImpl{}
}

func (s *HashImpl) Make(value string) (string, error) {
	return facades.Hash().Make(value)
}

func (s *HashImpl) Check(value, hash string) bool {
	return facades.Hash().Check(value, hash)
}
