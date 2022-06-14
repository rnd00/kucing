package helpers

import "encoding/base64"

func ToBase64String(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func ToBase64Bytes(b []byte) []byte {
	return []byte(ToBase64String(b))
}
