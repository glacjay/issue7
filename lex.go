package main

func initLex() {
	for _, s := range InitialSyms {
		sym := lookupSym(s.name)
		sym.lex = s.lex
	}
}
