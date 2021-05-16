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
	var (
		fs       = flag.NewFlagSet("app", flag.ExitOnError)
		maxWidth = fs.Int("max-width", 72, "set max width")
	)

	root := &ffcli.Command{
		Name:       "moul say",
		ShortUsage: "moulsay [OPTS] words...",
		ShortHelp:  "moulsay -h",
		FlagSet:    fs,
		Exec: func(ctx context.Context, args []string) error {
			return run(ctx, args, *maxWidth)
		},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context, args []string, maxWidth int) error {
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
