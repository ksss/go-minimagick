package mogrify

import (
	"fmt"
	"testing"
)

func TestResize(t *testing.T) {
	c := New()
	c.Resize("100x100")
	if fmt.Sprint(c) != "mogrify -resize 100x100" {
		t.Fatalf("fatal: " + c.String())
	}
}
