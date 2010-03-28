%{
package parse

import (
	"fmt"

	"./ast"
)
%}

%union {
	sym *ast.Sym
}

%token <sym> LXFUNC
%token <sym> LXNAME
%token <sym> LXPACKAGE

%type <sym> sym

%left NotPackage
%left LXPACKAGE

%%

file:
	package
	xdcl_list

package:
	%prec NotPackage
|	LXPACKAGE sym ';'

xdcl_list:
	{}
|	xdcl_list xdcl ';'

xdcl:
	{}
|	xfunc_dcl

xfunc_dcl:
	LXFUNC func_dcl func_body

func_dcl:
	dcl_name '(' ')'

func_body:
	{}
|	'{' '}'

dcl_name:
	sym

sym:
	LXNAME

%%

func Lex() int { return nextToken() }

func Error(fmtstr string, args ...interface{}) {}
