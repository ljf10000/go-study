package libexpr

const (
	OpAnd     Operator = 0 // &&
	OpOr      Operator = 1 // ||
	OpNot     Operator = 2 // !
	OpInclude Operator = 3 // =
	OpGe      Operator = 4 // >
	OpEqGe    Operator = 5 // >=
	OpLe      Operator = 6 // <
	OpEqLe    Operator = 7 // <=
	OpEq      Operator = 8 // ==
	OpNeq     Operator = 9 // !=
	OpEnd     Operator = 10
)

var operators = [OpEnd]string{
	OpAnd:     "&&",
	OpOr:      "||",
	OpNot:     "!",
	OpInclude: "=",
	OpGe:      ">",
	OpEqGe:    ">=",
	OpLe:      "<",
	OpEqLe:    "<=",
	OpEq:      "==",
	OpNeq:     "!=",
}

type Operator int

func (me Operator) String() string {
	return operators[me]
}
