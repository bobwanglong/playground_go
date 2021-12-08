package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/logrusorgru/aurora"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "bob's cli"
	app.Version = fmt.Sprintf("%s %s/%s", "1.0.0-cli", runtime.GOOS, runtime.GOARCH)
	app.Action = func(c *cli.Context) error {
		println("Greetings")
		return fmt.Errorf("%stest wrong", "aa")
		// return nil
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(aurora.Red("run wrong"))
		os.Exit(1)
	}
}
