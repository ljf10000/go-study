package libexpr

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	for i := 1; i < 10; i++ {
		RegisterKeyword(fmt.Sprintf("zone%d", i), ScopeZone)
	}

	for i := 1; i < 10; i++ {
		RegisterKeyword(fmt.Sprintf("obj%d", i), ScopeObject)
	}

	lines := []string{
		`a1`,
		`a1 a2`,
		`a1 a2 a3`,
		`a1 && a2`,
		`a1 && a2 && a3`,
		`a1 && a2 || a3`,
		`a1 || a2 || a3`,

		`obj1>0 obj2>=0 obj3==0 obj4!=0 obj5<0 obj6<=0`,
		`obj1>0 && obj2>=0 && obj3==0 && obj4!=0 && obj5<0 && obj6<=0`,
		`obj1>0 || obj2>=0 || obj3==0 || obj4!=0 || obj5<0 || obj6<=0`,
		`!obj1>0`,
		`!(obj1>0)`,
		`a1 ||!obj1>0`,
		`a1 ||!(obj1>0)`,
		`a1 zone1=xxx obj1>=0`,
		`a1 'v1 "v2"'`,
		`a1 'v1 \"v2\"'`,
		`a1 'v1 \"v2\"' && obj1 > 0||obj2<0`,
		`a1 a2||(a3 obj1<0)`,
		`a1 a2||(a3 obj1<0)&&(obj2!=0||obj3==0)`,
		`a1 a2||(a3 obj1<0)&&(obj2!=0||obj3==0||(obj4>=0&&obj5<=0))`,
		`a1 a2||(a3 obj1<0)&&(obj2!=0||obj3==0||(obj4>=0 && obj5<=0 ) )||! (obj6!= 0)`,
		`a1 a2||a3 zone1="v1 \'v2\'"&&(obj1>=0&&(obj2>1||obj3<1)||!(obj4==0&&obj5==0))||(obj6<0&&obj7>0)`,
	}

	for _, line := range lines {
		expr := Scan(line)
		expr = expr
	}
}