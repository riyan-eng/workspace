package util

import (
	"fmt"
	"server/env"
)

type sender struct{}

func NewSender() *sender {
	return &sender{}
}
func (m *sender) ResetPasswordToken(email, link, expired string) error {
	sender := NewGmailSender("SISAMBI", env.NewEnv().SMTP_EMAIL, env.NewEnv().SMTP_PASSWORD)
	subject := "Reset Password Verification"
	content := NewTemplate().EmailResetPassword(link, expired)
	to := []string{email}
	if err := sender.SendEmail(subject, content, to, nil, nil, nil); err != nil {
		return fmt.Errorf("failed sent email to: %v", email)
	}
	return nil
}
