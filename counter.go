package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/crackcomm/go-clitable"
	"strconv"
)

func counterUp(c *cli.Context) {
	counterName := "default"
	if c.NArg() > 0 {
		counterName = c.Args()[0]
	}
	initFlag := isConfigDirectoryExist()
	if initFlag == false {
		createNewCounterConfig()
	}
	config := unmarshalJSON()
	config.Counters[counterName]++
	config.marshalJSON()
	fmt.Println(config.Counters[counterName])
}

func counterDown(c *cli.Context) {
	counterName := "default"
	if c.NArg() > 0 {
		counterName = c.Args()[0]
	}
	initFlag := isConfigDirectoryExist()
	if initFlag == false {
		createNewCounterConfig()
	}
	config := unmarshalJSON()
	config.Counters[counterName]--
	config.marshalJSON()
	fmt.Println(config.Counters[counterName])
}

func counterList(c *cli.Context) {
	config := unmarshalJSON()
	table := clitable.New([]string{"name", "count"})
	for key, value := range config.Counters {
		table.AddRow(map[string]interface{}{"name": key, "count": strconv.FormatInt(value, 10)})
	}
	table.Print()
}
func counterDelete(c *cli.Context) {
	counterName := "default"
	if c.NArg() > 0 {
		counterName = c.Args()[0]
	}
	initFlag := isConfigDirectoryExist()
	if initFlag == false {
		return
	}
	config := unmarshalJSON()
	delete(config.Counters, counterName)
	config.marshalJSON()
}
func createNewCounterConfig() {
	config := newConfig()
	config.marshalJSON()
}
