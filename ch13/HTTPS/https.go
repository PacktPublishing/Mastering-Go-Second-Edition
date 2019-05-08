package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello there securely!\n"))
}

func main() {
	caCert, err := ioutil.ReadFile("client.crt")
	if err != nil {
		fmt.Println(err)
		return
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	cfg := &tls.Config{
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
	}

	srv := &http.Server{
		Addr:      ":8443",
		Handler:   &handler{},
		TLSConfig: cfg,
	}

	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}
