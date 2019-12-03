package jwt

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"

	"github.com/dgrijalva/jwt-go"
)

/*
# Generate private key for signing by auth service.
ssh-keygen -t rsa -b 4096 -m PEM -f jwt-rs256.key

# Generate public key for verifying by consuming.
openssl rsa -in jwt-rs256.key -pubout -outform PEM -out jwt-rs256.key.pub
*/

func ParsePrivateKeyPEM(b []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(b)
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func ReadPrivateKeyPEM(path string) (*rsa.PrivateKey, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParsePrivateKeyPEM(b)
}

func ParsePublicKeyPEM(b []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(b)
	v, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	if k, ok := v.(*rsa.PublicKey); ok {
		return k, nil
	}
	return nil, errors.New("not a public key")
}

func ReadPublicKeyPEM(path string) (*rsa.PublicKey, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return ParsePublicKeyPEM(b)
}

type Claims struct {
	Audience  string `json:"aud,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Id        string `json:"jti,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	Issuer    string `json:"iss,omitempty"`
	NotBefore int64  `json:"nbf,omitempty"`
	Subject   string `json:"sub,omitempty"`
}

// Sign creates and signs a JWT given the claims. It uses the RS256 signing method
// expecting a private key.
func Sign(key *rsa.PrivateKey, claims *Claims) (string, error) {
	var stdclaims *jwt.StandardClaims
	if claims != nil {
		// Convert to jwt-go type for validation.
		x := jwt.StandardClaims(*claims)
		stdclaims = &x
	}

	// Fixed signing method.
	tok := jwt.NewWithClaims(jwt.SigningMethodRS256, stdclaims)

	return tok.SignedString(key)
}

// Parse parses a signed JWT. It requires a public key to verify the signing.
func Parse(key *rsa.PublicKey, token string) (*Claims, error) {
	tok, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(tok *jwt.Token) (interface{}, error) {
		if _, ok := tok.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("wrong signing method")
		}
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	stdclaims, _ := tok.Claims.(*jwt.StandardClaims)
	if stdclaims != nil {
		claims := Claims(*stdclaims)
		return &claims, nil
	}
	return nil, nil
}
