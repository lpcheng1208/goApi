package request

import (
	"github.com/dgrijalva/jwt-go"
)

// Custom claims structure
type CustomClaims struct {
	UserId  string
	UserRid int64
	jwt.StandardClaims
}
