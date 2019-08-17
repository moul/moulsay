package moulsay

import (
	"strings"

	"moul.io/asciimoul"
)

func bodyLinesWithPadding() []string {
	lines := strings.Split(asciimoul.Body, "\n")
	longest := 0
	for _, line := range lines {
		if len(line) > longest {
			longest = len(line)
		}
	}
	for i, line := range lines {
		lines[i] = line + strings.Repeat(" ", longest-len(line))
	}
	return lines
}

func justifyCenter(s string, n int) string {
	lines := strings.Split(s, "\n")
	for i, line := range lines {
		if len(line) < n {
			lines[i] = strings.Repeat(" ", (n-len(line))/2) + line
		}
	}
	return strings.Join(lines, "\n")
}

func trimEmptyLines(input []string) []string {
	left := 0
	right := len(input)
	for i := 0; i < len(input); i++ {
		if strings.TrimSpace(input[i]) == "" {
			left++
		} else {
			break
		}
	}
	for i := right; i > left; i-- {
		if strings.TrimSpace(input[i-1]) == "" {
			right--
		} else {
			break
		}
	}
	return input[left:right]
}

func emptyLines(amount int, width int) []string {
	if amount <= 1 {
		return []string{}
	}
	lines := make([]string, amount)
	emptyLine := strings.Repeat(" ", width)
	for i := 0; i < amount; i++ {
		lines[i] = emptyLine
	}
	return lines
}
