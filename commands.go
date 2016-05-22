package main

import (
	"github.com/codegangsta/cli"
)

// Commands list
var Commands = []cli.Command{
	commandUp,
	commandDown,
	commandList,
	commandDelete,
}
var commandUp = cli.Command{
	Name:   "up",
	Usage:  "counter increment +1",
	Action: counterUp,
}
var commandDown = cli.Command{
	Name:   "down",
	Usage:  "counter decrement +1",
	Action: counterDown,
}
var commandList = cli.Command{
	Name:   "list",
	Usage:  "counter list display",
	Action: counterList,
}
var commandDelete = cli.Command{
	Name:   "delete",
	Usage:  "counter delete",
	Action: counterDelete,
}
