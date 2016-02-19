package cli_helpers

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/docker/docker/pkg/homedir"
)

func FixHOME(app *cli.App) {
	appBefore := app.Before

	app.Before = func(c *cli.Context) error {
		key := homedir.Key()
		if os.Getenv(key) != "" {
			return nil
		}

		value := homedir.Get()
		if value == "" {
			return fmt.Errorf("the %q is not set", key)
		}

		os.Setenv(key, value)

		if appBefore != nil {
			return appBefore(c)
		}
		return nil
	}
}
