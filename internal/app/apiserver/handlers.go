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

func (s *APIServer) HandleHome() http.HandlerFunc {

	type Data struct {
		Message   string
		Encrypted string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ts, err := template.ParseFiles("./ui/html/caesar.html")

		if err != nil {
			s.logger.Fatal(err)
			http.Error(w, "Internal Server Error", 500)
			return
		}

		if r.Method == http.MethodGet {

			err = ts.Execute(w, nil)

			if err != nil {
				s.logger.Println(err.Error())
				http.Error(w, "Internal Server Error", 500)
			}

		} else {

			body, err := ioutil.ReadAll(r.Body)
			if err == nil {
				s.logger.Println(string(body))
			}

			words := strings.Split(string(body)[8:], "+")

			var builder strings.Builder
			for _, word := range words {
				fmt.Fprintf(&builder, "%s ", word)
			}
			message := builder.String()
			message = message[:len(message)-1]

			values := encryption.InputValues{
				Message:  message,
				Language: "eng",
			}

			var caesar encryption.EncryptionMethod = &ciphers.CaesarCipher{}
			encryption.GetRandomKey(caesar, &values)
			encryptedMessage := encryption.Encrypt(caesar, values)

			s.logger.Println(message)
			s.logger.Println(encryptedMessage)

			data := Data{
				Message:   message,
				Encrypted: encryptedMessage,
			}

			err = ts.Execute(w, data)

			if err != nil {
				s.logger.Println(err.Error())
				http.Error(w, "Internal Server Error", 500)
			}

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
