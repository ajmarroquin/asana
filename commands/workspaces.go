package commands

import (
	"fmt"

	"https://github.com/ajmarroquin/asana/api"
	"https://github.com/urfave/cli"
)

func Workspaces(c *cli.Context) {
	for _, w := range api.Me().Workspaces {
		fmt.Printf("%16d %s\n", w.Id, w.Name)
	}
}
