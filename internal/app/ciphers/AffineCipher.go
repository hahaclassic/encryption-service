package ciphers

import (
	crypto "crypto/rand"
	"fmt"

	"math/big"
	"strconv"
	"strings"
	"time"

	mathfunc "github.com/hahaclassic/learning-rest-api.git/internal/app/pkgs/MathFunc"
)

type AffineCipher struct {
	alphabet string
}

func (cipher *AffineCipher) GetAlphabet(language string) {

	// Оставлена возможность для шифрования предложений на русском языке
	if language == "eng" {
		cipher.alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890 !?.,"
	}

}

func (cipher *AffineCipher) GetRandomKey() string {
	len := len(cipher.alphabet)
	var key int

	for {
		currentTime := time.Now().Unix()
		safeNum1, _ := crypto.Int(crypto.Reader, big.NewInt(currentTime))
		safeNum2, _ := crypto.Int(crypto.Reader, big.NewInt(currentTime))

		keyA := int(safeNum1.Int64()) % len
		keyB := int(safeNum2.Int64()) % len

		if mathfunc.GreatestCommonDiv(keyA, len) == 1 {
			key = keyA*len + keyB
			break
		}
	}

	return fmt.Sprintf("%d", key)
}

func (cipher *AffineCipher) EncryptMessage(message string, key string) string {
	alphabet_len, message_len := len(cipher.alphabet), len(message)
	int_key, _ := strconv.Atoi(key)
	keyA := int_key / alphabet_len
	keyB := int_key % alphabet_len

	translated := make([]rune, message_len, message_len)
	var symbolIndex, translatedIndex int

	for i, symbol := range message {

		symbolIndex = strings.Index(cipher.alphabet, string(symbol))

		if symbolIndex == -1 {
			translated[i] = symbol
		} else {
			translatedIndex = (symbolIndex*keyA + keyB) % alphabet_len
			translated[i] = rune(cipher.alphabet[translatedIndex])
		}
	}

	return string(translated)
}

func (cipher *AffineCipher) DecryptMessage(message string, key string) string {
	alphabet_len, message_len := len(cipher.alphabet), len(message)
	int_key, _ := strconv.Atoi(key)
	keyA := int_key / alphabet_len
	keyB := int_key % alphabet_len
	modInverse := mathfunc.FindModInverse(keyA, alphabet_len)

	translated := make([]rune, message_len, message_len)
	var symbolIndex, translatedIndex int

	for i, symbol := range message {

		symbolIndex = strings.Index(cipher.alphabet, string(symbol))

		if symbolIndex == -1 {
			translated[i] = symbol
		} else {
			translatedIndex = (symbolIndex - keyB) * modInverse % alphabet_len

			if translatedIndex < 0 {
				translatedIndex += alphabet_len
			}

			translated[i] = rune(cipher.alphabet[translatedIndex])
		}
	}

	return string(translated)
}
