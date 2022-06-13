package compare

import (
	"fmt"
)

func MakeKey(bd []byte) (string, error) {
	// get several data from bd
	// first and last should be easily taken
	// but it usually starts with `data:image/...`
	length := len(bd)
	var lastIdx string
	var usableKeys []byte

	// build keys
	lastIdx = string(bd[len(bd)-1])
	for i := 1; i < length; i *= 2 {
		usableKeys = append(usableKeys, bd[i])
	}

	return fmt.Sprintf("%s%s%d", lastIdx, string(usableKeys), length), nil
}
