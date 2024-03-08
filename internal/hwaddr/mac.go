package hwaddr

import (
	"encoding/hex"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	validChars = func() map[rune]struct{} { //nolint:gochecknoglobals // reasonable
		chars := map[rune]struct{}{}
		for _, r := range "0123456789abcdef" {
			chars[r] = struct{}{}
		}
		return chars
	}()
)

func filterHex(r rune) rune {
	r = unicode.ToLower(r)
	if _, ok := validChars[r]; !ok {
		return -1
	}
	return r
}

func ParseAddr(s string) (MAC, error) {
	const numHexDigits = 12

	normalizedMac := strings.Map(filterHex, s)

	if utf8.RuneCountInString(normalizedMac) != numHexDigits {
		return [6]byte{}, fmt.Errorf("invalid MAC address: %q", s)
	}

	d, err := hex.DecodeString(normalizedMac)
	if err != nil {
		return [6]byte{}, err
	}
	return [6]byte(d), nil
}

type MAC [6]byte

func (m MAC) String() string {
	return fmt.Sprintf(
		"%02x:%02x:%02x:%02x:%02x:%02x",
		m[0], m[1], m[2], m[3], m[4], m[5],
	)
}
