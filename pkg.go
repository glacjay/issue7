package ast

import (
	"fmt"
)

type Pkg struct {
	name   string
	path   string
	prefix string
}

var (
	localPkg *Pkg
)

var allPackages = make(map[string]*Pkg)

func InitPackages() {
	// TODO
	localPkg = GetPkg("")
	localPkg.prefix = `""`
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
