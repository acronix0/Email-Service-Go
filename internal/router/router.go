package router

import (
	"encoding/json"
	"log/slog"

	"github.com/IBM/sarama"
	"github.com/acronix0/Email-Service-Go/internal/domain"
	"github.com/acronix0/Email-Service-Go/internal/service"
)

type MessageRouter struct {
	emailService service.EmailService
	log *slog.Logger

}
type ResetMessage struct{
	email string
	newPassword string
}
func NewMessageRouter(emailService service.EmailService, log *slog.Logger) *MessageRouter {
  return &MessageRouter{emailService: emailService, log: log}
}

func (r *MessageRouter) RouteMessage(message *sarama.ConsumerMessage) {
	switch message.Topic {
  case ResetTopic:
		var resetMessage = ResetMessage{}
		json.Unmarshal(message.Value, &resetMessage)
    r.emailService.SendReset(resetMessage.email, resetMessage.newPassword)
  case OrderTopic:
		var orderMessage = domain.Order{}
    //r.emailService.SendOrder(orderMessage)
		 r.emailService.SendOrder(orderMessage)
  default:
    r.log.Debug("Received message from unknown topic: %s", message.Topic)
  }
}

