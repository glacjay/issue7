package main

import (
	"container/list"
	"fmt"
	"os"
)

var (
	OutFile string
	InFiles = list.New()
)

func main() {
	parseArgs()
	initPkgs()
	initLex()
	initType()
	initGen()
	for inFile := range InFiles.Iter() {
		err := setInputFile(inFile.(string))
		if err != nil {
			fmt.Printf("Cannot open and read input file '%s': %v",
				inFile, err)
			os.Exit(1)
		}
		Parse()
	}
	if OutFile == "" {
		OutFile = fmt.Sprintf("%s.%c", LocalPkg.name, TheChar)
	}
	testDclStack()
	makePkg(LocalPkg.name)
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
				OutFile = param
				i += delta
			default:
				usage()
			}
		} else {
			InFiles.PushBack(argv)
		}
	}

	if InFiles.Len() < 1 {
		usage()
	}
}

func usage() {
	print(`flags:
  -o file    -- specify output file
`)
	os.Exit(0)
}
