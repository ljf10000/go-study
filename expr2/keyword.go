package libexpr

func deftKeyword() *Keyword {
	return &Keyword{
		Key: "__all__",
		//List:  []string{Empty},
		Scope: ScopeAll,
	}
}

type Keyword struct {
	Key  string
	Type int
	//List  []string
	Scope Scope
}

func (me *Keyword) String() string {
	return me.Key
}
