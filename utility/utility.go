package utility

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func GetUniHashString(content string) string {
	h := md5.New()
	io.WriteString(h, content)

	hString := hex.EncodeToString(h.Sum(nil))
	return hString
}
