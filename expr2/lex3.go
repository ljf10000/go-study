package libexpr

import (
	. "asdf"
)

func lex3Scan(tokens []*Token) *Expr {
	lex := newLex3(tokens)

	lex.start()
	lex.stop()
	Log.Info(lex.DumpString())

	return lex.expr
}

func newLex3(tokens []*Token) *lex3 {
	return &lex3{
		tokens: tokens,
		expr:   newExpr(eFsmInit),
	}
}

type lex3 struct {
	tokens []*Token
	expr   *Expr
}

func (me *lex3) DumpString() string {
	s := "lex3:" + Crlf
	s += me.expr.LString(0)

	return s
}

func (me *lex3) TokenPanic(token *Token) {
	Panic("add token[%s] on fsm[%s]", token, me.expr.fsm)
}

func (me *lex3) start() {
	tokens := me.tokens

	for len(tokens) > 0 {
		tokens = me.scan(tokens)
	}
}

func (me *lex3) stop() {
	switch me.expr.fsm {
	case eFsmAtomic, eFsmSingleOk, eFsmExprOk:
		DoNothing()
	case eFsmExpr:
		me.expr = me.expr.children[0]
	default:
		Panic("left fsm[%s]", me.expr.fsm)
	}
}

func (me *lex3) scan(tokens []*Token) []*Token {
	if len(tokens) == 0 {
		return tokens
	}

	switch me.expr.fsm {
	case eFsmInit:
		return me.scanOnFsmInit(tokens)
	case eFsmAtomic:
		return me.scanOnFsmAtomic(tokens)
	case eFsmSingle:
		return me.scanOnFsmSingle(tokens)
	case eFsmSingleOk:
		return me.scanOnFsmSingleOk(tokens)
	case eFsmExpr:
		return me.scanOnFsmExpr(tokens)
	case eFsmExprMulti:
		return me.scanOnFsmExprMulti(tokens)
	case eFsmExprOk:
		return me.scanOnFsmExprOk(tokens)
	case eFsmExprMultiMore:
		return me.scanOnFsmExprMultiMore(tokens)
	default:
		Panic("scan on fsm[%d]", me.expr.fsm)

		return tokens
	}
}

func (me *lex3) scanOnFsmInit(tokens []*Token) []*Token {
	var token *Token
	token, tokens = pickFirstToken(tokens)

	switch token.t {
	case TypeSingle:
		// []
		// [] + SINGLE
		//    ++>
		// [SINGLE]
		me.expr.setFsm(eFsmSingle)
		me.expr.logic = token.Logic()
	case TypeMulti:
		// []
		// [] + MULTI
		// [MULTI]
		// !!! ERROR !!!
		me.TokenPanic(token)
	case TypeExprAtomic, TypeExprRaw:
		// []
		// [] + m-expr
		//    ++>
		// [expr]
		me.expr.setFsm(eFsmExpr)
		me.expr.pushExpr(token.expr())
	default:
		me.TokenPanic(token)
	}

	return tokens
}

func (me *lex3) scanOnFsmAtomic(tokens []*Token) []*Token {
	var token *Token
	token, tokens = pickFirstToken(tokens)

	me.TokenPanic(token)

	return tokens
}

func (me *lex3) scanOnFsmSingle(tokens []*Token) []*Token {
	var token *Token

	token, tokens = pickFirstToken(tokens)

	switch token.t {
	case TypeSingle:
		// [SINGLE]
		// [SINGLE] + SINGLE

		// !!! ERROR !!!
		me.TokenPanic(token)
	case TypeMulti:
		// [SINGLE]
		// [SINGLE] + MULTI

		// !!! ERROR !!!
		me.TokenPanic(token)
	case TypeExprAtomic, TypeExprRaw:
		// [SINGLE]
		// [SINGLE] + m-expr
		//    ++>
		// [expr]
		me.expr.setFsm(eFsmSingleOk)
		me.expr.pushExpr(token.expr())
	default:
		me.TokenPanic(token)
	}

	return tokens
}

func (me *lex3) scanOnFsmSingleOk(tokens []*Token) []*Token {
	var token *Token
	token, tokens = pickFirstToken(tokens)

	me.TokenPanic(token)

	return tokens
}

func (me *lex3) scanOnFsmExpr(tokens []*Token) []*Token {
	var expr *Expr
	var token *Token

	token, tokens = pickFirstToken(tokens)

	switch token.t {
	case TypeSingle:
		// [expr]
		// [expr] + SINGLE
		// [expr] + AND + SINGLE
		// [expr] + AND + [SINGLE + next(expr)]
		//    ==>
		// [expr] + AND + expr
		//     ++>
		// [expr AND expr]
		expr, tokens = pickFirstExpr(tokens)

		expr1 := newExprSingleOk(expr)

		me.expr.setFsm(eFsmExprOk)
		me.expr.logic = LogicAnd
		me.expr.pushExpr(expr1)
	case TypeMulti:
		// [expr]
		// [expr] + MULTI
		me.expr.setFsm(eFsmExprMulti)
		me.expr.logic = token.Logic()
	case TypeExprAtomic, TypeExprRaw:
		// [expr]
		// [expr] + m-expr
		// [expr] + AND + m-expr
		//    ++>
		// [expr + AND + m-expr]
		me.expr.setFsm(eFsmExprOk)
		me.expr.logic = LogicAnd
		me.expr.pushExpr(token.expr())
	default:
		me.TokenPanic(token)
	}

	return tokens
}

func (me *lex3) scanOnFsmExprMulti(tokens []*Token) []*Token {
	var expr *Expr
	var token *Token

	token, tokens = pickFirstToken(tokens)

	switch token.t {
	case TypeSingle:
		// [expr + MULTI]
		// [expr + MULTI] + SINGLE
		// [expr + MULTI] + [SINGLE + next(expr)]
		//    ==>
		// [expr + MULTI] + expr
		//    ++>
		// [expr + MULTI + expr]
		expr, tokens = pickFirstExpr(tokens)
		expr1 := newExprSingleOk(expr)

		me.expr.setFsm(eFsmExprOk)
		me.expr.pushExpr(expr1)
	case TypeMulti:
		// [expr + MULTI]
		// [expr + MULTI] + MULTI
		me.TokenPanic(token)
	case TypeExprAtomic, TypeExprRaw:
		// [expr + MULTI]
		// [expr + MULTI] + m-expr
		//    ++>
		// [expr + MULTI + expr]
		me.expr.setFsm(eFsmExprOk)
		me.expr.pushExpr(token.expr())
	default:
		me.TokenPanic(token)
	}

	return tokens
}

func (me *lex3) scanOnFsmExprOk(tokens []*Token) []*Token {
	var expr *Expr
	var token *Token

	token, tokens = pickFirstToken(tokens)

	switch token.t {
	case TypeSingle:
		// [expr + MULTI + expr]
		// [expr + MULTI + expr] + SINGLE
		// [expr + MULTI + expr] + AND + SINGLE
		expr, tokens = pickFirstExpr(tokens)
		if me.expr.logic == LogicAnd {
			// [expr + AND + expr] + AND + [SINGLE + next(expr)]
			//    ==>
			// [expr + AND + expr] + AND + expr
			//    ++>
			// [expr + AND + expr]
			expr1 := newExprSingleOk(expr)
			me.expr.pushExpr(expr1)
		} else {
			// [expr + OR  + expr] + AND + SINGLE
			// [expr + OR  + expr] + AND + [SINGLE + next(expr)]
			//    ==>
			// [expr + OR  + expr] + AND + expr
			//    ==>
			// [expr AND expr]
			expr1 := newExprSingleOk(expr)
			expr2 := me.expr
			me.expr = newAndExpr(expr2, expr1)
		}
	case TypeMulti:
		// [expr + MULTI + expr]
		// [expr + MULTI + expr] + MULTI
		if me.expr.logic == token.Logic() {
			// [expr + MULTI + expr] + MULTI(same)
			//    ++>
			// [expr + MULTI + expr + MULTI]
			me.expr.fsm = eFsmExprMultiMore
		} else {
			// [expr + MULTI + expr] + MULTI(diff)
			//    ==>
			// expr + MULTI(diff)
			expr := me.expr

			me.expr = newExpr(eFsmExprMulti)
			me.expr.logic = token.Logic()
			me.expr.pushExpr(expr)
		}
	case TypeExprAtomic, TypeExprRaw:
		// [expr + MULTI + expr]
		// [expr + MULTI + expr] + m-expr
		// [expr + MULTI + expr] + AND + m-expr
		if me.expr.logic == LogicAnd {
			// [expr + AND + expr] + AND + m-expr
			// ++>
			// [expr + AND + expr]
			me.expr.pushExpr(token.expr())
		} else {
			// [expr + OR + expr] + AND + m-expr
			//    ==>
			// expr + AND + m-expr
			me.expr = newAndExpr(me.expr, token.expr())
		}
	default:
		me.TokenPanic(token)
	}

	return tokens
}

func (me *lex3) scanOnFsmExprMultiMore(tokens []*Token) []*Token {
	var expr *Expr
	var token *Token

	token, tokens = pickFirstToken(tokens)

	switch token.t {
	case TypeSingle:
		// [expr + MULTI + expr + MULTI]
		// [expr + MULTI + expr + MULTI] + SINGLE
		// [expr + MULTI + expr + MULTI] + SINGLE + next(m-expr)
		// [expr + MULTI + expr + MULTI] + expr
		// [expr + MULTI + expr]
		expr, tokens = pickFirstExpr(tokens)

		me.expr.setFsm(eFsmExprOk)
		me.expr.pushExpr(expr)
	case TypeMulti:
		// [expr + MULTI + expr + MULTI]
		// [expr + MULTI + expr + MULTI] + MULTI
		// !!! ERROR !!!
		me.TokenPanic(token)
	case TypeExprAtomic, TypeExprRaw:
		// [expr + MULTI + expr + MULTI]
		// [expr + MULTI + expr + MULTI] + m-expr
		//    ++>
		// [expr + MULTI + expr]
		me.expr.setFsm(eFsmExprOk)
		me.expr.pushExpr(token.expr())
	default:
		me.TokenPanic(token)
	}

	return tokens
}
