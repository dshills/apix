package token

import (
	"testing"
	"time"
)

//TestToken is a test
func TestToken(t *testing.T) {
	key := []byte("26BF237B95964852625A2C27988C3")
	SetSecret(key)
	c := NewClaims(1, 15*time.Minute)
	c.SetIssuer("token_test")
	c.SetSubject("test")
	tok, err := c.Token()
	if err != nil {
		t.Fatal(err)
	}

	c, err = Decode(tok)
	if err != nil {
		t.Fatal(err)
	}
}
