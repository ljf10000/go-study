package libexpr

import (
	. "asdf"
)

const (
	TypeValue      Type = 0
	TypeOperator   Type = 1
	TypeKeyWord    Type = 2
	TypeSingle     Type = 3
	TypeMulti      Type = 4
	TypeExprAtomic Type = 5
	TypeExprRaw    Type = 6
	TypeEnd        Type = 7
)

var exprTypes = [TypeEnd]string{
	TypeValue:      "value",
	TypeOperator:   "operator",
	TypeKeyWord:    "keyword",
	TypeSingle:     "single",
	TypeMulti:      "multi",
	TypeExprAtomic: "a-expr",
	TypeExprRaw:    "r-expr",
}

type Type int

func (me Type) IsGood() bool {
	return me >= 0 && me < TypeEnd
}

func (me Type) String() string {
	if me.IsGood() {
		return exprTypes[me]
	} else {
		return Unknow
	}
}
