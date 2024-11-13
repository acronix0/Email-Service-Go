package service

import "net/smtp"

func (c *SMTPClient) SendReset(to, password string) error {
	msg := []byte("test reset")

	return smtp.SendMail("smtp.example.com:25", c.auth, "sender@example.com", []string{to}, msg)
}