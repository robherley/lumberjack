package es

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	elasticsearch6 "github.com/elastic/go-elasticsearch/v6"
)

func grabTLSCerts() (*tls.Config, error) {
	tlsCfg := tls.Config{}
	tlsCfg.RootCAs = x509.NewCertPool()

	files, err := ioutil.ReadDir("certs/")
	if err != nil {
		return nil, err
	}

	// read over our entire cert directory and add the certs
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".pem") {
			ca, err := ioutil.ReadFile("certs/" + file.Name())
			if err != nil {
				return nil, errors.New("unable to configure db tls: " + err.Error())
			}
			tlsCfg.RootCAs.AppendCertsFromPEM(ca)
		}
	}

	return &tlsCfg, nil
}

// connect will grab TLS certs and connect to our ES instance
func connect() (*elasticsearch6.Client, error) {
	tls, err := grabTLSCerts()
	if err != nil {
		return nil, err
	}

	fullURL := fmt.Sprintf("https://%s:%s", os.Getenv("ELASTIC_HOST"), os.Getenv("ELASTIC_PORT"))

	es6Config := elasticsearch6.Config{
		Addresses: []string{fullURL},
		Username:  os.Getenv("ELASTIC_USER"),
		Password:  os.Getenv("ELASTIC_PASS"),
		Transport: &http.Transport{
			TLSClientConfig: tls,
		},
	}

	return elasticsearch6.NewClient(es6Config)
}
