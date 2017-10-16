package libexpr

/*
import (
	. "asdf"
	"fmt"
)

type Term struct {
	Level  int
	Tokens []*Token
}

func (me *Term) String() string {
	s := Empty
	count := len(me.Tokens)

	for k, v := range me.Tokens {
		s += v.RawString()
		if k < count-1 {
			s += " "
		}
	}

	return s
}

func (me *Term) pushToken(token *Token) {
	me.Tokens = append(me.Tokens, token)

	Log.Info("push term token:%s", token.String())
}

func (me *Term) popToken() {
	if len(me.Tokens) > 0 {
		me.Tokens = []*Token{}

		Log.Info("pop term token")
	}
}

func (me *lex) popTerm() {
	if len(me.term.Tokens) > 0 {
		me.terms = append(me.terms, me.term)
	}

	me.term.popToken()
}

func (me *lex) ScanTokens() {
	tokens := me.tokens

	for len(tokens) > 0 {
		tokens = me.scanTokens(tokens)
	}

	me.popTerm()
}

func (me *lex) scanTokens(tokens []*Token) []*Token {
	token := tokens[0]
	tokens = tokens[1:]

	if token.IsLp() {
		me.popTerm()
		me.term.pushToken(token)

		for k, v := range tokens {
			me.term.pushToken(v)

			if v.L == token.L && v.IsRp() {
				me.popTerm()

				return tokens[k+1:]
			}
		}

		Panic("not close ():%d", token.L)
	} else {
		me.term.pushToken(token)
		me.popTerm()
	}

	return tokens
}

func (me *lex) DumpTerms() string {
	s := ""

	s += fmt.Sprintf("Line:%s\n", me.line)
	s += "Terms:\n"
	for _, v := range me.terms {
		s += fmt.Sprintf("\t%s\n", v.String())
	}

	return s
}

func TokenScan(line string) []*Token {
	lex := &lexToken{
		line:   line,
		tokens: []*Token{},
	}

	lex.start()
	Log.Info(lex.DumpTokens())

	return lex.tokens
}
*/
