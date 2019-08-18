package main // import "moul.io/moulsay"

import (
	"fmt"
	"io/ioutil"
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
	message := strings.Join(c.Args().Slice(), " ")
	if message == "" {
		in, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		message = string(in)
	}
	out, err := moulsay.Say(message, c.Int("max-width"))
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
