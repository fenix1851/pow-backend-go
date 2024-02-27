package pow

import "fmt"

// findNonce searches for a nonce that makes the hash of the data and the nonce have leadingZerosCount leading zeros
func FindNonce(data int64, leadingZerosCount int64) int64 {
	var nonce int64
	for {
		if IsHashCorrect(data, nonce, leadingZerosCount) {
			fmt.Println() // Print a newline to move to the next line
			return nonce
		}
		nonce++
	}
}
