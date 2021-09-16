// SPDX-FileCopyrightText: 2021 Danil Korennykh <bluskript@gmail.com>
//
// SPDX-License-Identifier: AGPL-3.0-or-later

package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
	"os"

	"github.com/harmony-development/legato/errwrap"
)

type Manager interface {
	GetPublicKey() []byte
}

type Ed25519KeyManager struct {
	pubKey  ed25519.PublicKey
	privKey ed25519.PrivateKey
}

func NewEd25519KeyManagerFromFile(privKeyPath string, pubKeyPath string) (Manager, error) {
	var privKeyFile, pubKeyFile *os.File

	var err error

	if privKeyFile, err = os.Open(privKeyPath); err != nil {
		return nil, errwrap.Wrap(err, "failed to open private key file")
	}

	if pubKeyFile, err = os.Open(pubKeyPath); err != nil {
		return nil, errwrap.Wrap(err, "failed to open public key file")
	}

	return NewEd25519KeyManager(privKeyFile, pubKeyFile)
}

func NewEd25519KeyManager(privKeyReader io.Reader, pubKeyReader io.Reader) (Manager, error) {
	privKey, err := io.ReadAll(privKeyReader)
	if err != nil {
		return nil, errwrap.Wrap(err, "failed to read private key")
	}

	pubKey, err := io.ReadAll(pubKeyReader)
	if err != nil {
		return nil, errwrap.Wrap(err, "failed to read public key")
	}

	return &Ed25519KeyManager{
		privKey: ed25519.PrivateKey(privKey),
		pubKey:  ed25519.PublicKey(pubKey),
	}, nil
}

func (manager *Ed25519KeyManager) GetPublicKey() []byte {
	return manager.pubKey
}

func WriteEd25519KeysToFile(privKeyPath string, pubKeyPath string) error {
	privFile, err := os.Create(privKeyPath)
	if err != nil {
		return errwrap.Wrap(err, "failed to create private key file")
	}
	defer privFile.Close()

	pubFile, err := os.Create(pubKeyPath)
	if err != nil {
		return errwrap.Wrap(err, "failed to create public key file")
	}
	defer pubFile.Close()

	return WriteEd25519Keys(privFile, pubFile)
}

func WriteEd25519Keys(privKeyWriter io.Writer, pubKeyWriter io.Writer) error {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return errwrap.Wrap(err, "failed to generate ed25519 key paid")
	}

	if _, err := privKeyWriter.Write(priv); err != nil {
		return errwrap.Wrap(err, "failed to write private key")
	}

	if _, err := pubKeyWriter.Write(pub); err != nil {
		return errwrap.Wrap(err, "failed to write public key")
	}

	return nil
}
