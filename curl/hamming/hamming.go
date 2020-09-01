// Package curl implements the Curl hashing function.
package hamming

import (
	. "github.com/iotaledger/iota.go/consts"
	"github.com/iotaledger/iota.go/curl"
	. "github.com/iotaledger/iota.go/trinary"
)

// Hamming calculates a nonce such that when chunk with the returned nonce is absorbed, the resulting hash will be normalized.
func Hamming(c *curl.Curl, chunk Trits, offset, end, security int) (Trits, error) {
	if len(chunk) > HashTrinarySize {
		return nil, ErrInvalidTritsLength
	}

	for {
		// compute the hash without modifying the state of c
		h := c.Clone()
		if err := h.Absorb(chunk); err != nil {
			return nil, err
		}
		hash := h.MustSqueeze(HashTrinarySize)
		if check(hash, security) {
			break
		}

		// increment the nonce
		for i := offset; i < end; i++ {
			if chunk[i] < MaxTritValue {
				chunk[i]++
				break
			}
			chunk[i] = MinTritValue
		}
	}
	return chunk[offset:end], nil
}

func check(hash Trits, security int) bool {
	sum := 0
	for j := 0; j < security; j++ {
		for k := j * HashTrinarySize / 3; k < (j+1)*HashTrinarySize/3; k++ {
			sum += int(hash[k])
		}
		if sum == 0 && j < security-1 {
			return false
		}
	}
	return sum == 0
}
