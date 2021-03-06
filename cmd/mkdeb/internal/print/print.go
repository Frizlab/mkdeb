package print

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

func Error(s string, args ...interface{}) {
	fmt.Fprint(os.Stderr, ansi.Color("Error: ", "red"))
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func Start(s string, args ...interface{}) {
	print("green", s, args...)
}

func Step(s string, args ...interface{}) {
	print("blue", s, args...)
}

func print(color, s string, args ...interface{}) {
	fmt.Print(ansi.Color("*** ", color))
	fmt.Printf(s+"\n", args...)
}
