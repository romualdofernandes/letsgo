package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"
)

//AutocertConfig ...
type AutocertConfig struct {
	Debug       bool
	HTTPPort    int
	HTTPSPort   int
	ProxyURL    string
	RenewBefore int
}

func main() {

	m := &autocert.Manager{
		Prompt:      autocert.AcceptTOS,
		Cache:       autocert.DirCache("secret-dir"),
		HostPolicy:  autocert.HostWhitelist("golang.romualdo.com.br"),
		RenewBefore: 30,
		Client:      nil,
		Email:       "",
		ForceRSA:    false,
	}

	go http.ListenAndServe(":http", m.HTTPHandler(nil))

	s := &http.Server{
		Addr:      ":https",
		TLSConfig: &tls.Config{GetCertificate: m.GetCertificate},
	}
	log.Fatal(s.ListenAndServeTLS("", ""))

}
