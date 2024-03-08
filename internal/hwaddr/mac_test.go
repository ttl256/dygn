package hwaddr_test

import (
	"testing"

	"github.com/ttl256/dygn/internal/hwaddr"
)

func TestParseAddr(t *testing.T) {
	validAddrTests := []struct {
		input string
		want  string
	}{
		{"00:11:22:33:44:55", "00:11:22:33:44:55"},
		{"00-11-22-33-44-55", "00:11:22:33:44:55"},
		{"0011.2233.4455", "00:11:22:33:44:55"},
		{"001122334455", "00:11:22:33:44:55"},
		{"aa:BB-cc.DdeE:f:F", "aa:bb:cc:dd:ee:ff"},
		{"000000000000", "00:00:00:00:00:00"},
		{"FFFFFFFFFFFF", "ff:ff:ff:ff:ff:ff"},
	}

	invalidAddrTests := []string{
		"0:11:22:33:44:55",
		"00:11:22:33:44:55:AA",
		"00-11-22-33-44-5",
		"T011.2233.4455",
		"kind of mac addr",
	}

	t.Run("valid MAC", func(t *testing.T) {
		for _, tt := range validAddrTests {
			got, err := hwaddr.ParseAddr(tt.input)
			if err != nil {
				t.Error("did not expect an error:", err)
			}
			if got.String() != tt.want {
				t.Errorf("got %q, want %q", got.String(), tt.want)
			}
		}
	})

	t.Run("invalid MAC", func(t *testing.T) {
		for _, input := range invalidAddrTests {
			_, err := hwaddr.ParseAddr(input)
			if err == nil {
				t.Error("wanted an error but did not get one")
			}
		}
	})
}
