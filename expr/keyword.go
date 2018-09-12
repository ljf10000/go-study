package libexpr

import (
	. "asdf"
)

func deftKeyword() *Keyword {
	return &Keyword{
		Key:   "__all__",
		Scope: ScopeAll,
		//List:  []string{Empty},
	}
}

type Keyword struct {
	Key   string
	Name  string
	Type  int
	Scope Scope
	//List  []string
}

func (me *Keyword) String() string {
	return me.Key
}

func (me *Keyword) checkOp(op Op) {
	switch me.Scope {
	case ScopeAll, ScopeZone:
		if op != OpInclude {
			Panic("zone key(%s) must use include op", me)
		}
	case ScopeObject:
		if op == OpInclude {
			Panic("object key(%s) cann't use include op", me)
		}
	}
}
