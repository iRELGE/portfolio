package createjwt

import "github.com/dgrijalva/jwt-go"

// JwtCustomClaims : struct store all attribute claimed on jwt
type JwtCustomClaims struct {
	ID      int    `boil:"ID" json:"ID" toml:"ID" yaml:"ID"`
	Email   string `boil:"Email" json:"Email,omitempty" toml:"Email" yaml:"Email,omitempty"`
	RolesID int    `boil:"roles_id" json:"roles_id" toml:"roles_id" yaml:"roles_id"`

	jwt.StandardClaims
}
