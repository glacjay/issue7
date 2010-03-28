package ast

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
	Name string
	Lex  int
}

var allSyms = make(map[*Pkg](map[string]*Sym))

func LookupSym(name string) *Sym {
	return LookupPkgSym(name, localPkg)
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
