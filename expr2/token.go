package libexpr

import (
	. "asdf"
	"fmt"
)

func pickFirstExpr(tokens []*Token) (*Expr, []*Token) {
	var token *Token

	token, tokens = pickFirstToken(tokens)
	if nil == token {
		panic("no found single' next")
	} else {
		switch token.t {
		case TypeExprAtomic, TypeExprRaw:
		default:
			Panic("single's next is %s", token)
		}
	}

	return token.expr(), tokens
}

func pickFirstToken(tokens []*Token) (*Token, []*Token) {
	if len(tokens) > 0 {
		return tokens[0], tokens[1:]
	} else {
		return nil, tokens
	}
}

type Token struct {
	t Type
	v interface{}
}

func (me *Token) RawString() string {
	switch me.t {
	case TypeValue, TypeExprRaw:
		return me.Value()
	case TypeKeyWord:
		return me.Keyword().keyword
	case TypeExprAtomic:
		return me.expr().atomic.String()
	default:
		return me.Buildin().String()
	}
}

func (me *Token) String() string {
	return fmt.Sprintf("%s: %s", me.t, me.RawString())
}

func (me *Token) TString() string {
	return "Token" + me.String()
}

func (me *Token) Keyword() *Keyword {
	return me.v.(*Keyword)
}

func (me *Token) Buildin() Buildin {
	return me.v.(Buildin)
}

func (me *Token) Logic() Logic {
	return me.v.(Logic)
}

func (me *Token) Op() Op {
	return me.v.(Op)
}

func (me *Token) Value() string {
	return me.v.(string)
}

func (me *Token) aexpr() *Expr {
	return me.v.(*Expr)
}

func (me *Token) rexpr() *Expr {
	return Scan(me.Value())
}

func (me *Token) expr() *Expr {
	switch me.t {
	case TypeExprAtomic:
		return me.aexpr()
	case TypeExprRaw:
		return me.rexpr()
	default:
		panic("token isn't expr")
		return nil
	}
}
