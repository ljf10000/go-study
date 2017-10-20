package libexpr

import (
	. "asdf"
	"strings"
)

//========================================================

type escapeInfo struct {
	s string
	c rune
}

var exprEscapes = []escapeInfo{
	escapeInfo{
		s: `\n`,
		c: '\n',
	},
	escapeInfo{
		s: `\t`,
		c: '\t',
	},
}

func hasEscapePrefix(s string) (string, rune, bool) {
	for _, v := range exprEscapes {
		if strings.HasPrefix(s, v.s) {
			Log.Info("has escape prefix:%s", v.s)

			return v.s, v.c, true
		}
	}

	return Empty, 0, false
}

//========================================================

//========================================================

//========================================================

//========================================================
