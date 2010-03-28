package ast

import (
	"container/list"
	"fmt"
)

type Node struct {
	Op    int
	Sym   *Sym
	Type  *Type
	Class int

	Addrable bool

	FuncName *Node
	FuncDcl  *NodeList

	FuncType *Node
	Defn     *Node

	GenVar int
}

type NodeList struct {
	*list.List
}

var (
	ExternDcl *NodeList

	CurFunc *Node
)

func MakeNode(op int, left, right *Node) *Node {
	n := new(Node)
	n.Op = op
	return n
}

func DclName(sym *Sym) *Node {
	if DclCtx == PEXTERN && sym.Block <= 1 {
		if sym.Def == nil {
			oldName(sym)
		}
		if sym.Def.Op == OPNONAME {
			return sym.Def
		}
	}

	n := newName(sym)
	n.Op = OPNONAME
	return n
}

var (
	genInit int
	genType int
	genVar  int
)

func RenameInit(n *Node) *Node {
	s := n.Sym
	if s == nil || s.Name != "init" {
		return n
	}
	genInit++
	s = LookupSym(fmt.Sprintf("init\xc2\xb7%d", genInit))
	return newName(s)
}

func DeclareFuncHeader(n *Node) {
	if n.FuncName != nil {
		n.FuncName.Op = OPNAME
		declare(n.FuncName, PFUNC)
		n.FuncName.Defn = n
	}
}

func NewNodeList() *NodeList {
	return &NodeList{list.New()}
}

func oldName(s *Sym) *Node {
	n := s.Def
	if n == nil {
		n = newName(s)
		n.Op = OPNONAME
		s.Def = n
	}

	return n
}

func newName(s *Sym) *Node {
	n := MakeNode(OPNAME, nil, nil)
	n.Sym = s
	n.Type = nil
	n.Addrable = true
	return n
}

func declare(n *Node, ctx int) {
	if isBlank(n) {
		return
	}

	s := n.Sym
	gen := 0
	if ctx == PEXTERN {
		ExternDcl.PushBack(n)
	} else {
		if CurFunc == nil && ctx == PAUTO {
			// TODO
		}
		if CurFunc != nil {
			CurFunc.FuncDcl.PushBack(n)
		}
		if n.Op == OPTYPE {
			genType++
			gen = genType
		} else if n.Op == OPNAME {
			genVar++
			gen = genVar
		}
		pushDcl(s)
	}

	s.Block = CurBlock
	s.Def = n
	n.GenVar = gen
	n.Class = ctx
}

func pushDcl(s *Sym) *Sym {
	d := pushSym()
	copySym(d, s)
	return d
}

func isBlank(n *Node) bool {
	if n == nil || n.Sym == nil {
		return false
	}
	return n.Sym.Name == "_"
}
