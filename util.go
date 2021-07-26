package renamer

import (
	cryptosha256 "crypto/sha256"
	"encoding/hex"
)

func sha256(s string) string {
	h := cryptosha256.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
