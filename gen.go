package ast

const ( // ???
	PXXX = iota

	PEXTERN
	PAUTO
	PPARAM
	PPARAMOUT
	PPARAMREF
	PFUNC

	PHEAP = 1 << 7
)

var (
	DclCtx int

	MaxBlock int
	CurBlock int
)

func InitGen() {
	MaxBlock = 1
	DclCtx = PEXTERN
}
