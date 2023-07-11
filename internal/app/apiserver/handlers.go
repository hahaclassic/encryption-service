package apiserver

import (
	"io"
	"net/http"
)

func (s *APIServer) HandleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Welcome")
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
