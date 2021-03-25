package commands

import (
	"fmt"
	"regexp"
	"time"

<<<<<<< HEAD
	"https://github.com/urfavee/cli"

	"https://github.com/ajmarroquin/asana/api"
=======
	"github.com/codegangsta/cli"

	"github.com/thash/asana/api"
>>>>>>> parent of 91b12b0 (changing codegansta to urfave, and thash to ajmarroquin)
)

const (
	DateRegexp = "[:digit:]{4}-[:digit:]{2}-[:digit:]{2}"
)

func DueOn(c *cli.Context) {
	taskId := api.FindTaskId(c.Args().First(), true)
	task := api.Update(taskId, "due_on", toDate(c.Args().First()))
	fmt.Println("set due on [ " + task.Due_on + " ] :" + task.Name)
}

func toDate(str string) string {
	switch {
	case regexp.MustCompile(DateRegexp).MatchString(str):
		return str
	case regexp.MustCompile("today").MatchString(str):
		return time.Now().Format("2006-01-02")
	case regexp.MustCompile("tomorrow").MatchString(str):
		d, _ := time.ParseDuration("24h")
		return time.Now().Add(d).Format("2006-01-02")
	default:
		// Asana API should return err.
		return str
	}
}
