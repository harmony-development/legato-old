package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
	"os"
)

type KeyManager interface {
	GetPublicKey() []byte
}

type Ed25519KeyManager struct {
	pubKey  ed25519.PublicKey
	privKey ed25519.PrivateKey
}

func NewEd25519KeyManagerFromFile(privKeyPath string, pubKeyPath string) (KeyManager, error) {
	var privKeyFile, pubKeyFile *os.File
	var err error
	if privKeyFile, err = os.Open(privKeyPath); err != nil {
		return nil, err
	}
	if pubKeyFile, err = os.Open(pubKeyPath); err != nil {
		return nil, err
	}
	return NewEd25519KeyManager(privKeyFile, pubKeyFile)
}

func NewEd25519KeyManager(privKeyReader io.Reader, pubKeyReader io.Reader) (KeyManager, error) {
	privKey, err := io.ReadAll(privKeyReader)
	if err != nil {
		return nil, err
	}
	pubKey, err := io.ReadAll(pubKeyReader)
	if err != nil {
		return nil, err
	}
	return &Ed25519KeyManager{
		privKey: ed25519.PrivateKey(privKey),
		pubKey:  ed25519.PublicKey(pubKey),
	}, nil
}

func (manager *Ed25519KeyManager) GetPublicKey() []byte {
	return manager.pubKey
}

func WriteEd25519Keys(privKeyPath string, pubKeyPath string) error {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}
	if err := os.WriteFile(privKeyPath, priv, 0644); err != nil {
		return err
	}
	if err := os.WriteFile(pubKeyPath, pub, 0644); err != nil {
		return err
	}
	return nil
}
