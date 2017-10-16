package libexpr

import (
	. "asdf"
	"fmt"
)

type Token struct {
	typ   Type
	level int
	value interface{}
}

func (me *Token) RawString() string {
	switch me.typ {
	case TypeTerm:
		panic("invalid term token type")
		return Empty
	case TypeObject:
		return me.Object().String()
	case TypeValue, TypeString:
		return me.Value()
	case TypeLogic, TypeOperator, TypeKeyWord:
		return me.Buildin().String()
	default:
		return BuildinUnknow.String()
	}
}

func (me *Token) String() string {
	return fmt.Sprintf("%s[%d:%s]", me.RawString(), me.level, me.typ)
}

func (me *Token) PushChar(c rune) {
	me.value = me.Value() + string(c)
}

func (me *Token) Buildin() Buildin {
	return me.value.(Buildin)
}

func (me *Token) Value() string {
	return me.value.(string)
}

func (me *Token) Object() *Object {
	return me.value.(*Object)
}

func (me *Token) IsLp() bool {
	return me.typ == TypeKeyWord && me.Buildin() == BuildinLp
}

func (me *Token) IsRp() bool {
	return me.typ == TypeKeyWord && me.Buildin() == BuildinRp
}

type lexToken struct {
	line   string
	token  *Token
	tokens []*Token
	quot   rune
	level  int
}
