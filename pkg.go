package ast

import (
	"fmt"
	"os"
)

type Pkg struct {
	Name   string
	Path   string
	Prefix string
}

var (
	LocalPkg *Pkg
)

var allPackages = make(map[string]*Pkg)

func InitPkgs() {
	// TODO
	LocalPkg = GetPkg("")
	LocalPkg.Prefix = `""`
}

func GetPkg(path string) *Pkg {
	p, ok := allPackages[path]
	if ok {
		return p
	}

	p = &Pkg{"", path, pathToPrefix(path)}
	allPackages[path] = p
	return p
}

func MakePkg(name string) {
	if LocalPkg.Name == "" {
		if name == "_" {
			fmt.Fprintf(os.Stderr, "invalid package name _")
		}
		LocalPkg.Name = name
	} else {
		if LocalPkg.Name != name {
			fmt.Fprintf(os.Stderr, "package %s; expected %s", name, LocalPkg.Name)
		}

		for _, s := range allSyms[LocalPkg] {
			if s.Def == nil {
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
