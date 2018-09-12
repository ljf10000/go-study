package libexpr

import (
	. "asdf"
	"fmt"
)

func RegisterKeyword(K Keyword) {
	v, ok := exprBuildinKeywords[K.Key]
	if ok && v.Scope != K.Scope {
		Panic("have register keyword:%s", v)
	}

	keyword := K
	exprBuildinKeywords[keyword.Key] = &keyword
	//    &Keyword{
	//		Key:   keyword,
	//		Type:  Type,
	//		Scope: scope,
	//		// List:  strings.Split(keyword, "."),
	//	}
}

func LineExpr(line string) *Expr {
	return lex3Scan(lex2Scan(lex1Scan(line)))
}

func exprExampl(expr *Expr, level int) string {
	s := "{" + Crlf

	if expr.IsAtomic() {
		atomic := expr.Atomic

		s += TabN(level+1) + atomic.TypeString() + "{" + Crlf

		s += TabN(level+2) + fmt.Sprintf("k: %s,", atomic.K) + Crlf
		s += TabN(level+2) + fmt.Sprintf("op: %s,", atomic.Op) + Crlf
		s += TabN(level+2) + fmt.Sprintf("v: %s,", atomic.V) + Crlf

		s += TabN(level+1) + "}," + Crlf
	} else {
		s += TabN(level+1) + fmt.Sprintf("op: %s,", expr.Logic) + Crlf

		for idx, v := range expr.Children {
			tmp := expr.TypeString() + exprExampl(v, level+1)

			s += TabN(level+1) + fmt.Sprintf("%d: %s,", idx, tmp) + Crlf
		}
	}

	s += TabN(level) + "}"

	return s
}

func ExprExampl(line string) string {
	expr := LineExpr(line)

	s := "lex3:" + Crlf
	s += exprExampl(expr, 0)

	return s
}
