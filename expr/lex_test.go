package libexpr

import (
	"testing"
)

func Test1(t *testing.T) {
	lines := []string{
		`a b c`,
		`a 'b "c"'`,
		`a 'b \"c\"'`,
		`a1 a2||a3 a="b \'c\'"&&(e>=0||f<10&&(g>1||h<1)||(i==0&&k==0))||(m<0&&n>0)`,
	}

	for _, line := range lines {
		LineScan(line)
	}
}
