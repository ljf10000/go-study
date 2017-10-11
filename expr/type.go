package libexpr

import (
	. "asdf"
	"strings"
)

const (
	TypeLogic  Type = 0
	TypeOp     Type = 1
	TypeString Type = 2
	TypeEnd    Type = 3
)

type Type int

const (
	ScopeAll  Scope = 0
	ScopeZone Scope = 1
	ScopePath Scope = 2
	ScopeEnd  Scope = 3
)

type Scope int

const (
	LogicAnd Logic = 0
	LogicOr  Logic = 1
	LogicNot Logic = 2
	LogicEnd Logic = 3
)

var logics = [LogicEnd]string{
	LogicAnd: "&&",
	LogicOr:  "||",
	LogicNot: "!",
}

type Logic int

func (me Logic) String() string {
	return logics[me]
}

func hasLogicPrefix(s string) (Logic, bool) {
	for i := Logic(0); i < LogicEnd; i++ {
		v := logics[i]

		if strings.HasPrefix(s, v) {
			Log.Info("has logic prefix:%s", v)

			return i, true
		}
	}

	return LogicEnd, false
}

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

var operators = [OpEnd]string{
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
	return operators[me]
}

func hasOpPrefix(s string) (Op, bool) {
	for i := Op(0); i < OpEnd; i++ {
		v := operators[i]

		if strings.HasPrefix(s, v) {
			Log.Info("has op prefix:%s", v)

			return i, true
		}
	}

	return OpEnd, false
}

type EscapeInfo struct {
	s string
	c rune
}

var escapes = []EscapeInfo{
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
	for _, v := range escapes {
		if strings.HasPrefix(s, v.s) {
			Log.Info("has escape prefix:%s", v.s)

			return v.s, v.c, true
		}
	}

	return Empty, 0, false
}
