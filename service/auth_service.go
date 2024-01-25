package service

import (
	"crypto/rsa"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/prapsky/sawitpro/entity"
)

type JwtAuthService struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

type JwtAuthServiceOptions struct {
	PrivateKey string
	PublicKey  string
}

type AuthService interface {
	CreateToken(user *entity.User) (string, error)
}

func NewJwtAuthService(opts JwtAuthServiceOptions) *JwtAuthService {
	pem, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(opts.PrivateKey))
	if err != nil {
		panic(err)
	}

	cert, err := jwt.ParseRSAPublicKeyFromPEM([]byte(opts.PublicKey))
	if err != nil {
		panic(err)
	}

	return &JwtAuthService{
		privateKey: pem,
		publicKey:  cert,
	}
}

func (s *JwtAuthService) CreateToken(user *entity.User) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub":  strconv.FormatInt(int64(user.ID), 10),
		"name": user.FullName,
		"iat":  now.Unix(),
		"exp":  now.Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(s.privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
