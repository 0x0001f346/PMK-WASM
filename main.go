package main

import (
	"crypto/sha1"
	"encoding/hex"
	"syscall/js"
	"unicode/utf8"

	"golang.org/x/crypto/pbkdf2"
)

func main() {
	js.Global().Set("generateWpa2Pmk", GenerateWpa2PmkWrapper())
	<-make(chan bool)
}

func GenerateWpa2PmkWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "error"
		}
		return GenerateWpa2Pmk(args[0].String(), args[1].String())
	})

	return jsonFunc
}

func GenerateWpa2Pmk(passphrase string, ssid string) string {
	if !isValidPassphrase(passphrase) {
		return "error"
	}

	if !isValidSSID(ssid) {
		return "error"
	}

	return hex.EncodeToString(
		pbkdf2.Key(
			[]byte(passphrase),
			[]byte(ssid),
			4096,
			32,
			sha1.New,
		),
	)
}

func isValidPassphrase(passphrase string) bool {
	length_passphrase := utf8.RuneCountInString(passphrase)

	if length_passphrase < 8 || length_passphrase > 63 {
		return false
	}

	// ASCII characters from 32 to 126 are valid
	for i := 0; i < len(passphrase); i++ {
		if passphrase[i] < 32 || passphrase[i] > 126 {
			return false
		}
	}

	return true
}

func isValidSSID(ssid string) bool {
	length_ssid := utf8.RuneCountInString(ssid)

	if length_ssid < 1 || length_ssid > 32 {
		return false
	}

	// ASCII characters from 32 to 126 are valid
	for i := 0; i < len(ssid); i++ {
		if ssid[i] < 32 || ssid[i] > 126 {
			return false
		}
	}

	return true
}
