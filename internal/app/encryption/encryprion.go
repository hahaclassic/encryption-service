package encryption

type Values struct {
	OperationType    string `json:"operation_type"`
	OriginalMessage  string `json:"original"`
	ConvertedMessage string `json:"converted"`
	Key              string `json:"key"`
	Language         string `json:"language"`
}

type EncryptionMethod interface {
	EncryptMessage(message string, key string) string
	DecryptMessage(message string, key string) string
	GetAlphabet(language string)
	GetRandomKey() string
}

func Encrypt(cipher EncryptionMethod, data *Values) {
	cipher.GetAlphabet(data.Language)

	data.ConvertedMessage = cipher.EncryptMessage(data.OriginalMessage, data.Key)
}

func Decrypt(cipher EncryptionMethod, data *Values) {
	cipher.GetAlphabet(data.Language)

	data.ConvertedMessage = cipher.DecryptMessage(data.OriginalMessage, data.Key)
}

func GetRandomKey(cipher EncryptionMethod, data *Values) {
	cipher.GetAlphabet(data.Language)

	data.Key = cipher.GetRandomKey()
}
