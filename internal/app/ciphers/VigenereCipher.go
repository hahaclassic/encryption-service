package ciphers

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"strings"
	"time"
)

type VigenereCipher struct {
	alphabet string
}

func (cipher *VigenereCipher) GetAlphabet(language string) {

	// Оставлена возможность для шифрования предложений на русском языке
	if language == "eng" {
		cipher.alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890 !?.,"
	}

}

func (cipher *VigenereCipher) GetRandomKey() string {
	len := len(cipher.alphabet)

	currentTime := time.Now().Unix()
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(currentTime))

	if err != nil {
		panic(err)
	}

	num := int(safeNum.Int64()) % len
	return fmt.Sprintf("%d", num)
}

func (cipher *VigenereCipher) EncryptMessage(message string, key string) string {
	alphabet_len, message_len := len(cipher.alphabet), len(message)
	keyIndex := 0

	translated := make([]rune, message_len, message_len)
	var translatedIndex int

	for i, symbol := range message {

		translatedIndex = strings.Index(cipher.alphabet, string(symbol))

		if translatedIndex == -1 {
			translated[i] = symbol
		} else {
			translatedIndex += strings.Index(cipher.alphabet, string(key[keyIndex]))
			translatedIndex %= alphabet_len

			translated[i] = rune(cipher.alphabet[translatedIndex])

			keyIndex = (keyIndex + 1) % len(key)
		}
	}

	return string(translated)
}

func (cipher *VigenereCipher) DecryptMessage(message string, key string) string {
	alphabet_len, message_len := len(cipher.alphabet), len(message)
	keyIndex := 0

	translated := make([]rune, message_len, message_len)
	var translatedIndex int

	for i, symbol := range message {

		translatedIndex = strings.Index(cipher.alphabet, string(symbol))

		if translatedIndex == -1 {
			translated[i] = symbol
		} else {
			translatedIndex -= strings.Index(cipher.alphabet, string(key[keyIndex]))

			if translatedIndex < 0 {
				translatedIndex += alphabet_len
			}

			translated[i] = rune(cipher.alphabet[translatedIndex])

			keyIndex = (keyIndex + 1) % len(key)
		}
	}

	return string(translated)
}
