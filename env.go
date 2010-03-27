package env

import (
	"fmt"
	"os"
)

const (
	theChar   = '8'
	theString = "386"
)

var (
	GoRoot string
	GoOS   string
	GoArch string = theString
)

func init() {
	GoRoot = os.Getenv("GOROOT")
	GoOS = os.Getenv("GOOS")
	if GoRoot == "" || GoOS == "" {
		fmt.Fprintf(os.Stderr, "$GOROOT and $GOOS need to be set correctly.\n")
		os.Exit(1)
	}
}
