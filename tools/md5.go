package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(s string) (string, error) {
	h := md5.New()
	_, err := h.Write([]byte(s))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
