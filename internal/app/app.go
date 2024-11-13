package app

import (
	"context"
	"log/slog"
	"os"

	"github.com/acronix0/Email-Service-Go/internal/config"
	"github.com/acronix0/Email-Service-Go/internal/kafka"
)

type App struct {
	serviceProvider *serviceProvider
	log             *slog.Logger
	config *config.Config
}

func NewApp(ctx context.Context) (*App, error){
	a:= &App{config: config.MustLoad()}
	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initLogger,
		a.initServiceProvider,
	}
	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
func (a *App) initServiceProvider (_ context.Context) error{
	a.serviceProvider = NewServiceProvider(a.config) 
  return nil
}
func (a *App) Run(ctx context.Context) error{
	 return kafka.Run(ctx, a.config, a.serviceProvider.GetRouter())
}
func (a *App)initLogger(_ context.Context) error {
	switch a.config.Env {
	case config.EnvLocal:
		a.log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case config.EnvProd:
		a.log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}
	return nil
}