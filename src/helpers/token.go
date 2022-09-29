package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecrets = []byte(os.Getenv("JWT_KEYS"))

type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// digunakan di service
func NewToken(username, role string) *Claims {
	return &Claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
}

func (c *Claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString(mySecrets)
}

// lempar ke authValidate
func CheckToken(token, role string) (bool, error) {
	tokens, err := jwt.ParseWithClaims(token, &Claims{Role: role}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecrets), nil
	})

	if err != nil {
		return false, err
	}

	claim := tokens.Claims.(*Claims)
	if claim.Role == role {
		return tokens.Valid, nil
	} else {
		if claim.Role == "admin" {
			return tokens.Valid, nil
		} else {
			return false, err
		}
	}
}

// juga digunakan di authValidate
func EksportToken(token string) (*Claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecrets), nil
	})

	if err != nil {
		return nil, err
	}

	claim := tokens.Claims.(*Claims)
	return claim, nil
}
