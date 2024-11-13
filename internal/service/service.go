package service

import (
	"net/smtp"

	"github.com/acronix0/Email-Service-Go/internal/domain"
)

type EmailService interface {
	SendOrder(order domain.Order) error
	SendReset(to, password string) error
}

type SMTPClient struct {
	auth smtp.Auth
	orderRecipientEmail string
	smtpHost string
}

func NewSMTPClient(userName, password, host, orderRecipient string) *SMTPClient {
	auth := smtp.PlainAuth("", userName, password, host)
  return &SMTPClient{auth: auth, orderRecipientEmail: orderRecipient, smtpHost: host}
}

