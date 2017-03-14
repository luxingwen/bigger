package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"strings"
)

func EncodePasswd(src, str string) string {
	s := src + str + src
	s = strings.ToUpper(s)
	t := md5.New()
	io.WriteString(t, s)
	io.WriteString(t, src)
	return hex.EncodeToString(t.Sum(nil))
}
