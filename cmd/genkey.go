package cmd

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
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
	pubKeyBytes, err := x509.MarshalPKIXPublicKey(&pubKey)
	if err != nil {
		logrus.Fatal(err)
	}
	var privKeyPem = &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}
	var pubKeyPem = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	}
	privKeyOut, err := os.Create("harmony-key.pem")
	if err != nil {
		logrus.Fatal(err)
	}
	defer privKeyOut.Close()
	pubKeyOut, err := os.Create("harmony-key.pub")
	if err != nil {
		logrus.Fatal(err)
	}
	defer pubKeyOut.Close()
	if err := pem.Encode(privKeyOut, privKeyPem); err != nil {
		logrus.Fatal(err)
	}
	if err := pem.Encode(pubKeyOut, pubKeyPem); err != nil {
		logrus.Fatal(err)
	}
}
