package parse

import (
	"io/ioutil"
	"os"
	"unicode"
	"utf8"

	"./ast"
)

type IO struct {
	inFile string
	src    []byte
	ch     int
	index  int
	w      int
	nlSemi bool
}

func (io *IO) readEntireFile() (err os.Error) {
	io.src, err = ioutil.ReadFile(io.inFile)
	return err
}

func (io *IO) next() int {
	if len(io.src[io.index:]) == 0 {
		io.ch = -1
		io.w = 0
		return io.ch
	}

	io.ch, io.w = utf8.DecodeRune(io.src[io.index:])
	io.index += io.w
	return io.ch
}

var curIO *IO

func SetInputFile(name string) os.Error {
	curIO = new(IO)
	curIO.inFile = name
	err := curIO.readEntireFile()
	if err != nil {
		return err
	}
	CurBlock = 1
	curIO.next()
	return nil
}

func nextToken() int {
	tok := nextRealToken()
	return tok
}

func nextRealToken() int {
	c := curIO.ch

	for unicode.IsSpace(c) {
		if c == '\n' && curIO.nlSemi {
			return ';'
		}
		curIO.next()
	}

	if c < 0 {
		return c
	}

	if unicode.IsLetter(c) || c == '_' {
		buf := ""
		for ; unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_'; c = curIO.next() {
			buf += string(c)
		}

		s := ast.LookupSym(buf)
		switch s.Lex {
		// TODO LXIGNORE and LoopHack
		}
		yylval.sym = s
		return s.Lex
	}

	if unicode.IsDigit(c) {
		// TODO
	}

	switch c {
	case '(', ')', '{': // TODO LoopHack
		goto lx
	default:
		goto lx
	}

	return 0

lx:
	curIO.next()
	return c
}
