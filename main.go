package main

import (
	"crypto/tls"
	_ "embed"
	"encoding/xml"
	"os"

	"github.com/amdonov/xmlsig"
)

//go:embed cert-key.pem
var key []byte

//go:embed cert.pem
var cert []byte

type Test1 struct {
	XMLName   xml.Name `xml:"urn:envelope Envelope"`
	ID        string   `xml:",attr"`
	Data      string   `xml:"urn:envelope Data"`
	Signature *xmlsig.Signature
}

func example() error {
	cert, err := tls.X509KeyPair(cert, key)
	if err != nil {
		return err
	}
	signer, err := xmlsig.NewSigner(cert)
	if err != nil {
		return err
	}
	doc := Test1{
		Data: "Hello, World!",
		ID:   "_1234",
	}
	sig, err := signer.CreateSignature(doc)
	if err != nil {
		return err
	}
	doc.Signature = sig
	encoder := xml.NewEncoder(os.Stdout)
	return encoder.Encode(doc)
}

func main() {
	err := example()
	if err != nil {
		panic(err)
	}
}
