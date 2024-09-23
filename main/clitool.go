package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func cliinit() {
	app := cli.NewApp()

	app.Action = func(content *cli.Context) error {

		var naming [100]string
		for i := 0; i < 100; i++ {
			var strc string = strconv.Itoa(i)
			naming[i] = "人生" + strc
			app.Name = naming[i]
			app.Usage = naming[i]
			fmt.Println(content.Args().Get(i))
		}
		return nil
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "cat, dog",
			Usage: "Echo with animal",
		},
	}

	app.Version = "0.1"
	app.Run(os.Args)
}
