package protocol

type ConnectResponse struct {
	Token string `json:"token"`
	Nonce string `json:"nonce"`
}
