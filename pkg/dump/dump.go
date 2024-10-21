// Package dump contains functions to dump data in various formats.
package dump

import "fmt"

// ByteSliceAsHexDump prints a slice of bytes in a hex dump format.
func ByteSliceAsHexDump(b []byte) {
	n := (len(b) + 15) &^ 15
	for i := 0; i < n; i++ {
		if i%16 == 0 {
			fmt.Printf("%04x  ", i)
		}
		if i < len(b) {
			fmt.Printf(" %02x", b[i])
		}
		if i%16 == 15 {
			fmt.Print("\n")
		}
	}
}
