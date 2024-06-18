// Package hotp implements https://datatracker.ietf.org/doc/html/rfc4226
package hotp

import (
	"crypto/hmac"
	"crypto/sha1"
)

// Generate generates a HTOP password
func Generate(secret []byte, counter uint64) uint32 {
	data := []byte{
		byte(counter >> 56), // byte() returns the lowest 8 bits
		byte(counter >> 48), // we need big endian so first shift all
		byte(counter >> 40), // the way to right to get the highest
		byte(counter >> 32), // order byte
		byte(counter >> 24),
		byte(counter >> 16),
		byte(counter >> 8),
		byte(counter >> 0),
	}
	hmacer := hmac.New(sha1.New, secret)
	hmacer.Write(data)
	hmac := hmacer.Sum(nil) // 20 bytes

	return truncate(hmac)
}

func truncate(hash []byte) (code uint32) {
	// dynamic truncate
	offset := hash[19] & 0xF // val between 0 and 15
	// return last 31 bits
	binCode := (uint32(hash[offset])&0x7F)<<24 |
		uint32(hash[offset+1])<<16 |
		uint32(hash[offset+2])<<8 |
		uint32(hash[offset+3])
	// HOTP value generation
	// RFC4226: MUST support 6 digits (10^6=1000000); SHOULD 7 & 8 & more
	// We stick with the minimal implementation
	return uint32(binCode % (1000000))
}
