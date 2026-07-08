package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5WithFile(file io.Reader) string {
	m := md5.New()
	_, err := io.Copy(m, file)
	if err != nil {
		return ""
	}

	sum := m.Sum(nil)
	return hex.EncodeToString(sum)
}
