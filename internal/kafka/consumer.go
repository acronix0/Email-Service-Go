package kafka

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/IBM/sarama"
	"github.com/acronix0/Email-Service-Go/internal/config"
	"github.com/acronix0/Email-Service-Go/internal/router"
)

type ConsumerGroupHandler struct{
	config *config.Config
	router *router.MessageRouter
}


func NewConsumerGroupHandler(config *config.Config, router *router.MessageRouter) *ConsumerGroupHandler {
  return &ConsumerGroupHandler{config: config, router: router}
}
func (h ConsumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		h.router.RouteMessage(message)
	}
	return nil
}

func Run(ctx context.Context, cfg *config.Config, router *router.MessageRouter) error{
	config := sarama.NewConfig()
  config.Consumer.Return.Errors = true
  config.Version = sarama.V3_3_0_0

  consumerGroup, err := sarama.NewConsumerGroup(cfg.KafkaConfig.BootstrapServer, cfg.KafkaConfig.ConsumerGroup, config)
  if err!= nil {
    return err
  }
	
  go func() {
    for {
			handler := NewConsumerGroupHandler(cfg,router)
    	if err := consumerGroup.Consume(ctx, cfg.KafkaConfig.Topics, handler); err != nil {
     	  log.Fatalf("Consume error: %v", err)
     	}
    }
  }()


	sigterm := make(chan os.Signal, 1)
  signal.Notify(sigterm, os.Interrupt)
  <-sigterm
	return nil
}


func (ConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
    return nil
}

func (ConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
    return nil
}