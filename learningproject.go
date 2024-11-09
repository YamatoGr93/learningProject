package main

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const base58Alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
const base85Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!#$%&()*+-;<=>?@^_`{|}~"

func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	message := "Hello, World!"

	encodedBase64 := encodeBase64(message)
	encodedASCII := encodeASCII(message)
	encodedHex := encodeHex(message)
	encodedURL := encodeURL(message)
	encodedROT13 := encodeROT13(message)
	encodedBinary := encodeBinary(message)
	encodedOctal := encodeOctal(message)
	encodedCaesar := encodeCaesar(message, 3)
	encodedAtbash := encodeAtbash(message)
	encodedBase32 := encodeBase32(message)
	encodedBase58 := encodeBase58(message)
	encodedBase85 := encodeBase85(message)

	response := fmt.Sprintf(
		"Original: %s\nEncoded (Base64): %s\nEncoded (ASCII): %s\nEncoded (Hex): %s\nEncoded (URL): %s\nEncoded (ROT13): %s\nEncoded (Binary): %s\nEncoded (Octal): %s\nEncoded (Caesar): %s\nEncoded (Atbash): %s\nEncoded (Base32): %s\nEncoded (Base58): %s\nEncoded (Base85): %s",
		message, encodedBase64, encodedASCII, encodedHex, encodedURL, encodedROT13, encodedBinary, encodedOctal, encodedCaesar, encodedAtbash, encodedBase32, encodedBase58, encodedBase85,
	)
	_, err := fmt.Fprint(w, response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func encodeBase64(message string) string {
	return base64.StdEncoding.EncodeToString([]byte(message))
}

func encodeASCII(message string) string {
	asciiEncoded := ""
	for _, char := range message {
		asciiEncoded += fmt.Sprintf("%d ", char)
	}
	return asciiEncoded
}

func encodeHex(message string) string {
	return fmt.Sprintf("%x", message)
}

func encodeURL(message string) string {
	return url.QueryEscape(message)
}

func encodeROT13(message string) string {
	rot13 := func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return 'A' + (r-'A'+13)%26
		} else if r >= 'a' && r <= 'z' {
			return 'a' + (r-'a'+13)%26
		}
		return r
	}
	return strings.Map(rot13, message)
}

func encodeBinary(message string) string {
	binaryEncoded := ""
	for _, char := range message {
		binaryEncoded += fmt.Sprintf("%08b ", char)
	}
	return binaryEncoded
}

func encodeOctal(message string) string {
	octalEncoded := ""
	for _, char := range message {
		octalEncoded += fmt.Sprintf("%o ", char)
	}
	return octalEncoded
}

func encodeCaesar(message string, shift int) string {
	caesar := func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return 'A' + (r-'A'+rune(shift))%26
		} else if r >= 'a' && r <= 'z' {
			return 'a' + (r-'a'+rune(shift))%26
		}
		return r
	}
	return strings.Map(caesar, message)
}

func encodeAtbash(message string) string {
	atbash := func(r rune) rune {
		if r >= 'A' && r <= 'Z' {
			return 'Z' - (r - 'A')
		} else if r >= 'a' && r <= 'z' {
			return 'z' - (r - 'a')
		}
		return r
	}
	return strings.Map(atbash, message)
}

func encodeBase32(message string) string {
	return base32.StdEncoding.EncodeToString([]byte(message))
}

func encodeBase58(message string) string {
	num := big.NewInt(0).SetBytes([]byte(message))
	base := big.NewInt(int64(len(base58Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	var encoded string
	for num.Cmp(zero) > 0 {
		num.DivMod(num, base, mod)
		encoded = string(base58Alphabet[mod.Int64()]) + encoded
	}

	return encoded
}

func encodeBase85(message string) string {
	num := big.NewInt(0).SetBytes([]byte(message))
	base := big.NewInt(int64(len(base85Alphabet)))
	zero := big.NewInt(0)
	mod := &big.Int{}

	var encoded string
	for num.Cmp(zero) > 0 {
		num.DivMod(num, base, mod)
		encoded = string(base85Alphabet[mod.Int64()]) + encoded
	}

	return encoded
}
