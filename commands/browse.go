package commands

import (
	"strconv"
	"os/exec"

<<<<<<< HEAD
	"https://github.com/urfave/cli"

	"https://github.com/ajmarroquin/asana/api"
	"https://github.com/ajmarroquin/asana/config"
	"https://github.com/ajmarroquin/asana/utils"
=======
    "github.com/codegangsta/cli"

	"github.com/thash/asana/api"
	"github.com/thash/asana/config"
	"github.com/thash/asana/utils"
>>>>>>> parent of 91b12b0 (changing codegansta to urfav, and thash to ajmarroquin)
)

func Browse(c *cli.Context) {
	taskId := api.FindTaskId(c.Args().First(), true)
	url := "https://app.asana.com/0/" + strconv.Itoa(config.Load().Workspace) + "/" + taskId
	launcher, err := utils.BrowserLauncher()
	utils.Check(err)
	cmd := exec.Command(launcher, url)
	err = cmd.Start()
	utils.Check(err)
}
