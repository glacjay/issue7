package env

import (
	"fmt"
	"os"
)

const (
	TheChar   = '8'
	TheString = "386"
)

var (
	GoRoot string
	GoOS   string
	GoArch string = TheString
)

func init() {
	GoRoot = os.Getenv("GOROOT")
	GoOS = os.Getenv("GOOS")
	if GoRoot == "" || GoOS == "" {
		fmt.Fprintf(os.Stderr, "$GOROOT and $GOOS need to be set correctly.\n")
		os.Exit(1)
	}
}
