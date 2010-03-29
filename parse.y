%{
package main

import (
	"fmt"
)
%}

%union {
	sym  *Sym
	node *Node
	list *NodeList
}

%token <sym> LXFUNC
%token <sym> LXNAME
%token <sym> LXPACKAGE

%type <sym> sym

%type <node> dcl_name
%type <node> xfunc_dcl func_dcl

%type <list> xdcl xdcl_list

%left NotPackage
%left LXPACKAGE

%%

file:
	package
	xdcl_list
	{ PkgTree.PushBackNList($2) }

package:
	%prec NotPackage
	{
		Error("package statement must be first")
		makePkg("main")
	}
|	LXPACKAGE sym ';'
	{ makePkg($2.name) }

xdcl_list:
	{ $$ = newNodeList() }
|	xdcl_list xdcl ';'

xdcl:
	{}
|	xfunc_dcl
	{
		$$ = newNodeList()
		$$.PushBack($1)
	}

xfunc_dcl:
	LXFUNC func_dcl func_body
	{
		$$ = $2
	}

func_dcl:
	dcl_name '(' ')'
	{
		$$ = makeNode(OPDCLFUNC, nil, nil)
		$$.funcName = $1
		$$.funcName = renameInit($1)
		n := makeNode(OPTFUNC, nil, nil)
		$$.funcName.funcType = n
	}

func_body:
	{}
|	'{' '}'

dcl_name:
	sym
	{ $$ = dclName($1) }

sym:
	LXNAME

%%

func Lex() int { return nextToken() }

func Error(fmtstr string, args ...interface{}) {}
