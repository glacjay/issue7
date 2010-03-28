package ast

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
	Name  string
	Lex   int
	Def   *Node
	Block int
	Pkg   *Pkg
}

var allSyms = make(map[*Pkg](map[string]*Sym))
var dclStack = list.New()

func LookupSym(name string) *Sym {
	return LookupPkgSym(name, LocalPkg)
}

func LookupPkgSym(name string, pkg *Pkg) (s *Sym) {
	p, ok := allSyms[pkg]
	if ok {
		s, ok = p[name]
		if ok {
			return s
		}
	} else {
		allSyms[pkg] = make(map[string]*Sym)
		p = allSyms[pkg]
	}

	s = new(Sym)
	s.Name = name
	p[name] = s
	return s
}

func pushSym() *Sym {
	s := new(Sym)
	dclStack.PushBack(s)
	return s
}

func copySym(d, s *Sym) {
	d.Pkg = s.Pkg
	d.Name = s.Name
	d.Def = s.Def
	d.Block = s.Block
}
