package encryption

type Values struct {
	Message  string
	Key      string
	Language string
}

type EncryptionMethod interface {
	EncryptMessage(message string, key string) string
	DecryptMessage(message string, key string) string
	GetAlphabet(language string)
	GetRandomKey() string
}

func Encrypt(cipher EncryptionMethod, values Values) string {
	cipher.GetAlphabet(values.Language)
	encryptedMessage := cipher.EncryptMessage(values.Message, values.Key)

	return encryptedMessage
}

func Decrypt(cipher EncryptionMethod, values Values) string {
	cipher.GetAlphabet(values.Language)
	decryptedMessage := cipher.DecryptMessage(values.Message, values.Key)

	return decryptedMessage
}

func GetRandomKey(cipher EncryptionMethod, values *Values) {
	cipher.GetAlphabet(values.Language)
	values.Key = cipher.GetRandomKey()
}
