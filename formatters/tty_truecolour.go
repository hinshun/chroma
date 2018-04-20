package formatters

import (
	"fmt"
	"io"
	"strings"

	"github.com/alecthomas/chroma"
)

// TTY16m is a true-colour terminal formatter.
var TTY16m = Register("terminal16m", chroma.FormatterFunc(trueColourFormatter))

func trueColourFormatter(w io.Writer, style *chroma.Style, it chroma.Iterator) error {
	tokens := it.Tokens()
	lines := splitTokensIntoLines(tokens)
	lineDigits := len(fmt.Sprintf("%d", len(lines)))

	for index, tokens := range lines {
		line := 1 + index

		fmt.Fprint(w, trueColourAnsi(style.Get(chroma.LineNumbers)))
		fmt.Fprintf(w, "%*d ", lineDigits, line)

		for _, token := range tokens {
			entry := style.Get(token.Type)
			if !entry.IsZero() {
				fmt.Fprint(w, trueColourAnsi(entry))
			}

			value := token.Value
			if strings.Contains(value, "\t") {
				value = strings.Replace(value, "\t", "  ", -1)
			}
			fmt.Fprint(w, value)

			if !entry.IsZero() {
				fmt.Fprint(w, "\033[0m")
			}
		}
	}
	return nil
}

func splitTokensIntoLines(tokens []*chroma.Token) (out [][]*chroma.Token) {
	line := []*chroma.Token{}
	for _, token := range tokens {
		for strings.Contains(token.Value, "\n") {
			parts := strings.SplitAfterN(token.Value, "\n", 2)
			// Token becomes the tail.
			token.Value = parts[1]

			// Append the head to the line and flush the line.
			clone := token.Clone()
			clone.Value = parts[0]
			line = append(line, clone)
			out = append(out, line)
			line = nil
		}
		line = append(line, token)
	}
	if len(line) > 0 {
		out = append(out, line)
	}
	return
}

func trueColourAnsi(entry chroma.StyleEntry) string {
	out := ""
	if entry.Bold == chroma.Yes {
		out += "\033[1m"
	}
	if entry.Underline == chroma.Yes {
		out += "\033[4m"
	}
	if entry.Colour.IsSet() {
		out += fmt.Sprintf("\033[38;2;%d;%d;%dm", entry.Colour.Red(), entry.Colour.Green(), entry.Colour.Blue())
	}
	if entry.Background.IsSet() {
		out += fmt.Sprintf("\033[48;2;%d;%d;%dm", entry.Background.Red(), entry.Background.Green(), entry.Background.Blue())
	}
	return out
}
