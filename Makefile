all: issue7

issue7: main.8
	8l -o issue7 main.8

main.8: main.go ast.a parse.a
	8g $<

%.a: %.8
	gopack grc $@ $<

ast.8: lex.go node.go op.go pkg.go sym.go type.go
	8g -o $@ $^

parse.8: scan.go parse.go gen.go
	8g -o $@ $^

parse.go: parse.y
	goyacc -o $@ $<

clean:
	-rm -rf *.[8a]
	-rm -rf issue7
