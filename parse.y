%{
package parse

import (
	"fmt"

	"./ast"
)
%}

%union {
	sym  *ast.Sym
	node *ast.Node
	list *ast.NodeList
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

package:
	%prec NotPackage
	{
		Error("package statement must be first")
		ast.MakePkg("main")
	}
|	LXPACKAGE sym ';'
	{ ast.MakePkg($2.Name) }

xdcl_list:
	{ $$ = ast.NewNodeList() }
|	xdcl_list xdcl ';'

xdcl:
	{}
|	xfunc_dcl
	{
		$$ = ast.NewNodeList()
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
		$$ = ast.MakeNode(ast.OPDCLFUNC, nil, nil)
		$$.FuncName = $1
		$$.FuncName = ast.RenameInit($1)
		n := ast.MakeNode(ast.OPTFUNC, nil, nil)
		$$.FuncName.FuncType = n
	}

func_body:
	{}
|	'{' '}'

dcl_name:
	sym
	{ $$ = ast.DclName($1) }

sym:
	LXNAME

%%

func Lex() int { return nextToken() }

func Error(fmtstr string, args ...interface{}) {}
