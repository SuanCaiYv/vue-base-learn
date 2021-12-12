package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func main() {
	// InitSysRoleDB()
	signingString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "test",
		Subject:   "access",
		Audience:  []string{0: "aaa", 1: "bbb"},
		ExpiresAt: &jwt.NumericDate{Time: time.UnixMilli(time.Now().Unix() + 2*time.Hour.Milliseconds())},
		NotBefore: &jwt.NumericDate{Time: time.Now()},
		IssuedAt:  &jwt.NumericDate{Time: time.Now()},
		ID:        "0001",
	}).SigningString()
	fmt.Println(signingString, err)
}
