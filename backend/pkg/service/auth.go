package service

import (
	"backend/pkg/model"
	"backend/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	tokenTTL = 12 * time.Hour
)

var salt, signingKey string

func init() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	signingKey = os.Getenv("SigningKey")
	salt = os.Getenv("SALT")
}

type tokenClaims struct {
	jwt.StandardClaims
	Username string
	Roles    []string `json:"roles"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User, role string) (int, error) {
	user.Password = GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user, role)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	roles, err := s.repo.GetUser(username, GeneratePasswordHash(password))
	if err != nil {
		return "", err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		username,
		roles,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) GenerateOAthToken(username string, roles []string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		username,
		roles,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (string, []string, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return "", nil, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return "", nil, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.Username, claims.Roles, nil
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
