package server

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

// generar un nuevo jwt para la autenticacion del usuario
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"nombre": username,
	})
	res, err := token.SignedString([]byte("mi clave super secreta"))
	if err != nil {
		return "", err
	}
	return res, nil
}



// TODO: YA NO SE QUE estoy haciendo porque che kane'o
func ComprobarJWT(receivedToken string) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return "hola", nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["email"], claims["birthday"])
	} else {
		fmt.Println(err)
	}
}
