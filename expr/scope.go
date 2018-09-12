package libexpr

import (
	. "asdf"
)

const (
	ScopeAll    Scope = 0
	ScopeZone   Scope = 1
	ScopeObject Scope = 2
	ScopeEnd    Scope = 3
)

var exprScopes = [ScopeEnd]string{
	ScopeAll:    "all",
	ScopeZone:   "zone",
	ScopeObject: "object",
}

type Scope int

func (me Scope) IsGood() bool {
	return me >= 0 && me < ScopeEnd
}

func (me Scope) String() string {
	if me.IsGood() {
		return exprScopes[me]
	} else {
		return Unknow
	}
}
