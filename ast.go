// Header {{{1
//////////////////////////////////////////////////////////////////////////////
package ast

import (
	"container/list"
	"fmt"
)


// Lexical {{{1
//////////////////////////////////////////////////////////////////////////////
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


// Type {{{1
//////////////////////////////////////////////////////////////////////////////
const (
	TYXXX = iota

	TYPTR32
	TYPTR64

	TYFUNC
	TYSTRUCT

	NTYPE
)

const (
	ptrType = TYPTR32
)

type Type struct {
	EType int
	Type  *Type

	IsFuncArg bool
}

type TypeProp struct {
	typ *Type

	Direct  int // direct used type
	OkForEq bool
}

var Types [NTYPE]TypeProp

func InitType() {
	for i, tp := range Types {
		tp.Direct = i
	}

	t := &Types[TYFUNC]
	t.Direct = ptrType
	t.OkForEq = true
	t.typ = FuncType(nil, nil, nil)
}

func GetType(etype int) *Type {
	t := new(Type)
	t.EType = etype
	return t
}

func FuncType(this *Node, in, out NodeList) *Type {
	t := GetType(TYFUNC)
	t.Type = GenStruct(nil, TYFUNC)
	return t
}

func GenStruct(l NodeList, etype int) *Type {
	funcArg := false
	if etype == TYFUNC {
		funcArg = true
		etype = TYSTRUCT
	}
	t := GetType(etype)
	t.IsFuncArg = funcArg
	return t
}


// OP {{{1
//////////////////////////////////////////////////////////////////////////////
const (
	OPXXX = iota
)


// Symbol {{{1
//////////////////////////////////////////////////////////////////////////////
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


// Package {{{1
//////////////////////////////////////////////////////////////////////////////
type Pkg struct {
	name   string
	path   string
	prefix string
}

var (
	localPkg *Pkg
)

var allPackages = make(map[string]*Pkg)

func InitPackages() {
	// TODO
	localPkg = GetPkg("")
	localPkg.prefix = `""`
}

func GetPkg(path string) *Pkg {
	p, ok := allPackages[path]
	if ok {
		return p
	}

	p = &Pkg{"", path, pathToPrefix(path)}
	allPackages[path] = p
	return p
}

func pathToPrefix(path string) string {
	prefix := ""
	for _, ch := range path {
		if ch <= ' ' || ch == '.' || ch == '%' || ch == '"' {
			prefix += fmt.Sprintf("%%%02x", ch)
		} else {
			prefix += string(ch)
		}
	}
	return prefix
}


// Node {{{1
//////////////////////////////////////////////////////////////////////////////
type Node struct{}

type NodeList *list.List
