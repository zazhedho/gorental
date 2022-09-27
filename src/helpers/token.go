package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySecrets = []byte(os.Getenv("JWT_KEYS"))

type claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

// digunakan di service
func NewToken(username, role string) *claims {
	return &claims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
}

func (c *claims) Create() (string, error) {
	tokens := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tokens.SignedString(mySecrets)
}

// lempar ke authValidate
func CheckToken(token, role string) (bool, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{Role: role}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecrets), nil
	})

	if err != nil {
		return false, err
	}

	claim := tokens.Claims.(*claims)
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
func EksportToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(mySecrets), nil
	})

	if err != nil {
		return nil, err
	}

	claim := tokens.Claims.(*claims)
	return claim, nil
}
