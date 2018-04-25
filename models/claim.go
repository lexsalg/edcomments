package models

import jwt "github.com/dgrijalva/jwt-go"

// Claim clase claim
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
