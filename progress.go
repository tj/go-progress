// Package progress provides a simple terminal progress bar.
package progress

import (
	"fmt"
	"io"
	"math"
	"strings"
)

// Bar is a progress bar.
type Bar struct {
	Delimiter string  // Delimiter for the bar ("|").
	Filled    string  // Filled section representation ("█").
	Empty     string  // Empty section representation ("░")
	Total     float64 // Total value.
	Width     int     // Width of the bar.

	value float64
	text  string
}

// New returns a new bar with the given total.
func New(total float64) *Bar {
	return &Bar{
		Delimiter: "|",
		Filled:    "█",
		Empty:     "░",
		Total:     total,
		Width:     60,
	}
}

// NewInt returns a new bar with the given total.
func NewInt(total int) *Bar {
	return New(float64(total))
}

// Text sets the text value.
func (b *Bar) Text(s string) {
	b.text = s
}

// Value sets the value.
func (b *Bar) Value(n float64) {
	if n > b.Total {
		panic("Bar update value cannot be greater than the total")
	}
	b.value = n
}

// ValueInt sets the value.
func (b *Bar) ValueInt(n int) {
	b.Value(float64(n))
}

// String returns the progress bar.
func (b *Bar) String() string {
	p := b.value / b.Total
	filled := math.Ceil(float64(b.Width) * p)
	empty := math.Floor(float64(b.Width) - filled)

	s := fmt.Sprintf("%3.0f%% ", p*100)
	s += b.Delimiter
	s += strings.Repeat(b.Filled, int(filled))
	s += strings.Repeat(b.Empty, int(empty))
	s += b.Delimiter
	s += " " + b.text

	return s
}

// WriteTo the given writer.
func (b *Bar) WriteTo(w io.Writer) (int64, error) {
	s := fmt.Sprintf("\r   %s ", b.String())
	_, err := io.WriteString(w, s)
	return int64(len(s)), err
}
