package token

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const minSecretKeySize = 32

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrExpiredToken
		}
		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(tokenString, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(jwt.ClaimsValidator)
		if ok && errors.Is(verr.Validate(), ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

// VerifyToken checks if the token is valid or not
//func (maker *JWTMaker) VerifyToken(tokenString string) (*Payload, error) {
//	// Parse the token without validating the signature
//	jwtToken, err := jwt.ParseWithClaims(tokenString, &Payload{}, func(token *jwt.Token) (interface{}, error) {
//		// Check the signing method
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, ErrInvalidToken
//		}
//		return []byte(maker.secretKey), nil
//	})
//
//	if err != nil {
//		return nil, ErrInvalidToken
//	}
//
//	// Check if the token is valid
//	if !jwtToken.Valid {
//		return nil, ErrInvalidToken
//	}
//
//	// Extract claims
//	payload, ok := jwtToken.Claims.(*Payload)
//	if !ok {
//		return nil, ErrInvalidToken
//	}
//
//	return payload, nil
//}
