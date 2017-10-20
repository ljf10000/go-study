package libexpr

import (
	. "asdf"
)

func lex2Scan(tokens []*Token) []*Token {
	lex := newLex2(tokens)

	lex.start()
	lex.stop()
	Log.Info(lex.DumpString())
	return lex.tokens
}

func newLex2(tokens []*Token) *lex2 {
	return &lex2{
		raw:    tokens,
		tokens: []*Token{},
		atomic: newAtomic(),
	}
}

type lex2 struct {
	raw    []*Token
	tokens []*Token
	atomic *Atomic
}

func (me *lex2) DumpString() string {
	s := "lex2:" + Crlf
	for _, token := range me.tokens {
		s += Tab + token.String() + Crlf
	}

	return s
}

func (me *lex2) TokenPanic(token *Token) {
	Panic("add token[%s] on fsm[%s]", token, me.atomic.fsm)
}

func (me *lex2) pushAtomic() {
	me.pushToken(&Token{
		t: TypeExprAtomic,
		v: newAtomicExpr(me.atomic),
	})

	me.atomic = newAtomic()
}

func (me *lex2) pushLogicAnd() {
	me.pushToken(&Token{
		t: TypeMulti,
		v: LogicAnd,
	})
}

func (me *lex2) pushToken(token *Token) {
	Log.Info("push token:%s", token)
	me.tokens = append(me.tokens, token)
}

func (me *lex2) start() {
	tokens := me.raw

	for len(tokens) > 0 {
		tokens = me.scan(tokens)
	}
}

func (me *lex2) stop() {
	atomic := me.atomic

	switch atomic.fsm {
	case aFsmValue:
		atomic.toDeft()
		me.pushAtomic()
	case aFsmInit:
		DoNothing()
	default:
		Panic("left fsm[%s]", atomic.fsm)
	}
}

func (me *lex2) scan(tokens []*Token) []*Token {
	if len(tokens) == 0 {
		return tokens
	}

	switch me.atomic.fsm {
	case aFsmInit:
		return me.scanOnFsmInit(tokens)
	case aFsmKey:
		return me.scanOnFsmKey(tokens)
	case aFsmValue:
		return me.scanOnFsmValue(tokens)
	case aFsmKeyOp:
		return me.scanOnFsmKeyOp(tokens)
	case aFsmOk:
		Panic("scan on fsm[%d]", me.atomic.fsm)
	}

	return tokens
}

func (me *lex2) scanOnFsmInit(tokens []*Token) []*Token {
	var token *Token
	token, tokens = pickFirstToken(tokens)

	atomic := me.atomic

	switch token.t {
	case TypeKeyWord:
		// 0. []
		// 1. [] + keyword
		// 2. keyword
		atomic.k = token.Keyword()
		atomic.setFsm(aFsmKey)
	case TypeValue:
		// 0. []
		// 1. [] + value
		// 2. value
		atomic.v = token.Value()
		atomic.setFsm(aFsmValue)
	case TypeSingle:
		// 0. []
		// 1. [] + SINGLE
		//    add SINGLE
		// 2. []
		me.pushToken(token)
	case TypeMulti:
		// 0. []
		// 1. [] + MULTI
		//    add MULTI
		// 2. []
		if len(me.tokens) > 0 {
			me.pushToken(token)
		} else {
			Panic("first is %s", token)
		}
	case TypeExprRaw:
		// 0. []
		// 1. [] + EXPR
		//    add EXPR
		// 2. []
		me.pushToken(token)
	default:
		me.TokenPanic(token)
	}

	return tokens
}

func (me *lex2) scanOnFsmKey(tokens []*Token) []*Token {
	var token *Token
	token, tokens = pickFirstToken(tokens)

	atomic := me.atomic

	switch token.t {
	case TypeOperator:
		// 0. keyword
		// 1. keyword + op
		atomic.op = token.Buildin()
		atomic.setFsm(aFsmKeyOp)
	default:
		me.TokenPanic(token)
	}

	return tokens
}
func (me *lex2) scanOnFsmValue(tks []*Token) []*Token {
	token, tokens := pickFirstToken(tks)

	atomic := me.atomic

	switch token.t {
	case TypeValue:
		// 0. value
		// 1. value + value
		// 2. ALL INCLUDE value + AND + value
		// 3. atomic AND value
		atomic.toDeft()

		me.pushAtomic()
		me.pushLogicAnd()

		tokens = me.scan(tks)
	case TypeKeyWord:
		// 0. value
		// 1. value + keyword
		// 2. ALL INCLUDE value + AND + keyword
		// 3. atomic AND keyword
		atomic.toDeft()

		me.pushAtomic()
		me.pushLogicAnd()

		tokens = me.scan(tks)
	case TypeSingle:
		// 0. value
		// 1. value + SINGLE
		// 2. ALL INCLUDE value + AND + SINGLE
		// 3. atomic AND SINGLE
		atomic.toDeft()

		me.pushAtomic()
		me.pushLogicAnd()

		tokens = me.scan(tks)
	case TypeMulti:
		// 0. value
		// 1. value + SINGLE
		// 2. ALL INCLUDE value + MULTI
		// 3. atomic + MULTI
		atomic.toDeft()

		me.pushAtomic()
		me.pushToken(token)
	case TypeExprRaw:
		// 0. value
		// 1. value + EXPR
		// 2. ALL INCLUDE value + AND + EXPR
		// 3. atomic AND EXPR
		atomic.toDeft()

		me.pushAtomic()
		me.pushLogicAnd()

		tokens = me.scan(tks)
	default:
		me.TokenPanic(token)
	}

	return tokens
}

func (me *lex2) scanOnFsmKeyOp(tokens []*Token) []*Token {
	var token *Token
	token, tokens = pickFirstToken(tokens)

	atomic := me.atomic

	switch token.t {
	case TypeValue:
		// 0. keyword + op
		// 1. keyword + op + value
		//    push atomic
		// 2. []
		atomic.v = token.Value()
		atomic.setFsm(aFsmOk)

		me.pushAtomic()
	default:
		me.TokenPanic(token)
	}

	return tokens
}
