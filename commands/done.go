package commands

import (
	"fmt"

<<<<<<< HEAD
	"https://github.com/urfave/cli"

	"https://github.com/ajmarroquin/asana/api"
=======
	"github.com/codegangsta/cli"

	"github.com/thash/asana/api"
>>>>>>> parent of 91b12b0 (changing codegansta to urfav, and thash to ajmarroquin)
)

func Done(c *cli.Context) {
	task := api.Update(api.FindTaskId(c.Args().First(), false), "completed", "true")
	fmt.Println("DONE! : " + task.Name)
}
