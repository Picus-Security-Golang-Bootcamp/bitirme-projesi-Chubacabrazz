package jwt

import (
	"fmt"
	"strings"
	"time"

	"github.com/Chubacabrazz/picus-storeApp/pkg/config"
	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateJWT(email string) (string, error)
	DecodeJWT(tokenString string) (*TokenClaim, bool, error)
}

type TokenClaim struct {
	Email string `json:"Email"`
	jwt.StandardClaims
}

type jwtService config.JWTConfig

func (user *jwtService) GenerateJWT(email string) (string, error) {

	jwtKey := []byte(user.SecretKey)

	claims := TokenClaim{
		email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(user.SessionTime) * time.Minute).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	return tokenString, err
}

func (user *jwtService) DecodeJWT(tokenString string) (*TokenClaim, bool, error) {

	var secret = []byte(user.SecretKey)

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "expired") {
			return nil, false, nil
		}
		return nil, false, err
	}

	if claims, ok := token.Claims.(*TokenClaim); ok {
		return claims, token.Valid, nil
	}

	return nil, false, err
}

func NewJWTService(SecretKey string, SessionTime int) JWTService {
	return &jwtService{SecretKey: SecretKey, SessionTime: SessionTime}
}
