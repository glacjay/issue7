package ast

const (
	LXFUNC = iota
	LXPACKAGE
)

func InitLex() {
	for _, s := range InitialSyms {
		sym := LookupSym(s.name)
		sym.Lex = s.lex
	}
}
