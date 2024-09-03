package auth

import (
	"github.com/goravel/framework/contracts/mail"
	"github.com/goravel/framework/facades"

	"goravel/app/utils"
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
		Html:    "<a href=\"" + utils.AppURL() + "/reset-password?token=" + token + "\">Reset Password</a>",
	}).Send()
}

func (s *MailImpl) SendVerificationEmail(email string, id, hash string) error {
	return facades.Mail().To([]string{email}).Content(mail.Content{
		Subject: "Email Verification",
		Html:    "<a href=\"" + utils.AppURL() + "/verify-email/" + id + "/" + hash + "\">Verify Email</a>",
	}).Send()
}
