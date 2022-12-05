package tools

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

var clave = []byte("mi clave super secreta")

// generar un nuevo jwt para la autenticacion del usuario
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": username,
	})
	res, err := token.SignedString(clave)
	if err != nil {
		return "", err
	}
	return res, nil
}

// comprueba el estado del jwt para revisar
func ComprobarJWT(receivedToken string) error {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return clave, nil
	})
	// mal parseo
	if err != nil {
		return err
	}
	// token invalido
	if _, ok := token.Claims.(jwt.MapClaims); ok != true || token.Valid == false {
		return fmt.Errorf("Token invalido")
	}
	return nil
}

// jwt midleware to protect authentication
func JwtMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// no proteger "/user/" porque no necesita token para crear o iniciar sesion
		if r.URL.Path == "/user/" {
			next.ServeHTTP(w, r)
			return
		}
		// extraer el token
		token, err := extractToken(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
			return
		}
		// comprobar el token
		if err := ComprobarJWT(token); err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token"))
			return
		}
		// pasar al siguiente midleware
		next.ServeHTTP(w, r)
	})
}

func extractToken(r *http.Request) (string, error) {
	// get the token from the request
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
	// comprobar el estado del token
	if len(authHeader) != 2 {
		return "", fmt.Errorf("Malformed Token")
	}
	// comprobar si el token es correcto
	token := strings.TrimPrefix(authHeader[1], ": ") // eliminar esta parte de la request
	return token, nil
}
