package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve files from the current directory.
	fs := http.FileServer(http.Dir("../"))
	http.Handle("/", fs)

	// Define the certificate and key file names.
	certFile := "cert.pem"
	keyFile := "key.pem"

	log.Println("Server is listening on https://localhost:8443")
	err := http.ListenAndServeTLS(":8443", certFile, keyFile, nil)
	if err != nil {
		log.Fatal("ListenAndServeTLS: ", err)
	}
}
