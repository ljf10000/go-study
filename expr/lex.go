package libexpr

import (
	. "asdf"
	"errors"
	"fmt"
	"unicode/utf8"
)

type Term []Token

type Lex struct {
	line  string
	token string
	term  Term
	terms []Term
	quot  rune
	level int
}

func NewLex(line string) *Lex {
	return &Lex{
		line: line,
		term: Term{},
	}
}

func (me *Lex) DumpToken() string {
	s := ""

	s += fmt.Sprintf("Line:%s\n", me.line)
	s += "Tokens:\n"
	for _, token := range me.term {
		s += fmt.Sprintf("\t%s\n", token.value)
	}

	return s
}

func (me *Lex) TokenScan() error {
	return nil
}

func (me *Lex) tokenScan(term Term) (Term, error) {
	count := len(term)
	if 0 == count {

	} else if 1 == count {

	} else {

	}

	return nil, nil
}

func (me *Lex) LineScan() error {
	for line := me.line; Empty != line; {
		line = me.lineHandle(line)
	}

	return me.endLineScan()
}

func (me *Lex) endLineScan() error {
	// check ' or " close
	switch me.quot {
	case '\'':
		if 0 != me.quot {
			return errors.New(`not close '`)
		}
	case '"':
		if 0 != me.quot {
			return errors.New(`not close "`)
		}
	}

	// end scan
	me.pushToken()

	return nil
}

func (me *Lex) lineHandle(line string) string {
	switch me.quot {
	case 0:
		return me.lineNormalHandle(line)
	case '\'':
		return me.lineSingleQuotHandle(line)
	case '"':
		return me.lineDoubleQuotHandle(line)
	default:
		return line
	}
}

func (me *Lex) lineNormalHandle(line string) string {
	c, _ := utf8.DecodeRuneInString(line)
	Log.Info("handle normal line:%s", line)

	Len := 0
	if v, ok := hasOpPrefix(line); ok {
		Len = len(v.String())

		me.pushOp(v)
	} else if v, ok := hasLogicPrefix(line); ok {
		Len = len(v.String())

		me.pushLogic(v)
	} else if v, ok := hasKeyPrefix(line); ok {
		Len = len(v.String())

		me.pushKey(v)
	} else {
		Len = len(string(c))

		switch c {
		case ' ', '\t':
			me.skipChar(c)
			me.pushToken()
		case '"', '\'':
			me.skipChar(c)
			me.quot = c
			me.token = Empty

			Log.Info("begin quot:%c", c)
		default:
			me.pushChar(c)
		}
	}

	return line[Len:]
}

func (me *Lex) lineSingleQuotHandle(line string) string {
	c, _ := utf8.DecodeRuneInString(line)
	Log.Info("handle single quot line:%s", line)

	switch c {
	case me.quot:
		me.closeQuot()
	default:
		me.pushChar(c)
	}

	return line[len(string(c)):]
}

func (me *Lex) escape(line string) string {
	if s, e, ok := hasEscapePrefix(line); ok {
		me.pushChar(e)

		line = line[len(s):]
	}

	return line
}

func (me *Lex) lineDoubleQuotHandle(line string) string {
	return me.lineSingleQuotHandle(me.escape(line))
}

func (me *Lex) closeQuot() {
	// close quot
	me.quot = 0
	me.pushToken()

	Log.Info("end quot")
}

func (me *Lex) pushToken() {
	if Empty != me.token {
		Log.Info("save auto token:%s", me.token)
		me.term = append(me.term, Token{
			Type:  TypeString,
			value: me.token,
		})
	}

	me.token = Empty
}

func (me *Lex) pushOp(token Op) {
	me.pushToken()

	Log.Info("save Op token:%s", token)
	me.term = append(me.term, Token{
		Type:  TypeOp,
		value: token,
	})
}

func (me *Lex) pushLogic(token Logic) {
	me.pushToken()

	Log.Info("save Logic token:%s", token)
	me.term = append(me.term, Token{
		Type:  TypeLogic,
		value: token,
	})
}

func (me *Lex) pushKey(token Key) {
	me.pushToken()

	Log.Info("save Key token:%s", token)
	me.term = append(me.term, Token{
		Type:  TypeKey,
		value: token,
	})
}

func (me *Lex) skipChar(c rune) {
	// do nothing
}

func (me *Lex) pushChar(c rune) {
	me.token += string(c)
}
