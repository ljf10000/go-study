package libbignumber

import (
	. "asdf"
)

func BigAdd(sa, sb string) string {
	a := []byte(sa)
	b := []byte(sb)

	ByteReverse(a)
	ByteReverse(b)
}

func BigMulti(a, b string) string {

}
