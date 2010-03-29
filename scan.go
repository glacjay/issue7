package main

import (
	"io/ioutil"
	"os"
	"unicode"
	"utf8"
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

var CurIO *IO

func setInputFile(name string) os.Error {
	CurIO = new(IO)
	CurIO.inFile = name
	err := CurIO.readEntireFile()
	if err != nil {
		return err
	}
	CurBlock = 1
	CurIO.next()
	return nil
}

func nextToken() int {
	tok := nextRealToken()
	return tok
}

func nextRealToken() int {
	c := CurIO.ch

	for unicode.IsSpace(c) {
		if c == '\n' && CurIO.nlSemi {
			return ';'
		}
		CurIO.next()
	}

	if c < 0 {
		return c
	}

	if unicode.IsLetter(c) || c == '_' {
		buf := ""
		for ; unicode.IsLetter(c) || unicode.IsDigit(c) || c == '_'; c = CurIO.next() {
			buf += string(c)
		}

		s := lookupSym(buf)
		switch s.lex {
		// TODO LXIGNORE and LoopHack
		}
		yylval.sym = s
		return s.lex
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
	CurIO.next()
	return c
}
