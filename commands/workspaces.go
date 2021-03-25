package commands

import (
	"fmt"
<<<<<<< HEAD

	"https://github.com/ajmarroquin/asana/api"
	"https://github.com/urfavee/cli"
=======
	"github.com/codegangsta/cli"
	"github.com/thash/asana/api"
>>>>>>> parent of 91b12b0 (changing codegansta to urfave, and thash to ajmarroquin)
)

func Workspaces(c *cli.Context) {
	for _, w := range api.Me().Workspaces {
		fmt.Printf("%16d %s\n", w.Id, w.Name)
	}
}
