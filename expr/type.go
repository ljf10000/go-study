package libexpr

import (
	. "asdf"
	"strings"
)

func hasPrefix(s string, maps []string) (int, bool) {
	count := len(maps)
	for i := 0; i < count; i++ {
		v := maps[i]

		if strings.HasPrefix(s, v) {
			Log.Info("has prefix:%s", v)

			return i, true
		}
	}

	return 0, false
}

//==============================================================================

const (
	TypeString Type = 0
	TypeLogic  Type = 1
	TypeOp     Type = 2
	TypeKey    Type = 3
	TypeEnd    Type = 4
)

type Type int

//==============================================================================

const (
	ScopeAll  Scope = 0
	ScopeZone Scope = 1
	ScopePath Scope = 2
	ScopeEnd  Scope = 3
)

type Scope int

//==============================================================================
const (
	TkOpBegin Tk = 0
	TkEqGe    Tk = TkOpBegin
	TkGe      Tk = 1
	TkEqLe    Tk = 2
	TkLe      Tk = 3
	TkEq      Tk = 4
	TkInc     Tk = 5
	TkNeq     Tk = 6
	TkOpEnd   Tk = 7

	TkLogicBegin Tk = TkOpEnd
	TkAnd        Tk = TkLogicBegin
	TkOr         Tk = 8
	TkNot        Tk = 9
	TkLogicEnd   Tk = 10
	TkLp         Tk = 10
	TkRp         Tk = 11
	TkEnd        Tk = 12
)

type Tk int

const (
	LogicAnd Logic = 0
	LogicOr  Logic = 1
	LogicNot Logic = 2
	LogicEnd Logic = 3
)

var exprLogics = [LogicEnd]string{
	LogicAnd: "&&",
	LogicOr:  "||",
	LogicNot: "!",
}

type Logic int

func (me Logic) String() string {
	return exprLogics[me]
}

func hasLogicPrefix(s string) (Logic, bool) {
	idx, ok := hasPrefix(s, exprLogics[:])

	return Logic(idx), ok
}

//==============================================================================

const (
	OpEqGe    Op = 0 // >=
	OpGe      Op = 1 // >
	OpEqLe    Op = 2 // <=
	OpLe      Op = 3 // <
	OpEq      Op = 4 // ==
	OpInclude Op = 5 // =
	OpNeq     Op = 6 // !=
	OpEnd     Op = 7
)

var exprOperators = [OpEnd]string{
	OpEqGe:    ">=",
	OpGe:      ">",
	OpEqLe:    "<=",
	OpLe:      "<",
	OpEq:      "==",
	OpInclude: "=",
	OpNeq:     "!=",
}

type Op int

func (me Op) String() string {
	return exprOperators[me]
}

func hasOpPrefix(s string) (Op, bool) {
	idx, ok := hasPrefix(s, exprOperators[:])

	return Op(idx), ok
}

//==============================================================================

const (
	KeyLP  Key = 0 // (
	KeyRP  Key = 1 // )
	KeyEnd Key = 2
)

var exprKeys = [KeyEnd]string{
	KeyLP: "(",
	KeyRP: ")",
}

type Key int

func (me Key) String() string {
	return exprKeys[me]
}

func hasKeyPrefix(s string) (Key, bool) {
	idx, ok := hasPrefix(s, exprKeys[:])

	return Key(idx), ok
}

//==============================================================================

type EscapeInfo struct {
	s string
	c rune
}

var exprEscapes = []EscapeInfo{
	EscapeInfo{
		s: `\n`,
		c: '\n',
	},
	EscapeInfo{
		s: `\t`,
		c: '\t',
	},
}

func hasEscapePrefix(s string) (string, rune, bool) {
	for _, v := range exprEscapes {
		if strings.HasPrefix(s, v.s) {
			Log.Info("has escape prefix:%s", v.s)

			return v.s, v.c, true
		}
	}

	return Empty, 0, false
}
