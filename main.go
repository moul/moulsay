package main // import "moul.io/moulsay"

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/peterbourgon/ff/v3/ffcli"
	"moul.io/moulsay/moulsay"
)

func main() {
	fs := flag.NewFlagSet("app", flag.ExitOnError)
	maxWidth := fs.Int("max-width", 72, "set max width")

	root := &ffcli.Command{
		Name:       "moul say",
		ShortUsage: "moulsay + word",
		ShortHelp:  "moulsay word",
		LongHelp:   "moulsay -max-width=? word for different max width from moul to word (minimum max-width = 27)",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			err := run(ctx, args, *maxWidth)
			if err != nil {
				return err
			}
			return nil
		},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(c context.Context, args []string, maxWidth int) error {
	message := strings.Join(args, " ")
	if message == "" {
		in, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return err
		}
		message = string(in)
	}
	out, err := moulsay.Say(message, maxWidth)
	if err != nil {
		return err
	}
	fmt.Println(out)
	return nil
}
