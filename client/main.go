package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	client := getClient()
	resp, err := client.Get("https://the-server:8080")
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Client error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Status: %s  Body: %s\n", resp.Status, string(body))
}

func getClient() *http.Client {
	cp := x509.NewCertPool()
	data, _ := ioutil.ReadFile("../ca/minica.pem")
	cp.AppendCertsFromPEM(data)

	c, _ := tls.LoadX509KeyPair("cert.pem", "key.pem")

	config := &tls.Config{
		RootCAs:      cp,
		Certificates: []tls.Certificate{c},
	}

	client := &http.Client{
		Timeout: time.Minute * 3,
		Transport: &http.Transport{
			TLSClientConfig: config,
		},
	}
	return client
}
