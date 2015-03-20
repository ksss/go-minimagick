package info

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ksss/go-minimagick/identify"
)

func Width(path string) (int, error) {
	return format(path, "%w")
}

func Height(path string) (int, error) {
	return format(path, "%h")
}

func format(path string, format string) (int, error) {
	out, err := identify.New().Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed: %s\n", err)
		return -1, err
	}
	return strconv.Atoi(string(out))
}
