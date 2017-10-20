package libexpr

import (
	. "asdf"
)

func deftKeyword() *Keyword {
	return &Keyword{
		keyword: "__all__",
		list:    []string{Empty},
		scope:   ScopeAll,
	}
}

type Keyword struct {
	keyword string
	list    []string
	scope   Scope
}

func (me *Keyword) String() string {
	return me.keyword
}
