// bgc-hmac project main.go
package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

var secret = "1234567890ABCDEF1234567890ABCDEF"

func main() {
	mess := "00000000"
	result, _ := getHmacString(mess)
	fmt.Println(result)
}

func getHmacString(str string) (string, error) {
	return getHmacByte([]byte(str))
}

func getHmacByte(b []byte) (string, error) {
	encKey, err := hex.DecodeString(secret)
	if err != nil {
		return "", err
	}
	mac := hmac.New(sha256.New, encKey)
	in := normalizeData(b)
	mac.Write(in)
	return strings.ToUpper(hex.EncodeToString(mac.Sum(nil)[:16])), nil
}

func normalizeData(indata []byte) []byte {
	var buffer bytes.Buffer
	for _, b := range indata {
		// Line endings should be removed
		if b != 13 && b != 10 {
			// Include only the accepted values
			if b >= 32 && b <= 126 {
				buffer.WriteByte(b)
			} else {
				// Check if we have any of our special cases ÅÄÖ etc ;)
				switch b {
				case 201:
					buffer.WriteByte(64)
				case 196:
					buffer.WriteByte(91)
				case 214:
					buffer.WriteByte(92)
				case 197:
					buffer.WriteByte(93)
				case 220:
					buffer.WriteByte(94)
				case 233:
					buffer.WriteByte(96)
				case 228:
					buffer.WriteByte(123)
				case 246:
					buffer.WriteByte(124)
				case 229:
					buffer.WriteByte(125)
				case 252:
					buffer.WriteByte(126)
				default:
					// When the values are not accepted write byte 195 instead
					buffer.WriteByte(195)
				}
			}
		}
	}
	return buffer.Bytes()
}
