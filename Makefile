include $(GOROOT)/src/Make.$(GOARCH)

TARG = issue7
GOFILES = \
	env.go		\
	gen.go		\
	lex.go		\
	main.go		\
	node.go		\
	op.go		\
	parse.go	\
	pkg.go		\
	scan.go		\
	sym.go		\
	type.go		\

include $(GOROOT)/src/Make.cmd

parse.go: parse.y
	goyacc -o $@ $<
