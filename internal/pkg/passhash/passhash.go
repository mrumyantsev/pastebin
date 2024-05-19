package passhash

import (
	"crypto/sha1"
	"fmt"
)

func PasswordHash(password, salt string) string {
	if len(password) == 0 {
		return ""
	}

	hash := sha1.New()

	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
