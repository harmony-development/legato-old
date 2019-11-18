package handler

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
	"time"

	"github.com/bluskript/harmony-server/globals"
	"github.com/bluskript/harmony-server/socket"
	"github.com/dgrijalva/jwt-go"
	. "github.com/logrusorgru/aurora"
	"github.com/thanhpk/randstr"
)

type registerData struct {
	email    string
	username string
	password string
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

// RegisterHandler handles user registration in harmony
func RegisterHandler(raw interface{}, ws *socket.WebSocket) {
	if globals.HarmonyServer.DatabaseInstance == nil {
		log.Println(Red("Database Instance not available").Bold())
		return
	}

	rawMap, ok := raw.(map[string]interface{})
	if ok {
		var data registerData
		data.email, ok = rawMap["email"].(string)
		if !ok {
			regErr("Missing Email", ws)
			return
		}
		data.username, ok = rawMap["username"].(string)
		if !ok {
			regErr("Missing Username", ws)
			return
		}
		data.password, ok = rawMap["password"].(string)
		if !ok {
			regErr("Missing Password", ws)
			return
		}
		if !emailRegex.MatchString(data.email) {
			regErr("Invalid Email", ws)
			return
		}

		if len(data.username) >= 20 {
			regErr("Your name is too long", ws)
			return
		}

		if len(data.password) < 3 || len(data.password) > 30 {
			regErr("Password must be at least 3 characters and at most 30 characters long", ws)
			return
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(data.password), bcrypt.DefaultCost)

		if err != nil {
			whoops("REGISTER_ERROR", ws)
			return
		}

		// all inputs are valid here
		userid := randstr.Hex(16)
		_, err = globals.HarmonyServer.DatabaseInstance.Exec("INSERT INTO users (id, email, password, username) VALUES (?, ?, ?, ?)", userid, data.email, hash, data.username)

		if err != nil {
			log.Println(Red(err.Error()).Bold())
			regErr("Something went wrong while registering. The email specified might already be registered", ws)
			return
		}

		expireTime := time.Now().Add(30 * 24 * time.Hour).UTC()
		claims := &AuthToken{
			Userid: userid,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expireTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(globals.HarmonyServer.JwtSecret))
		if err != nil {
			log.Println(Red(err.Error()).Bold())
			regErr("Unable to create token", ws)
			return
		}
		register(tokenString, ws)
	}
}
