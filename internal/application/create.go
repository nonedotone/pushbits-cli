package application

import (
	log "github.com/sirupsen/logrus"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/ui"
)

const (
	createEndpoint = "/application"
)

type createCommand struct {
	options.AuthOptions
	Name                string `arg:"name" help:"The name of the application"`
	StrictCompatibility bool   `long:"compat" help:"Enforce strict compatibility with Gotify"`
}

func (c *createCommand) Run(s *options.Options) error {
	if len(c.Password) == 0 {
		c.Password = ui.GetCurrentPassword(c.Username)
	}

	data := map[string]interface{}{
		"name":                 c.Name,
		"strict_compatibility": c.StrictCompatibility,
	}

	resp, err := api.Post(c.URL, createEndpoint, c.Proxy, c.Username, c.Password, data)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)

	return nil
}
