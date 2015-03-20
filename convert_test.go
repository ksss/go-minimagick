package minimagick

import (
	"fmt"
	"testing"
)

func TestConvertResize(t *testing.T) {
	c := NewConvert()
	c.Resize("100x100")
	if fmt.Sprint(c) != "convert -resize 100x100" {
		t.Fatalf("fatal: " + c.String())
	}
}
