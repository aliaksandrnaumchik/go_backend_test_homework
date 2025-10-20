package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "qw45jk32"

// Claims хранит идентификатор пользователя в токене
type Claims struct {
	ID                   int    // идентификатор пользователя
	Name                 string // идентификатор пользователя
	jwt.RegisteredClaims        // базовый тип
}

// CreateToken создаёт токен
func CreateToken(id int, name string) (string, error) {
	// Создаём структуру Claims
	claims := Claims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // токен действителен 24 часа
		},
	}

	// Создаём токен с алгоритмом HMAC с SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken проверяет валидность JWT-токена и возвращает идентификатор пользователя
func ValidateToken(tokenString string, claims *Claims) bool {
	// парсим токен
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// проверяем алгоритм подписи
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неверный метод подписи")
		}
		return []byte(secretKey), nil
	})
	if err != nil || !token.Valid {
		return false
	}
	return true
}

func main() {
	for i, name := range []string{"Андрей", "Олег", "Наталья"} {
		token, err := CreateToken(i+1, name)
		if err != nil {
			return
		}
		var claim Claims
		if ValidateToken(token, &claim) {
			fmt.Println(claim.ID, claim.Name)
		}
	}
}
