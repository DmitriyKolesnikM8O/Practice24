package jwt

import (
	"context"
	"fmt"
	"github.com/DmitriyKolesnikM8O/Practice24/internal/config"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

var secretKey = []byte(config.GetConfig().SecretKey.Secret)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// для создания токенов
func GenerateJWT(username string) (string, error) {
	// Определите срок жизни токена
	expirationTime := time.Now().Add(24 * time.Hour)

	// Создайте claims
	claims := &MyClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// Создайте токен
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подпишите токен секретным ключом
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// для валидации (проверки) -- подпись и срок действия
func ValidateJWT(tokenString string) (*MyClaims, error) {
	claims := &MyClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Убедитесь, что метод подписи JWT является ожидаемым
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}

// middleware для проверки наличия и корректности
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// Сохраняем claims в контексте запроса, для использования в обработчиках
		ctx := r.Context()
		ctx = context.WithValue(ctx, "user", claims.Username)

		// Передаем управление следующему обработчику
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
