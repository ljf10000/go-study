package libexpr

import (
	. "asdf"
)

//========================================================
var exprSkips = []rune{
	CHAR_SPACE,
	CHAR_TAB,
	CHAR_CRLF,
}

func isSkip(c rune) bool {
	for _, v := range exprSkips {
		if v == c {
			return true
		}
	}

	return false
}

var exprEscapes = []rune{
	't',
	'n',
	CHAR_SLASH,
	CHAR_QUOT_DOUBLE,
	CHAR_QUOT_SINGLE,
}

func isEscape(c rune) bool {
	for _, v := range exprEscapes {
		if c == v {
			return true
		}
	}

	return false
}

func hasEscapePrefix(s string) (string, rune, bool) {
	if len(s) >= 2 && s[0] == CHAR_SLASH {
		c := rune(s[1])
		if isEscape(c) {
			prefix := string(CHAR_SLASH) + string(c)
			Log.Info("has escape prefix:%s", prefix)

			return prefix, c, true
		} else {
			Panic("invalid escape \\%s", string(c))
		}
	}

	return Empty, 0, false
}

//========================================================
