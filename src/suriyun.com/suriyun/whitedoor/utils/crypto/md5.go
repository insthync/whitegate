package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMD5Hash ... Sum string to md5 string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
