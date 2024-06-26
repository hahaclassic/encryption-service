package apiserver

import (
	"encoding/json"
	"errors"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/hahaclassic/learning-rest-api.git/internal/app/ciphers"
	"github.com/hahaclassic/learning-rest-api.git/internal/app/encryption"
)

func (s *APIServer) showPage(w http.ResponseWriter, r *http.Request, FilePath string) error {

	templ, err := template.ParseFiles(FilePath)

	if err != nil {
		return err
	}

	err = templ.Execute(w, nil)
	s.logger.Info("GET-request processed successfully")

	return err
}

func (s *APIServer) PostRequestHandler(w http.ResponseWriter, r *http.Request,
	cipher encryption.EncryptionMethod) error {

	body, _ := ioutil.ReadAll(r.Body)
	data := encryption.Values{}
	err := json.Unmarshal(body, &data)

	if err != nil {
		return err
	}

	s.logger.Println("Input Data: ", data)

	switch data.OperationType {
	case "Encrypt":
		encryption.Encrypt(cipher, &data)
	case "Decrypt":
		encryption.Decrypt(cipher, &data)
	case "GetRandomKey":
		encryption.GetRandomKey(cipher, &data)
	default:
		return errors.New("Wrong operation type")
	}

	JsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	s.logger.Println("Output Data: ", data)

	w.Header().Set("Content-Type", "application/json")
	w.Write(JsonData)

	return nil
}

func (s *APIServer) HandleHome() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome Home")
	}
}

func (s *APIServer) HandlerCaesar(w http.ResponseWriter, r *http.Request) {

	var err error

	if r.Method == http.MethodGet {
		err = s.showPage(w, r, "./ui/html/caesar.html")
	} else {
		err = s.PostRequestHandler(w, r, &ciphers.CaesarCipher{})
	}

	if err != nil {
		s.logger.Fatal(err)
		http.Error(w, "Internal Server Error", 500)
	}

}

func (s *APIServer) HandlerVigenere(w http.ResponseWriter, r *http.Request) {

	var err error

	if r.Method == http.MethodGet {
		err = s.showPage(w, r, "./ui/html/vigenere.html")
	} else {
		err = s.PostRequestHandler(w, r, &ciphers.VigenereCipher{})
	}

	if err != nil {
		s.logger.Fatal(err)
		http.Error(w, "Internal Server Error", 500)
	}

}

func (s *APIServer) HandlerAffine(w http.ResponseWriter, r *http.Request) {

	var err error

	if r.Method == http.MethodGet {
		err = s.showPage(w, r, "./ui/html/affine.html")
	} else {
		err = s.PostRequestHandler(w, r, &ciphers.AffineCipher{})
	}

	if err != nil {
		s.logger.Fatal(err)
		http.Error(w, "Internal Server Error", 500)
	}

}

func (s *APIServer) HandlerSimpleSubstitution(w http.ResponseWriter, r *http.Request) {

	var err error

	if r.Method == http.MethodGet {
		err = s.showPage(w, r, "./ui/html/simplesubstitution.html")
	} else {
		err = s.PostRequestHandler(w, r, &ciphers.SimpleSubstitutionCipher{})
	}

	if err != nil {
		s.logger.Fatal(err)
		http.Error(w, "Internal Server Error", 500)
	}

}
