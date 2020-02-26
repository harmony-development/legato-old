package rest

import (
	"encoding/json"
	"github.com/kataras/golog"
	"golang.org/x/crypto/bcrypt"
	"harmony-server/harmonydb"
	"net/http"
)

type LoginData struct {
	Email *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	WithCors(w)
	var data LoginData
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "missing POST body", http.StatusBadRequest)
	}
	if data.Email == nil || data.Password == nil {
		http.Error(w, "invalid username or password", http.StatusBadRequest)
	}
	var passwd, userid string
	if err = harmonydb.DBInst.QueryRow("SELECT password, id from users WHERE email=$1", data.Email).Scan(&passwd, &userid); err != nil {
		http.Error(w, "invalid email or password", http.StatusNotAcceptable)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(passwd), []byte(*data.Password)); err != nil {
		http.Error(w, "invalid email or password", http.StatusUnauthorized)
		return
	}
	if !getVisitor("login", getIP(r)).Allow() {
		http.Error(w, "you're sending too many login requests", http.StatusTooManyRequests)
	}
	token, err := makeToken(userid)
	if err != nil || token == nil {
		http.Error(w, "error generating token", http.StatusInternalServerError)
		return
	}
	_, err = w.Write([]byte(*token))
	if err != nil {
		golog.Warnf("error occurred while sending token : %v", err)
	}
}