package libexpr

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
