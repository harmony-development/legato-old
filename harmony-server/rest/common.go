package rest

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"net/http"
	"os"
	"time"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func sendResp(w http.ResponseWriter, data interface{}) {
	marshalled, err := json.Marshal(data)
	if err != nil {
		golog.Warnf("Error sending JSON Response : %v", err)
		return
	}
	_, _ = w.Write(marshalled)
}

// sourced from https://golangcode.com/get-the-request-ip-addr/
func getIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func makeToken(contents string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  contents,
		"exp": time.Now().UTC().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		golog.Warnf("Error signing token. Reason : %v", err) // pray to god this never happens
		return nil, err
	}
	return &tokenString, nil
}

func deleteFromFilestore(fileid string) {
	err := os.Remove(fmt.Sprintf("./filestore/%v", fileid))
	if err != nil {
		golog.Warnf("Error deleting from filestore : %v", err)
	}
}