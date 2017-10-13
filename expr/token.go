package libexpr

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

func (me *Token) Key() Key {
	return me.value.(Key)
}

func (me *Token) Value() string {
	return me.value.(string)
}
