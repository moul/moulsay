package moulsay

import (
	"fmt"
	"strings"

	"github.com/eidolon/wordwrap"
)

func Say(message string, totalMaxWidth int) (string, error) {
	art := bodyLinesWithPadding()
	art = trimEmptyLines(art)
	artWidth := len(art[0])
	if minWidth := artWidth + 2; totalMaxWidth < minWidth {
		return "", fmt.Errorf("--max-width should be at least %d", minWidth)
	}
	maxTextWidth := totalMaxWidth - artWidth - 1
	message = wordwrap.Wrapper(maxTextWidth, true)(message)
	message = justifyCenter(message, maxTextWidth)
	messageLines := strings.Split(message, "\n")
	messageLines = trimEmptyLines(messageLines)
	if len(messageLines) > len(art) {
		art = append(
			emptyLines(len(messageLines)-len(art), artWidth),
			art...,
		)
	} else {
		messageLines = append(
			messageLines,
			emptyLines((len(art)-len(messageLines))/3, 0)...,
		)
		messageLines = append(
			emptyLines(len(art)-len(messageLines), 0),
			messageLines...,
		)
	}
	lines := make([]string, len(art))
	for i := 0; i < len(art); i++ {
		lines[i] = strings.TrimRight(fmt.Sprintf("%s %s", art[i], messageLines[i]), " ")
	}
	return strings.Join(lines, "\n"), nil
}
