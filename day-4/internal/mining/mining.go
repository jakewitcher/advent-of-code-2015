package mining

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func FindHashWithPrefix(prefix string, secret string) int {
	number := 1
	for {
		h := md5.New()
		value := fmt.Sprintf("%s%d", secret, number)
		h.Write([]byte(value))

		hexHash := hex.EncodeToString(h.Sum(nil))
		if strings.HasPrefix(hexHash, prefix) {
			return number
		}

		number++
	}
}
