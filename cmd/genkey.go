package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"os"

	"github.com/sirupsen/logrus"
)

func GenKeys() {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		logrus.Fatal(err)
	}
	pubKey := key.PublicKey
	pubKeyBytes, err := asn1.Marshal(pubKey)
	if err != nil {
		logrus.Fatal(err)
	}
	var keyPem = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	var pubKeyPem = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	privKeyOut, err := os.Create("harmony-key.pem")
	if err != nil {
		logrus.Fatal(err)
	}
	pubKeyOut, err := os.Create("harmony-key.pub")
	if err != nil {
		logrus.Fatal(err)
	}
	if err := pem.Encode(privKeyOut, keyPem); err != nil {
		logrus.Fatal(err)
	}
	if err := pem.Encode(pubKeyOut, pubKeyPem); err != nil {
		logrus.Fatal(err)
	}
}
