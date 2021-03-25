package commands

import (
	"fmt"

	"github.com/urfav/cli"

	"github.com/ajmarroquin/asana/api"
)

func Done(c *cli.Context) {
	task := api.Update(api.FindTaskId(c.Args().First(), false), "completed", "true")
	fmt.Println("DONE! : " + task.Name)
}
