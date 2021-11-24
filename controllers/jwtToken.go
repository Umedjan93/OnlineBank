package controllers

import (
	"OnlineBank/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//GenerateToken - генерирует новый токен/generates new JWT token
func GenerateToken(userId int64, role string) (token string) {
	mySigningKey := []byte("MySecretCode")
	claims := models.MyClaims{
		userId,
		role,
		jwt.StandardClaims{
			//ExpiresAt - выставляет время жизни/действия Токена до конкретного
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
			Issuer:    "jwt",
		},
	}

	token2 := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token2.SignedString(mySigningKey)
	fmt.Println("token:", ss, err)
	return ss
}

//ParsToken - парсинг токена с id и ролью пользователя/parses token
func ParsToken(token string) (int, string, bool) {
	claims := jwt.MapClaims{}
	token1, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("MySecretCode"), nil
	})
	if err != nil {
		return 0, "", false
	}
	if claims, ok := token1.Claims.(jwt.MapClaims); ok && token1.Valid {
		id := claims["id"].(float64)
		return int(id), claims["role"].(string), true
	}
	return 0, "", false
}
