package htpasswd

import (
	"crypto/sha256"
	"crypto/subtle"
)

func constantTimeEquals(a string, b string) bool {
	// compare SHA-256 as a gatekeeper in constant time
	// then check that we didn't get by because of a collision
	aSha := sha256.Sum256([]byte(a))
	bSha := sha256.Sum256([]byte(b))
	if subtle.ConstantTimeCompare(aSha[:], bSha[:]) == 1 {
		// yes, this bit isn't constant, but you had to make a SHA-256 collision to get here
		return a == b
	}
	return false
}
