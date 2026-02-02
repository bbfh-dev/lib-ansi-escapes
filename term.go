package libescapes

import (
	"os"

	"golang.org/x/term"
)

var IsTerminal bool

var AllowANSI bool

func init() {
	IsTerminal = term.IsTerminal(int(os.Stdout.Fd()))
	AllowANSI = IsTerminal && os.Getenv("NO_COLOR") != "1"
}

func Optional(sequence string) string {
	if AllowANSI {
		return sequence
	}
	return ""
}
