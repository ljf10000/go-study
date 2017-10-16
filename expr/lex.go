package libexpr

import (
	. "asdf"
)

const (
	SINGLE_QUOT = '\''
	DOUBLE_QUOT = '"'
)

type lex struct {
	line      string
	level     int
	quot      rune
	token     *Token
	tkComplet bool
	atomic    *Atomic
	expr      *Expr
	root      *Expr
	stack     []*Expr
}

func newLex(line string) *lex {
	return &lex{
		line: line,
	}
}

func (me *lex) flush() {
	// todo
}

func (me *lex) Token(v string) *Token {
	if Empty != v {
		return &Token{
			typ:   TypeString,
			value: v,
			level: me.level,
		}
	} else {
		return nil
	}
}

func (me *lex) newToken(v string) *Token {
	return &Token{
		typ:   TypeString,
		value: v,
		level: me.level,
	}
}

func (me *lex) buildinToken(buildin Buildin) *Token {
	return &Token{
		typ:   buildin.Type(),
		value: buildin,
		level: me.level,
	}
}

func (me *lex) openToken() {
	me.token = me.newToken(Empty)
	me.tkComplet = false
}

func (me *lex) closeToken() {
	me.tkComplet = true
}

func (me *lex) openQuot(c rune) {
	me.closeToken()
	me.quot = c

	Log.Info("open quot")
}

func (me *lex) closeQuot() {
	me.closeToken()
	me.quot = 0

	Log.Info("close quot")
}

func (me *lex) openPar() {
	me.closeToken()
	me.level++

	Log.Info("open %d:quot", me.level)

	// todo
}

func (me *lex) closePar() {
	me.closeToken()
	// todo

	Log.Info("close %d:quot", me.level)
	me.level--
}

func (me *lex) pushBuildin(buildin Buildin) {
	token := me.buildinToken(buildin)

	me.pushToken(token)
}

func (me *lex) pushToken(token *Token) {
	me.closeToken()

	if nil != token {
		if nil != me.token {
			// todo
		} else {
			me.token = token
			me.closeToken()
		}
	}
}

func (me *lex) pushChar(c rune) {
	if nil == me.token {
		me.openToken()
	}

	if !me.tkComplet {
		me.token.PushChar(c)
	} else {
		// todo
	}
}

func (me *lex) stop() {
	// end scan
	me.flush()

	// check ' or " close
	switch me.quot {
	case SINGLE_QUOT:
		panic(`not close '`)
	case DOUBLE_QUOT:
		panic(`not close "`)
	}
}

func (me *lex) start() {
	line := me.line

	for len(line) > 0 {
		line = me.scan(line)
	}

	me.stop()
}

func (me *lex) scan(line string) string {
	if 0 == len(line) {
		return Empty
	} else if v, ok := hasBuildinPrefix(line); ok {
		return me.scanBuildin(v, line)
	} else {
		return me.scanNormal(line)
	}
}

func (me *lex) scanBuildin(buildin Buildin, line string) string {
	switch buildin {
	case BuildinLp:
		me.openPar()

		return line[1:]
	case BuildinRp:
		me.closePar()

		return line[1:]
	default:
		me.pushBuildin(buildin)

		return line[len(buildin.String()):]
	}
}

func (me *lex) scanNormal(line string) string {
	c, line := FirstRune(line)

	switch c {
	case ' ', '\t':
		// skip space
		me.closeToken()
	case SINGLE_QUOT:
		Log.Info("begin single quot")
		line = me.scanQuot(c, line, me.singleQuot)
	case DOUBLE_QUOT:
		Log.Info("begin double quot")
		line = me.scanQuot(c, line, me.doubleQuot)
	default:
		me.pushChar(c)

		line = me.scan(line)
	}

	return line
}

type QuotScan func(line string) (bool, string)

func (me *lex) scanQuot(quot rune, line string, scan QuotScan) string {
	closed := false

	me.openQuot(quot)

	for {
		closed, line = scan(line)
		if closed {
			return line
		}
	}
}

func (me *lex) singleQuot(line string) (bool, string) {
	closed := false

	var c rune
	c, line = FirstRune(line)

	switch c {
	case me.quot:
		me.closeQuot()
		closed = true
	default:
		me.pushChar(c)
	}

	return closed, line
}

func (me *lex) doubleQuot(line string) (bool, string) {
	return me.singleQuot(me.escape(line))
}

func (me *lex) escape(line string) string {
	if s, c, ok := hasEscapePrefix(line); ok {
		me.pushChar(c)

		line = line[len(s):]
	}

	return line
}

func Scan(line string) {
	lex := &lex{
		line: line,
	}

	lex.start()
}
