package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"

	"github.com/stretchr/testify/assert"
)

func generateKeyPair(t *testing.T, bits int) (*rsa.PrivateKey, *rsa.PublicKey) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		t.Fatal(err)
	}
	return privateKey, &privateKey.PublicKey
}

func generateKeyPairPEMs(t *testing.T, bits int) ([]byte, []byte) {
	privateKey, publicKey := generateKeyPair(t, bits)

	privateBuf := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	pubASN1, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
	}

	publicBuf := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubASN1,
	})

	return privateBuf, publicBuf
}

func TestParsePrivateKeyPEM(t *testing.T) {
	buf, _ := generateKeyPairPEMs(t, 512)
	_, err := ParsePrivateKeyPEM(buf)
	assert.NoError(t, err)
}

func TestParsePublicKeyPEM(t *testing.T) {
	_, buf := generateKeyPairPEMs(t, 512)
	_, err := ParsePublicKeyPEM(buf)
	assert.NoError(t, err)

}

func TestJWTVerifySigning(t *testing.T) {
	privateKey, publicKey := generateKeyPair(t, 512)

	token, err := Sign(privateKey, nil)
	assert.NoError(t, err)

	_, err = Parse(publicKey, token)
	assert.NoError(t, err)
}

func TestJWTVerifySigningClaims(t *testing.T) {
	privateKey, publicKey := generateKeyPair(t, 512)

	token, err := Sign(privateKey, &Claims{Issuer: "test"})
	assert.NoError(t, err)

	claims, err := Parse(publicKey, token)
	assert.NoError(t, err)
	assert.Equal(t, claims.Issuer, "test")
}
