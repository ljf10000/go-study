package libexpr

import (
	. "asdf"
	"strings"
)

func RegisterKeyword(keyword string, scope Scope) {
	v, ok := exprBuildinKeywords[keyword]
	if ok && v.scope != scope {
		Panic("have register keyword:%s", v)
	}

	exprBuildinKeywords[keyword] = &Keyword{
		keyword: keyword,
		scope:   scope,
		list:    strings.Split(keyword, "."),
	}
}

func Scan(line string) *Expr {
	return lex3Scan(lex2Scan(lex1Scan(line)))
}
