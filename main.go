package main // import "moul.io/moulsay"

import (
	"fmt"
	"log"
	"os"
	"strings"

	cli "gopkg.in/urfave/cli.v2"
	"moul.io/moulsay/moulsay"
)

func main() {
	app := cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{Name: "max-width", Aliases: []string{"w"}, Value: 72},
		},
		Action: run,
	}
	if err := app.Run(os.Args); err != nil {
		log.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	fmt.Println(moulsay.Say(strings.Join(c.Args().Slice(), " "), c.Int("max-width")))
	return nil
}
