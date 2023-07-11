package ciphers

import (
	crypto "crypto/rand"
	"math/big"
	"strings"
	"time"
)

type SimpleSubstitutionCipher struct {
	alphabet string
}

func (cipher *SimpleSubstitutionCipher) GetAlphabet(language string) {

	// Оставлена возможность для шифрования предложений на русском языке
	if language == "eng" {
		cipher.alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890 !?.,"
	}

}

func (cipher *SimpleSubstitutionCipher) GetRandomKey() string {
	len := len(cipher.alphabet)
	key := []rune(cipher.alphabet)

	for i := len - 1; i > 0; i-- {
		currentTime := time.Now().Unix()
		safeNum, _ := crypto.Int(crypto.Reader, big.NewInt(currentTime))
		newIndex := int(safeNum.Int64()) % i

		key[i], key[newIndex] = key[newIndex], key[i]
	}

	return string(key)
}

func (cipher *SimpleSubstitutionCipher) EncryptMessage(message string, key string) string {
	message_len := len(message)
	translated := make([]rune, message_len, message_len)
	var translatedIndex int

	for i, symbol := range message {

		translatedIndex = strings.Index(cipher.alphabet, string(symbol))

		if translatedIndex == -1 {
			translated[i] = symbol
		} else {
			translated[i] = rune(key[translatedIndex])
		}
	}

	return string(translated)
}

func (cipher *SimpleSubstitutionCipher) DecryptMessage(message string, key string) string {
	message_len := len(message)
	translated := make([]rune, message_len, message_len)
	var translatedIndex int

	for i, symbol := range message {

		translatedIndex = strings.Index(key, string(symbol))

		if translatedIndex == -1 {
			translated[i] = symbol
		} else {
			translated[i] = rune(cipher.alphabet[translatedIndex])
		}
	}

	return string(translated)
}
