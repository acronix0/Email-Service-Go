package app

import (
	"log/slog"
	"os"

	"github.com/acronix0/Email-Service-Go/internal/config"
	"github.com/acronix0/Email-Service-Go/internal/router"
	"github.com/acronix0/Email-Service-Go/internal/service"
)

type serviceProvider struct {
	router *router.MessageRouter
	
	SMTPClient *service.SMTPClient
	config *config.Config
	logger *slog.Logger
}
func NewServiceProvider(cfg *config.Config) *serviceProvider {
	return &serviceProvider{config:cfg}
}

func (s *serviceProvider) GetSMTPClient() *service.SMTPClient{
	if s.SMTPClient == nil {
		s.SMTPClient = service.NewSMTPClient(s.config.SMTPConfig.Username, s.config.SMTPConfig.Password, s.config.SMTPConfig.Host, s.config.OrderInfo.Recipient)
	}
	return s.SMTPClient
}
func (s *serviceProvider) GetRouter() *router.MessageRouter{
	if s.router == nil {
		s.router = router.NewMessageRouter(s.GetSMTPClient(), s.GetLogger())
	}
	return s.router
}
func (s *serviceProvider) GetLogger() *slog.Logger{
	if s.logger == nil {
		s.logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return s.logger
}
