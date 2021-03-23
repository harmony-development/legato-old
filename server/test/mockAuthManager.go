package test

import "crypto/rsa"

type MockAuthManager struct {
}

func (m MockAuthManager) MakeAuthToken(userID uint64, target, username, avatar string) (string, error) {
	panic("unimplemented")
}
func (m MockAuthManager) GetPublicKey(host string) (string, error) {
	panic("unimplemented")
}

func (m MockAuthManager) GetOwnPublicKey() *rsa.PublicKey {
	panic("unimplemented")
}