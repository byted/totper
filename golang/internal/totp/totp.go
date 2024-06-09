// Package totp implements https://datatracker.ietf.org/doc/html/rfc6238
package totp

import (
	"encoding/base32"
	"fmt"
	"time"

	"github.com/byted/totper/internal/hotp"
)

type TOTPer struct {
	secret    []byte
	timeSteps int64
}

// NewTOTPer initiates a TOTP generator based on a given Base32 encoded secret
func NewTOTPer(secret string) (*TOTPer, error) {
	decodedSecret, err := decodeBase32(secret)
	if err != nil {
		return nil, fmt.Errorf("unable to decode secret string: %v", err)
	}
	return &TOTPer{
		secret:    decodedSecret,
		timeSteps: int64(30),
	}, nil
}

// GetOTP gets the current OTP
//
// implements https://datatracker.ietf.org/doc/html/rfc6238
func (t *TOTPer) GetOTP() uint32 {
	timeCounter := time.Now().Unix() / t.timeSteps
	return hotp.Generate(t.secret, uint64(timeCounter))
}

func decodeBase32(src string) ([]byte, error) {
	dst := make([]byte, base32.StdEncoding.DecodedLen(len(src)))
	if _, err := base32.StdEncoding.Decode(dst, []byte(src)); err != nil {
		return nil, err
	}
	return dst, nil
}
