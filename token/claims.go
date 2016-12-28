package token

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Claims is signed data for the token
type Claims struct {
	UserID int64 `json:"user_id,omitempty"`
	jwt.StandardClaims
}

// NewClaims returns a claim for a user id
func NewClaims(userid int64, exp time.Duration) *Claims {
	return &Claims{
		userid,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(exp).Unix(),
		},
	}
}

// SetIssuer sets the issuer claim
func (c *Claims) SetIssuer(issuer string) {
	c.Issuer = issuer
}

// SetSubject sets the subject claim
func (c *Claims) SetSubject(subject string) {
	c.Subject = subject
}

// Token returns a token for claim
func (c *Claims) Token() (string, error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return tok.SignedString(secretKey)
}
