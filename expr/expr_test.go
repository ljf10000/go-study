package libexpr

import (
	. "asdf"
	"testing"
)

func Test1(t *testing.T) {
	lines := []string{
		`a b c`,
		`a 'b "c"'`,
		`a 'b \"c\"'`,
		`a1 a2 a3 a="b \'c\'"&&(e>=0||f<10)`,
	}

	for _, line := range lines {
		lex := &Lex{
			line: line,
		}

		lex.scan()
		Log.Info(lex.dump())
	}
}
