package pow

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

// IsHashCorrect checks if the hash is correct for the given data and nonce and leadingZerosCount
func IsHashCorrect(data int64, nonce int64, leadingZerosCount int64) bool {
	stringToHash := strconv.FormatInt(data, 10) + strconv.FormatInt(nonce, 10)

	hash := sha256.New()
	hash.Write([]byte(stringToHash))
	hashed := hex.EncodeToString(hash.Sum(nil))

	fmt.Printf("\rCurrent hash being checked: %s", hashed) // Print the current hash in the same line
	os.Stdout.Sync()                                       // Flush the buffer to immediately display the output

	for i := int64(0); i < leadingZerosCount; i++ {
		if hashed[i] != '0' {
			return false
		}
	}
	return true
}
