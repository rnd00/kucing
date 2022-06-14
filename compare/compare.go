package compare

import (
	"fmt"
)

func MakeKey(bd []byte) string {
	// get several data from bd
	// first and last should be easily taken
	// but it usually starts with `data:image/...`
	length := len(bd)
	var usableKeys []byte

	// usableKeys will be all different in length
	// takes 1-2-4-8-16-32-64-and so on
	for i := 1; i < length; i *= 2 {
		usableKeys = append(usableKeys, bd[i])
	}

	return fmt.Sprintf("%d-%s", length, string(usableKeys))
}
