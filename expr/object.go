package libexpr

import (
	. "asdf"
	"strings"
)

type Object struct {
	raw   string
	objs  []string
	scope Scope
}

func (me *Object) String() string {
	return me.raw
}

func newObject(s string) *Object {
	obj := &Object{
		raw: s,
	}

	if Empty == s {
		obj.scope = ScopeAll
	} else {
		obj.objs = strings.Split(s, ":")

		// todo
		match := false

		if match {
			obj.scope = ScopeZone
		} else {
			obj.scope = ScopeObject
		}
	}

	return obj
}
