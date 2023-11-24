package define

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var (
	Jwtkey        = "sys-admin"
	TokenExpire   = time.Now().Add(time.Second * 3600 * 24 * 7).Unix()
	RefreshExpire = time.Now().Add(time.Second * 3600 * 24 * 14).Unix()
	DefaultSize   = 10
)

type UserClaim struct {
	Id      uint
	Name    string
	IsAdmin bool
	jwt.StandardClaims
}
