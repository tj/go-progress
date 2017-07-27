package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/tj/go-progress"
	"github.com/tj/go/ansi"

	color "github.com/aybabtme/rgbterm"
)

// gray string.
func gray(s string) string {
	return color.FgString(s, 150, 150, 150)
}

// purple string.
func purple(s string) string {
	return color.FgString(s, 96, 97, 190)
}

func main() {
	ansi.HideCursor()
	defer ansi.ShowCursor()

	b := progress.NewInt(10)
	b.Width = 40
	b.StartDelimiter = gray("|")
	b.EndDelimiter = gray("|")
	b.Filled = purple("█")
	b.Empty = gray("░")

	for i := 0; i <= 10; i++ {
		b.ValueInt(i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		os.Stdout.WriteString(ansi.CenterLine(b.String()))
	}
}
