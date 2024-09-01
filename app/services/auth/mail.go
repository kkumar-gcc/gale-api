package auth

import (
	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/facades"
)

type Mail interface {
	SendPasswordResetEmail(email, token string) error
	SendVerificationEmail(email string, id, hash string) error
}

type MailImpl struct{}

func NewMailImpl() *MailImpl {
	return &MailImpl{}
}

func (s *MailImpl) SendPasswordResetEmail(email, token string) error {
	return facades.Mail().To([]string{email}).Content(mail.Content{
		Subject: "Password Reset",
		Html:    "<a href=\"http://localhost:3000/reset-password?token=" + token + "\">Reset Password</a>",
	}).Send()
}

func (s *MailImpl) SendVerificationEmail(email string, id, hash string) error {
	return facades.Mail().To([]string{email}).Content(mail.Content{
		Subject: "Email Verification",
		Html:    "<a href=\"http://localhost:3000/verify-email/" + id + "/" + hash + "\">Verify Email</a>",
	}).Send()
}
