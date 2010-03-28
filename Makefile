all: issue7

issue7: main.8
	8l -o issue7 main.8

main.8: main.go ast.a
	8g $<

ast.a: ast.8
	gopack grc $@ $<

ast.8: lex.go node.go op.go pkg.go sym.go type.go
	8g -o $@ $^

clean:
	-rm -rf *.[8a]
	-rm -rf issue7
