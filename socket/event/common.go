package event

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/golog"
	"harmony-server/socket"
	"os"
	"time"
)

var jwtSecret = os.Getenv("JWT_SECRET")

func regErr(ws *socket.Client, msg string) {
	ws.Send(&socket.Packet{
		Type: "RegisterError",
		Data: map[string]interface{}{
			"message": msg,
		},
	})
}

func loginErr(ws *socket.Client, msg string) {
	ws.Send(&socket.Packet{
		Type: "LoginError",
		Data: map[string]interface{}{
			"message": msg,
		},
	})
}

func deauth(ws *socket.Client) {
	ws.Send(&socket.Packet{
		Type: "Deauth",
		Data: map[string]interface{}{
			"message": "token is missing or invalid",
		},
	})
}

func sendToken(ws *socket.Client, id string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"exp": time.Now().UTC().Add(7 * 24 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		golog.Warnf("Error signing token. Reason : %v", err) // pray to god this never happens
		return
	}

	ws.Send(&socket.Packet{
		Type: "Token",
		Data: map[string]interface{}{
			"token": tokenString,
		},
	})
}

func verifyToken(tokenstr string) string {
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["id"].(string)
	} else {
		return ""
	}
}