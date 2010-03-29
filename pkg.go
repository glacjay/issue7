package main

import (
	"fmt"
	"os"
)

type Pkg struct {
	name   string
	path   string
	prefix string
}

var (
	LocalPkg *Pkg

	PkgTree *NodeList
)

var AllPackages = make(map[string]*Pkg)

func initPkgs() {
	// TODO
	LocalPkg = getPkg("")
	LocalPkg.prefix = `""`
}

func getPkg(path string) *Pkg {
	p, ok := AllPackages[path]
	if ok {
		return p
	}

	p = &Pkg{"", path, pathToPrefix(path)}
	AllPackages[path] = p
	return p
}

func makePkg(name string) {
	if LocalPkg.name == "" {
		if name == "_" {
			fmt.Fprintf(os.Stderr, "invalid package name _")
		}
		LocalPkg.name = name
	} else {
		if LocalPkg.name != name {
			fmt.Fprintf(os.Stderr, "package %s; expected %s", name, LocalPkg.name)
		}

		for _, s := range AllSyms[LocalPkg] {
			if s.def == nil {
				continue
			}
		}
	}
}

func pathToPrefix(path string) string {
	prefix := ""
	for _, ch := range path {
		if ch <= ' ' || ch == '.' || ch == '%' || ch == '"' {
			prefix += fmt.Sprintf("%%%02x", ch)
		} else {
			prefix += string(ch)
		}
	}
	return prefix
}
