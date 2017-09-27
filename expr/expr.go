package libexpr

const (
	OpAnd     Operator = 0 // &&
	OpOr      Operator = 1 // ||
	OpNot     Operator = 2 // !
	OpInclude Operator = 3 //
	OpGe      Operator = 4 // >
	OpEqGe    Operator = 5 // >=
	OpLe      Operator = 6 // <
	OpEqLe    Operator = 7 // <=
	OpEq      Operator = 8 // ==
	OpNeq     Operator = 9 // !=
	OpEnd     Operator = 10
)

type Operator int

const (
	ScopeALl  Scope = 0
	ScopeZone Scope = 1
	ScopePath Scope = 2
	ScopeEnd  Scope = 3
)

type Scope int

type Expr struct {
	Left  *Expr
	Rigth *Expr

	Op    Operator
	Scope Scope
	Key   string
	Value string
}

type ExprLex struct {
	Line          string
	Token         string
	Tokens        []string
	last          int
	current       int
	inQuot        bool
	inQarenthesis bool
	state         State
	Root          *Expr
}

func (me *ExprLex) Scan(line string) error {
	me.Line = line

	return me.scan()
}

func (me *ExprLex) scan() error {
	for idx, c := range me.Line {
		err := me.char(idx)
		if nil != err {
			return err
		}
	}

	return nil
}

func (me *ExprLex) one(idx int) error {
	c := me.Line[idx]

	switch c {
	case '"':
		if me.inQuot {

		} else {
			me.inQuot = true
		}
	case '(':
	case ')':
	case '\'':
	case ' ', '\t':
	default:
	}
}
