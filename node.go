package main

import (
	"container/list"
	"fmt"
)

type Node struct {
	op    int
	sym   *Sym
	type_ *Type
	class int

	addrable bool

	funcName *Node
	funcDcl  *NodeList

	funcType *Node
	defn     *Node

	genVar int
}

type NodeList struct {
	*list.List
}

func (nl *NodeList) PushBackNList(nl2 *NodeList) {
	nl.List.PushBackList(nl2.List)
}

var (
	ExternDcl *NodeList

	CurFunc *Node
)

func makeNode(op int, left, right *Node) *Node {
	n := new(Node)
	n.op = op
	return n
}

func dclName(sym *Sym) *Node {
	if DclCtx == PEXTERN && sym.block <= 1 {
		if sym.def == nil {
			oldName(sym)
		}
		if sym.def.op == OPNONAME {
			return sym.def
		}
	}

	n := newName(sym)
	n.op = OPNONAME
	return n
}

var (
	GenInit int
	GenType int
	GenVar  int
)

func renameInit(n *Node) *Node {
	s := n.sym
	if s == nil || s.name != "init" {
		return n
	}
	GenInit++
	s = lookupSym(fmt.Sprintf("init\xc2\xb7%d", GenInit))
	return newName(s)
}

func declareFuncHeader(n *Node) {
	if n.funcName != nil {
		n.funcName.op = OPNAME
		declare(n.funcName, PFUNC)
		n.funcName.defn = n
	}
}

func newNodeList() *NodeList {
	return &NodeList{list.New()}
}

func oldName(s *Sym) *Node {
	n := s.def
	if n == nil {
		n = newName(s)
		n.op = OPNONAME
		s.def = n
	}

	return n
}

func newName(s *Sym) *Node {
	n := makeNode(OPNAME, nil, nil)
	n.sym = s
	n.type_ = nil
	n.addrable = true
	return n
}

func declare(n *Node, ctx int) {
	if isBlank(n) {
		return
	}

	s := n.sym
	gen := 0
	if ctx == PEXTERN {
		ExternDcl.PushBack(n)
	} else {
		if CurFunc == nil && ctx == PAUTO {
			// TODO
		}
		if CurFunc != nil {
			CurFunc.funcDcl.PushBack(n)
		}
		if n.op == OPTYPE {
			GenType++
			gen = GenType
		} else if n.op == OPNAME {
			GenVar++
			gen = GenVar
		}
		pushDcl(s)
	}

	s.block = CurBlock
	s.def = n
	n.genVar = gen
	n.class = ctx
}

func pushDcl(s *Sym) *Sym {
	d := pushSym()
	copySym(d, s)
	return d
}

func isBlank(n *Node) bool {
	if n == nil || n.sym == nil {
		return false
	}
	return n.sym.name == "_"
}
