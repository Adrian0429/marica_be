package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTService interface {
	GenerateToken(id uuid.UUID, role string) string
	GenerateTokenForgot(id uuid.UUID, role string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetIDByToken(token string) (uuid.UUID, error)
	// GetMandorIDByToken(token string) (uuid.UUID, error)
}

type jwtCustomClaim struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "Template",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "Template"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(id uuid.UUID, role string) string {
	return j.GeneratedToken(id, role)
}

func (j *jwtService) GenerateTokenForgot(id uuid.UUID, role string) string {
	return j.GeneratedToken(id, role)
}

func (j *jwtService) GeneratedToken(id uuid.UUID, role string) string {
	claims := jwtCustomClaim{
		id,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 240)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (j *jwtService) GeneratedTokenForgot(id uuid.UUID, role string) string {
	claims := jwtCustomClaim{
		id,
		role,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 3)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (j *jwtService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, j.parseToken)
}

func (j *jwtService) GetIDByToken(token string) (uuid.UUID, error) {
	t_Token, err := j.ValidateToken(token)
	if err != nil {
		return uuid.Nil, err
	}
	claims := t_Token.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["id"])
	teamID, _ := uuid.Parse(id)
	return teamID, nil
}
