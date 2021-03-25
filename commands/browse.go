package commands

import (
	"os/exec"
	"strconv"

	"https://github.com/urfave/cli"

	"https://github.com/ajmarroquin/asana/api"
	"https://github.com/ajmarroquin/asana/config"
	"https://github.com/ajmarroquin/asana/utils"
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
