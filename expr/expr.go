package libexpr

type Expr struct {
	father   *Expr
	logic    Buildin
	children []*Atomic
}

func (me *Expr) IsRoot() bool {
	return nil == me.father
}

/*
func (me *Expr) pushToken(token *Token) {
	if TypeUnknow == token.T {
		panic("unknow token type")
	}

	count := len(me.children)

	if 0 == count {
		if me.Type() != TypeUnknow {
			// panic
		}

		switch token.T {
		case TypeKeyWord:
			if token.Tk() == TkLp {

			} else {

			}
		case TypeLogic:
			if token.Tk() == TkNot {
				// first token is !
				// do nothing
			} else {
				panic("push first logic token[!] to unknow expr")
			}
		case TypeOperator:
			Panic("push first operator token[%s] to unknow expr", TypeOperator)
		case TypeString:
		default:
		}
	} else {

	}

	switch me.Type() {
	case TypeUnknow:
		switch token.T {
		case TypeString:
		case TypeLogic:
		case TypeOperator:
		case TypeKeyWord:
		default:
			Panic("push %s token to unknow expr", token.T)
		}
	case TypeOperator:
		// ok
	default:
		Panic("invalid expr type:%s", me.Type())
	}

	me.children = append(me.children, token)
}

func (me *lex) ScanToken2(tokens []Token) *lex {
	for len(tokens) > 0 {
		tokens = me.scanToken2(tokens)
	}

	return me
}

func (me *lex) scanToken2(tokens []Token) []Token {
	token := tokens[0]
	tokens = tokens[1:]

	if nil == me.root {
		me.root = &Expr{
			Tk: TkUnknow,
		}

		me.expr = me.root
	}

	me.expr.pushToken(&token)

	return tokens
}
*/
