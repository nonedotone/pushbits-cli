package application

import (
	log "github.com/sirupsen/logrus"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/ui"
)

const (
	listEndpoint = "/application"
)

type listCommand struct {
	options.AuthOptions
}

func (c *listCommand) Run(s *options.Options) error {
	if len(c.Password) == 0 {
		c.Password = ui.GetCurrentPassword(c.Username)
	}

	resp, err := api.Get(c.URL, listEndpoint, c.Proxy, c.Username, c.Password)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)

	return nil
}
