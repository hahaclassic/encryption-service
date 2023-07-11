package ciphers

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"
)

type CaesarCipher struct {
	alphabet string
}

func (cipher *CaesarCipher) GetAlphabet(language string) {

	// Оставлена возможность для шифрования предложений на русском языке
	if language == "eng" {
		cipher.alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890 !?.,"
	}

}

func (cipher *CaesarCipher) GetRandomKey() string {
	len := len(cipher.alphabet)

	currentTime := time.Now().Unix()
	safeNum, err := crypto.Int(crypto.Reader, big.NewInt(currentTime))

	if err != nil {
		panic(err)
	}

	num := int(safeNum.Int64()) % len
	return fmt.Sprintf("%d", num)
}

func (cipher *CaesarCipher) EncryptMessage(message string, key string) string {
	alphabet_len, message_len := len(cipher.alphabet), len(message)
	int_key, _ := strconv.Atoi(key)
	int_key = int_key % alphabet_len

	translated := make([]rune, message_len, message_len)
	var symbolIndex, translatedIndex int

	for i, symbol := range message {

		symbolIndex = strings.Index(cipher.alphabet, string(symbol))

		if symbolIndex == -1 {
			translated[i] = symbol
		} else {
			translatedIndex = (symbolIndex + int_key) % alphabet_len
			translated[i] = rune(cipher.alphabet[translatedIndex])
		}
	}

	return string(translated)
}

func (cipher *CaesarCipher) DecryptMessage(message string, key string) string {
	alphabet_len, message_len := len(cipher.alphabet), len(message)
	int_key, _ := strconv.Atoi(key)
	int_key = int_key % alphabet_len

	translated := make([]rune, message_len, message_len)
	var symbolIndex, translatedIndex int

	for i, symbol := range message {

		symbolIndex = strings.Index(cipher.alphabet, string(symbol))

		if symbolIndex == -1 {
			translated[i] = symbol
		} else {
			translatedIndex = symbolIndex - int_key

			if translatedIndex < 0 {
				translatedIndex += alphabet_len
			}

			translated[i] = rune(cipher.alphabet[translatedIndex])
		}
	}

	return string(translated)
}
