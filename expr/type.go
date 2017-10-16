package libexpr

const (
	TypeUnknow   Type = 0
	TypeObject   Type = 1
	TypeValue    Type = 2
	TypeString   Type = 3
	TypeLogic    Type = 4
	TypeOperator Type = 5
	TypeKeyWord  Type = 6
	TypeTerm     Type = 7
	TypeEnd      Type = 8
)

var exprTypes = [TypeEnd]string{
	TypeUnknow:   "unknow",
	TypeObject:   "object",
	TypeValue:    "value",
	TypeString:   "string",
	TypeLogic:    "logic",
	TypeOperator: "operator",
	TypeKeyWord:  "keyword",
	TypeTerm:     "term",
}

type Type int

func (me Type) String() string {
	return exprTypes[me]
}
