package libexpr

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

type Lex struct {
	Root        *Expr
	Line        string
	Token       string
	Tokens      []string
	last        rune
	current     rune
	quot        rune
	inQuot      bool
	parentheses int
}

func (me *Lex) Scan(line string) error {
	me.Line = line

	return me.scan()
}

func (me *Lex) scan() error {
	for idx, c := range me.Line {
		err := me.char(idx, c)
		if nil != err {
			return err
		}
	}

	return nil
}

func (me *Lex) char(idx int, c rune) error {
	me.current = c

	defer func() {
		me.last = c
	}()

	if me.inQuot {
		return me.quotHandle(idx, c)
	}

	switch c {
	case '"', '\'':
		me.inQuot = true
		me.quot = c
		me.Token = ""
	case '(':
		me.parentheses++
		me.saveToken(string(c))
	case ')':
		me.parentheses--
		me.saveToken(string(c))
	case ' ', '\t':
		me.skip(c)
	default:
		me.save(c)
	}

	return nil
}

func (me *Lex) quotHandle(idx int, c rune) error {
	if c == me.quot {
		// close quot
		me.inQuot = false
		me.saveToken(me.Token)
	} else {
		me.save(c)
	}

	return nil
}

func (me *Lex) saveToken(token string) {
	me.Tokens = append(me.Tokens, token)
	me.Token = ""
}

func (me *Lex) skip(c rune) {
	// do nothing
}

func (me *Lex) save(c rune) {
	me.Token += string(c)
}
