package main

import (
	"fmt"
	"os"

	cli "gopkg.in/urfave/cli.v2"
)

func main() {
	app := &cli.App{
		Name:  "greet",
		Usage: "say a greeting",
		Action: func(c *cli.Context) error {
			fmt.Println("Greetings")
			return nil
		},
	}

	app.Run(os.Args)
}
