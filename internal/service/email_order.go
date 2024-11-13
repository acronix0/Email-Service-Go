package service

import (
	"net/smtp"

	"github.com/acronix0/Email-Service-Go/internal/domain"
)

func (c *SMTPClient) SendOrder(order domain.Order) error {
	msg := []byte("test order")
	err := smtp.SendMail(c.smtpHost + ":" + "587", c.auth,"support@dm-trade.pro", []string{c.orderRecipientEmail}, msg)
  return err
   
}