package apiserver

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/hahaclassic/learning-rest-api.git/internal/app/ciphers"
	"github.com/hahaclassic/learning-rest-api.git/internal/app/encryption"
)

func PostRequestHandler(w http.ResponseWriter, r *http.Request,
	cipher encryption.EncryptionMethod) encryption.Values {

	body, _ := ioutil.ReadAll(r.Body)

	words := strings.Split(string(body)[8:], "+")

	var builder strings.Builder
	for _, word := range words {
		fmt.Fprintf(&builder, "%s ", word)
	}
	message := builder.String()
	message = message[:len(message)-1]

	values := encryption.Values{
		Message:  message,
		Language: "eng",
	}

	encryption.GetRandomKey(cipher, &values)
	encryptedMessage := encryption.Encrypt(cipher, values)

	values = encryption.Values{
		Message:   message,
		Encrypted: encryptedMessage,
	}

	return encryptedData
}

func (s *APIServer) HandleHome() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ts, err := template.ParseFiles("./ui/html/caesar.html")

		if err != nil {
			s.logger.Fatal(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		data := encryption.Values{}

		if r.Method == http.MethodPost {
			// var caesar encryption.EncryptionMethod = &ciphers.CaesarCipher{}
			data = PostRequestHandler(w, r, &ciphers.CaesarCipher{})
		}

		err = ts.Execute(w, data)

		if err != nil {
			s.logger.Fatal(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

	}
}

func (s *APIServer) HandleCaesar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Caesar Cipher")
	}
}
func (s *APIServer) HandleVigenere() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Vigenere Cipher")
	}
}

func (s *APIServer) HandleSimpleSC() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "simple Substitution Cipher")
	}
}

func (s *APIServer) HandleAffine() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Affine CIpher")
	}
}
