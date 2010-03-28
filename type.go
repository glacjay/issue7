package ast

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
