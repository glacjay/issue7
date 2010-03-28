package main

import (
	"container/list"
	"os"

	"./ast"
	"./gen"
)

var (
	outFile string
	inFiles = list.New()
)

func main() {
	parseArgs()
	ast.InitPackages()
	ast.InitLex()
	ast.InitType()
	gen.InitGen()
}

func parseArgs() {
	for i := 1; i < len(os.Args); i++ {
		argv := os.Args[i]
		if argv[0] == '-' {
			if len(argv) == 1 {
				usage()
			}

			var param string
			var delta = 0
			if len(argv) > 2 {
				param = argv[2:]
			} else {
				if i == len(os.Args)-1 {
					param = ""
				} else {
					param = os.Args[i+1]
				}
				delta++
			}

			switch argv[1] {
			case 'o':
				if param == "" {
					usage()
				}
				outFile = param
				i += delta
			default:
				usage()
			}
		} else {
			inFiles.PushBack(argv)
		}
	}

	if inFiles.Len() < 1 {
		usage()
	}
}

func usage() {
	print(`flags:
  -o file    -- specify output file
`)
	os.Exit(0)
}
