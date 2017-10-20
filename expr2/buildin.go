package libexpr

import (
	. "asdf"
	"strings"
)

// must keep sort
const (
	BdBegin Buildin = 1 // begin 1

	OpBegin   Buildin = BdBegin
	OpGeEq    Buildin = OpBegin
	OpGe      Buildin = 2
	OpLeEq    Buildin = 3
	OpLe      Buildin = 4
	OpEq      Buildin = 5
	OpInclude Buildin = 6
	OpNeq     Buildin = 7
	OpEnd     Buildin = 8

	LogicMultiBegin Buildin = OpEnd
	LogicAnd        Buildin = LogicMultiBegin
	LogicOr         Buildin = 9
	LogicMultiEnd   Buildin = 10

	LogicSingleBegin Buildin = LogicMultiEnd
	LogicNot         Buildin = LogicSingleBegin
	LogicSingleEnd   Buildin = 11

	BdEnd Buildin = LogicSingleEnd
)

var exprBuildins = [BdEnd]string{
	OpGeEq:    ">=",
	OpGe:      ">",
	OpLeEq:    "<=",
	OpLe:      "<",
	OpEq:      "==",
	OpInclude: "=",
	OpNeq:     "!=",
	LogicAnd:  "&&",
	LogicOr:   "||",
	LogicNot:  "!",
}

var exprBuildinTypes = [BdEnd]Type{
	OpGeEq:    TypeOperator,
	OpGe:      TypeOperator,
	OpLeEq:    TypeOperator,
	OpLe:      TypeOperator,
	OpEq:      TypeOperator,
	OpInclude: TypeOperator,
	OpNeq:     TypeOperator,
	LogicAnd:  TypeMulti,
	LogicOr:   TypeMulti,
	LogicNot:  TypeSingle,
}

var exprBuildinKeywords = map[string]*Keyword{}

type Buildin int
type Logic = Buildin
type Op = Buildin

func (me Buildin) IsGood() bool {
	return me >= BdBegin && me < BdEnd
}

func (me Buildin) String() string {
	if me.IsGood() {
		return exprBuildins[me]
	} else {
		return Unknow
	}
}

func (me Buildin) Type() Type {
	if me.IsGood() {
		return exprBuildinTypes[me]
	} else {
		return TypeValue
	}
}

func (me Buildin) IsOp() bool {
	return TypeOperator == me.Type()
}

func (me Buildin) IsSingle() bool {
	return TypeSingle == me.Type()
}

func (me Buildin) IsMulti() bool {
	return TypeMulti == me.Type()
}

func hasBuildinPrefix(s string) (Buildin, bool) {
	idx, has := HasPrefix(s, exprBuildins[BdBegin:])

	return BdBegin + Buildin(idx), has
}

func hasBuildinKeywordPrefix(s string) (*Keyword, bool) {
	for k, v := range exprBuildinKeywords {
		if strings.HasPrefix(s, k) {
			return v, true
		}
	}

	return nil, false
}
