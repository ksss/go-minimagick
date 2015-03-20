package minimagick

import (
	"fmt"
	"testing"
)

func TestMogrifyResize(t *testing.T) {
	c := NewMogrify()
	c.Resize("100x100")
	if fmt.Sprint(c) != "mogrify -resize 100x100" {
		t.Fatalf("fatal: " + c.String())
	}
}
