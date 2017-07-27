package progress

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/tj/assert"
)

var preview = flag.Bool("preview", false, "Preview output rendering.")

func TestBar_previewWidth(t *testing.T) {
	if !*preview {
		t.SkipNow()
	}

	b := NewInt(10)
	b.Width = 25
	b.Empty = " "

	for i := 0; i <= 10; i++ {
		b.ValueInt(i)
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		b.WriteTo(os.Stdout)
	}
}

func TestBar_previewDefaults(t *testing.T) {
	if !*preview {
		t.SkipNow()
	}

	b := NewInt(20)

	for i := 0; i <= 20; i++ {
		b.ValueInt(i)
		b.Text(fmt.Sprintf("iteration %d", i))
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		b.WriteTo(os.Stdout)
	}
}

func TestBarString(t *testing.T) {
	b := NewInt(1000)
	assert.Equal(t, `  0% |░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| `, b.String())

	b.ValueInt(250)
	assert.Equal(t, ` 25% |███████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| `, b.String())

	b.ValueInt(750)
	assert.Equal(t, ` 75% |█████████████████████████████████████████████░░░░░░░░░░░░░░░| `, b.String())

	b.ValueInt(1000)
	assert.Equal(t, `100% |████████████████████████████████████████████████████████████| `, b.String())
}

func TestBarText(t *testing.T) {
	b := NewInt(1000)

	b.Text("Building")
	assert.Equal(t, `  0% |░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| Building`, b.String())

	b.Text("Installing")
	b.ValueInt(250)
	assert.Equal(t, ` 25% |███████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░| Installing`, b.String())
}
