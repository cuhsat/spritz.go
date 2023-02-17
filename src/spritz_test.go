// spritz_test.go
package spritz

import (
	"bytes"
	"testing"
)

func TestIntegration(t *testing.T) {
	t.Run("tests the integration of the encryption functionality", func(t *testing.T) {
		vectors := []string{"ABC", "spam", "arcfour"}

		for _, vector := range vectors {
			msg := []byte(vector)
			key := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xFF}

			Encrypt(msg, key)
			Decrypt(msg, key)

			if !bytes.Equal(msg, []byte(vector)) {
				t.Fatal("Message does not match expected result")
			}
		}
	})
}

func TestAlgorithm(t *testing.T) {
	t.Run("tests the algorithm with the official vectors from APPENDIX E", func(t *testing.T) {
		vectors := [][]byte{
			[]byte("ABC"),
			[]byte("spam"),
			[]byte("arcfour")}

		results := [][]byte{
			{0x02, 0x8F, 0xA2, 0xB4, 0x8B, 0x93, 0x4A, 0x18},
			{0xAC, 0xBB, 0xA0, 0x81, 0x3F, 0x30, 0x0D, 0x3A},
			{0xFF, 0x8C, 0xF2, 0x68, 0x09, 0x4C, 0x87, 0xB9}}

		for i, vector := range vectors {
			digest := make([]byte, 32)

			Hash(vector, digest)

			if !bytes.Equal(digest[:8], results[i]) {
				t.Fatal("Digest does not match expected result")
			}
		}
	})
}
