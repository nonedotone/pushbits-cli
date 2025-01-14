package user

import (
	log "github.com/sirupsen/logrus"

	"github.com/pushbits/cli/internal/api"
	"github.com/pushbits/cli/internal/options"
	"github.com/pushbits/cli/internal/ui"
)

const (
	listEndpoint = "/user"
)

type listCommand struct {
	options.AuthOptions
}

func (c *listCommand) Run(s *options.Options) error {
	password := ui.GetCurrentPassword(c.Username)

	resp, err := api.Get(c.URL, listEndpoint, c.Proxy, c.Username, password)
	if err != nil {
		log.Fatal(err)
	}

	ui.PrintJSON(resp)

	return nil
}
