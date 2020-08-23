package createjwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Createjwttoken it return token string and err
func Createjwttoken(u JwtCustomClaims) (string, error) {

	claims := &JwtCustomClaims{
		u.ID,
		u.Email,
		u.RolesID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secretToken"))
	if err != nil {
		return "", err
	}
	return t, err

}
