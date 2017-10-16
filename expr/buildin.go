package libexpr

import (
	. "asdf"
)

// must keep sort
const (
	BuildinUnknow Buildin = 0

	BuildinBegin         Buildin = 1
	BuildinOperatorBegin Buildin = BuildinBegin
	BuildinEqGe          Buildin = BuildinOperatorBegin
	BuildinGe            Buildin = 2
	BuildinEqLe          Buildin = 3
	BuildinLe            Buildin = 4
	BuildinEq            Buildin = 5
	BuildinInclude       Buildin = 6
	BuildinNeq           Buildin = 7
	BuildinOperatorEnd   Buildin = 8

	BuildinLogicBegin Buildin = BuildinOperatorEnd
	BuildinAnd        Buildin = BuildinLogicBegin
	BuildinOr         Buildin = 9
	BuildinNot        Buildin = 10
	BuildinLogicEnd   Buildin = 11

	BuildinKeyWordBegin Buildin = BuildinLogicEnd
	BuildinLp           Buildin = BuildinKeyWordBegin
	BuildinRp           Buildin = 12
	BuildinKeyWordEnd   Buildin = 13

	BuildinEnd Buildin = BuildinKeyWordEnd
)

var exprBuildins = [BuildinEnd]string{
	BuildinUnknow:  "unknow",
	BuildinEqGe:    ">=",
	BuildinGe:      ">",
	BuildinEqLe:    "<=",
	BuildinLe:      "<",
	BuildinEq:      "==",
	BuildinInclude: "=",
	BuildinNeq:     "!=",
	BuildinAnd:     "&&",
	BuildinOr:      "||",
	BuildinNot:     "!",
	BuildinLp:      "(",
	BuildinRp:      ")",
}

type Buildin int

func (me Buildin) String() string {
	return exprBuildins[me]
}

func (me Buildin) Type() Type {
	switch {
	case me >= BuildinOperatorBegin && me < BuildinOperatorEnd:
		return TypeOperator
	case me >= BuildinLogicBegin && me < BuildinLogicEnd:
		return TypeLogic
	case me >= BuildinKeyWordBegin && me < BuildinKeyWordEnd:
		return TypeKeyWord
	default:
		return TypeUnknow
	}
}

func hasBuildinPrefix(s string) (Buildin, bool) {
	idx, has := HasPrefix(s, exprBuildins[BuildinBegin:BuildinEnd])

	return Buildin(idx), has
}
