package app

import (
	"github.com/MikeMwita/tweeterapp/config"
	"github.com/MikeMwita/tweeterapp/internal"
	"github.com/MikeMwita/tweeterapp/internal/services"
)

type App struct {
	config *config.Config
}

func Initialize() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}
	app := &App{config: cfg}
	err = app.auth()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) auth() error {
	twClient, err := internal.NewClient(a.config)
	if err != nil {
		return err
	}
	tweetService := services.NewDefaultTwService(twClient)
	// Authenticate bot
	err = tweetService.Auth()
	if err != nil {
		return err
	}

	return nil
}
