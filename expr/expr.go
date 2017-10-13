package libexpr

type Expr struct {
	Left  *Expr
	Rigth *Expr
	Logic Logic

	Op    Op
	Scope Scope
	Key   string
	Value string
}

