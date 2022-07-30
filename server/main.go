package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	server := getServer()
	http.HandleFunc("/", myHandler)

	err := server.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}

}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello mTLS World!"))
}

func getServer() *http.Server {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../ca/minica.pem")
	cp.AppendCertsFromPEM(data)

	c, _ := tls.LoadX509KeyPair("cert.pem", "key.pem")

	tls := &tls.Config{
		Certificates:             []tls.Certificate{c},
		ClientCAs:                cp,
		ClientAuth:               tls.RequireAndVerifyClientCert,
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
	}

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tls,
	}
	return server
}
