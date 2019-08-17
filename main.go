package main // import "moul.io/moulsay"

import (
	"fmt"
	"os"
	"strings"

	"moul.io/moulsay/moulsay"
)

const totalMaxWidth = 72

func main() {
	fmt.Println(moulsay.Say(strings.Join(os.Args[1:], " "), totalMaxWidth))
}
