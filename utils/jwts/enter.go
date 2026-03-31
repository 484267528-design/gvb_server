package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayload struct {
	//Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     int    `json:"role"`
	UserID   int    `json:"user_id"`
}

var MySecret []byte

type CustomClaims struct {
	JwtPayload
	jwt.StandardClaims
}
