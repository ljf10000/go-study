package libexpr

import (
	. "asdf"
	"errors"
	"fmt"
	"unicode/utf8"
)

type Expr struct {
	Left  *Expr
	Rigth *Expr
	Logic Logic

	Op    Op
	Scope Scope
	Key   string
	Value string
}

type Token struct {
	Type  Type
	value interface{}
}

func (me *Token) Op() Op {
	return me.value.(Op)
}

func (me *Token) Logic() Logic {
	return me.value.(Logic)
}

func (me *Token) Value() string {
	return me.value.(string)
}

type Lex struct {
	line   string
	token  string
	tokens []string
	quot   rune
	inQuot bool
}

func (me *Lex) Scan(line string) error {
	me.line = line

	return me.scan()
}

func (me *Lex) dump() string {
	s := ""

	s += fmt.Sprintf("Line:%s\n", me.line)
	s += "Tokens:\n"
	for _, token := range me.tokens {
		s += fmt.Sprintf("\t%s\n", token)
	}

	return s
}

func (me *Lex) scan() error {
	var err error

	for line := me.line; Empty != line; {
		line, err = me.handle(line)
		if nil != err {
			return err
		}
	}
	me.saveToken(Empty)

	return nil
}

func (me *Lex) handle(line string) (string, error) {
	switch me.quot {
	case 0:
		return me.normalHandle(line)
	case '\'':
		return me.singleQuotHandle(line)
	case '"':
		return me.doubleQuotHandle(line)
	default:
		return Empty, errors.New("bad quot")
	}
}

func (me *Lex) normalHandle(line string) (string, error) {
	c, _ := utf8.DecodeRuneInString(line)
	Log.Info("handle normal line:%s", line)

	if v, ok := hasOpPrefix(line); ok {
		me.saveToken(v.String())

		return line[len(v.String()):], nil
	} else if v, ok := hasLogicPrefix(line); ok {
		me.saveToken(v.String())

		return line[len(v.String()):], nil
	} else {
		switch c {
		case ' ', '\t':
			me.skip(c)
			me.saveToken(Empty)
		case '"', '\'':
			me.skip(c)
			me.quot = c
			me.token = Empty

			Log.Info("begin quot:%c", c)
		case '(', ')':
			me.saveToken(string(c))
		default:
			me.save(c)
		}

		return line[len(string(c)):], nil
	}
}

func (me *Lex) singleQuotHandle(line string) (string, error) {
	c, _ := utf8.DecodeRuneInString(line)
	Log.Info("handle single quot line:%s", line)

	if c == me.quot {
		me.closeQuot()
	} else {
		me.save(c)
	}

	return line[len(string(c)):], nil
}

func (me *Lex) doubleQuotHandle(line string) (string, error) {
	c, _ := utf8.DecodeRuneInString(line)
	Log.Info("handle double quot line:%s", line)

	if s, e, ok := hasEscapePrefix(line); ok {
		me.save(e)

		return line[len(s):], nil
	}

	switch c {
	case me.quot:
		me.closeQuot()
	default:
		me.save(c)
	}

	return line[len(string(c)):], nil
}

func (me *Lex) closeQuot() {
	// close quot
	me.quot = 0
	me.saveToken(Empty)

	Log.Info("end quot")
}

func (me *Lex) saveToken(token string) {
	if Empty != me.token {
		Log.Info("save auto token:%s", me.token)
		me.tokens = append(me.tokens, me.token)
	}

	if Empty != token {
		Log.Info("save spec token:%s", token)
		me.tokens = append(me.tokens, token)
	}

	me.token = Empty
}

func (me *Lex) skip(c rune) {
	// do nothing
}

func (me *Lex) save(c rune) {
	me.token += string(c)
}
