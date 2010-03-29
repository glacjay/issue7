package main

const (
	TYXXX = iota

	TYPTR32
	TYPTR64

	TYFUNC
	TYSTRUCT

	NTYPE
)

const (
	PtrType = TYPTR32
)

type Type struct {
	etype int
	type_ *Type

	isFuncArg bool
}

type TypeProp struct {
	type_ *Type

	direct  int // direct used type
	okForEq bool
}

var Types [NTYPE]TypeProp

func initType() {
	for i, tp := range Types {
		tp.direct = i
	}

	t := &Types[TYFUNC]
	t.direct = PtrType
	t.okForEq = true
	t.type_ = funcType(nil, nil, nil)
}

func getType(etype int) *Type {
	t := new(Type)
	t.etype = etype
	return t
}

func funcType(this *Node, in, out *NodeList) *Type {
	t := getType(TYFUNC)
	t.type_ = genStruct(nil, TYFUNC)
	return t
}

func genStruct(l *NodeList, etype int) *Type {
	funcArg := false
	if etype == TYFUNC {
		funcArg = true
		etype = TYSTRUCT
	}
	t := getType(etype)
	t.isFuncArg = funcArg
	return t
}
