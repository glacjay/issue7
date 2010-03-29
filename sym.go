package main

import (
	"container/list"
)

type SymInit struct {
	name  string
	lex   int
	etype int
	op    int
}

var InitialSyms = [...]SymInit{
	SymInit{"func", LXFUNC, TYXXX, OPXXX},
	SymInit{"package", LXPACKAGE, TYXXX, OPXXX},
}

type Sym struct {
	name  string
	lex   int
	def   *Node
	block int
	pkg   *Pkg
}

var AllSyms = make(map[*Pkg](map[string]*Sym))
var DclStack = list.New()

func lookupSym(name string) *Sym {
	return lookupPkgSym(name, LocalPkg)
}

func lookupPkgSym(name string, pkg *Pkg) (s *Sym) {
	p, ok := AllSyms[pkg]
	if ok {
		s, ok = p[name]
		if ok {
			return s
		}
	} else {
		AllSyms[pkg] = make(map[string]*Sym)
		p = AllSyms[pkg]
	}

	s = new(Sym)
	s.name = name
	p[name] = s
	return s
}

func pushSym() *Sym {
	s := new(Sym)
	DclStack.PushBack(s)
	return s
}

func copySym(d, s *Sym) {
	d.pkg = s.pkg
	d.name = s.name
	d.def = s.def
	d.block = s.block
}
