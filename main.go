package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "counter"
	app.Usage = "counter app by make golang"
	app.Version = "0.0.1"
	app.Commands = Commands
	app.Run(os.Args)
}
