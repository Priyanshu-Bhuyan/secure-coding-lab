package main

import (
	"os"
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestRejectAlgNone(t *testing.T) {
	os.Setenv("JWT_SECRET", "testsecret")

	token := jwt.New(jwt.SigningMethodNone)

	tokenString, _ := token.SignedString(jwt.UnsafeAllowNoneSignatureType)

	err := ValidateJWT(tokenString)

	if err == nil {
		t.Fatal("alg:none token accepted")
	}
}
