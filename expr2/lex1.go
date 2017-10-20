package libexpr

import (
	. "asdf"
)

const (
	CHAR_QUOT_SINGLE = '\''
	CHAR_QUOT_DOUBLE = '"'
	CHAR_SPACE       = ' '
	CHAR_TAB         = '\t'
	CHAR_CRLF        = '\n'
	CHAR_LP          = '('
	CHAR_RP          = ')'
)

func lex1Scan(line string) []*Token {
	lex := newLex1(line)

	Log.Info("scan line:[%s] ...", line)
	lex.start()
	lex.stop()
	Log.Info(lex.DumpString())
	Log.Info("scan line:[%s] ok.", line)
	Log.Info("======================================================================")

	return lex.tokens
}

func newLex1(line string) *lex1 {
	return &lex1{
		raw:    line,
		tokens: []*Token{},
	}
}

type lex1 struct {
	raw    string
	quot   rune
	level  int
	token  string
	tokens []*Token
}

func (me *lex1) DumpString() string {
	s := "lex1:" + Crlf
	for _, token := range me.tokens {
		s += Tab + token.String() + Crlf
	}

	return s
}

func (me *lex1) flushToken() {
	if Empty != me.token {
		me.addString(me.token)

		me.token = Empty
	}
}

func (me *lex1) flushExpr() {
	if Empty != me.token {
		me.addExpr(me.token)

		me.token = Empty
	}
}

func (me *lex1) skipChar(c rune) {
	Log.Info("skip [%s]", string(c))

	me.flushToken()
}

func (me *lex1) addChar(c rune) {
	s := string(c)
	token := me.token

	me.token = token + s

	Log.Info("add char[%s] ==> token[%s] == token[%s]",
		s, token, me.token)
}

func (me *lex1) addToken(token *Token) {
	Log.Info("add token:%s", token)

	me.tokens = append(me.tokens, token)
}

func (me *lex1) addString(s string) {
	me.addToken(&Token{
		t: TypeValue,
		v: s,
	})
}

func (me *lex1) addBuildin(bd Buildin) {
	me.addToken(&Token{
		t: bd.Type(),
		v: bd,
	})
}

func (me *lex1) addKeyword(keyword *Keyword) {
	me.addToken(&Token{
		t: TypeKeyWord,
		v: keyword,
	})
}

func (me *lex1) addExpr(term string) {
	me.addToken(&Token{
		t: TypeExprRaw,
		v: term,
	})
}

func (me *lex1) openPa(strict bool) {
	if strict {
		me.skipChar(CHAR_LP)
	} else {
		me.addChar(CHAR_LP)
	}

	me.level++
	Log.Info("open %d(", me.level)
}

func (me *lex1) closePa() int {
	Log.Info("close %d)", me.level)
	me.level--

	if me.level > 0 {
		me.addChar(CHAR_RP)
	}

	return me.level
}

func (me *lex1) openQuot(c rune, strict bool) {
	me.quot = c
	if strict {
		me.flushToken()
	}

	Log.Info("open quot strict:%v", strict)
}

func (me *lex1) closeQuot(strict bool) {
	me.quot = 0
	if strict {
		me.flushToken()
	}

	Log.Info("close quot strict:%v", strict)
}

func (me *lex1) stop() {
	me.flushToken()

	// check ' or " close
	switch me.quot {
	case CHAR_QUOT_SINGLE:
		panic(`not close '`)
	case CHAR_QUOT_DOUBLE:
		panic(`not close "`)
	}

	if me.level > 0 {
		Panic(`not close )%d`, me.level)
	}
}

func (me *lex1) start() {
	line := me.raw

	for len(line) > 0 {
		line = me.scan(line)
	}
}

func (me *lex1) scan(line string) string {
	if v, ok := hasBuildinPrefix(line); ok {
		me.flushToken()
		me.addBuildin(v)

		return line[len(v.String()):]
	} else if v, ok := hasBuildinKeywordPrefix(line); ok {
		me.flushToken()
		me.addKeyword(v)

		return line[len(v.Key):]
	} else {
		return me.scanNormal(line)
	}
}

func (me *lex1) scanNormal(line string) string {
	c, line := FirstRune(line)

	switch c {
	case CHAR_SPACE, CHAR_TAB, CHAR_CRLF:
		me.skipChar(c)
	case CHAR_QUOT_SINGLE:
		line = me.scanQuot(c, line, true, me.singleQuot)
	case CHAR_QUOT_DOUBLE:
		line = me.scanQuot(c, line, true, me.doubleQuot)
	case CHAR_LP:
		line = me.scanTerm(line)
	default:
		me.addChar(c)
	}

	return line
}

func (me *lex1) scanTerm(line string) string {
	var c rune

	me.openPa(true)

	for len(line) > 0 {
		c, line = FirstRune(line)

		switch c {
		case CHAR_QUOT_SINGLE:
			line = me.scanQuot(c, line, false, me.singleQuot)
		case CHAR_QUOT_DOUBLE:
			line = me.scanQuot(c, line, false, me.doubleQuot)
		case CHAR_LP:
			me.openPa(false)
		case CHAR_RP:
			if 0 == me.closePa() {
				goto CLOSE
			}
		default:
			me.addChar(c)
		}
	}

CLOSE:
	me.flushExpr()

	return line
}

type funcLineLexScan func(line string, strict bool) (bool, string)

func (me *lex1) scanQuot(quot rune, line string, strict bool, scan funcLineLexScan) string {
	closed := false
	me.openQuot(quot, strict)

	for {
		closed, line = scan(line, strict)
		if closed {
			break
		}
	}

	me.closeQuot(strict)
	return line
}

func (me *lex1) singleQuot(line string, strict bool) (bool, string) {
	var c rune

	c, line = FirstRune(line)
	if c == me.quot {
		return true, line
	} else {
		me.addChar(c)

		return false, line
	}
}

func (me *lex1) doubleQuot(line string, strict bool) (bool, string) {
	return me.singleQuot(me.escape(line), strict)
}

func (me *lex1) escape(line string) string {
	if s, c, ok := hasEscapePrefix(line); ok {
		me.addChar(c)

		line = line[len(s):]
	}

	return line
}
