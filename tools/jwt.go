package tools

import (
	"fmt"
	"log"
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

// Comprobar el jwt, parsearlo, revisar la firma y la validez del token
func ComprobarJWT(receivedToken string) (string, error) {
	token, err := jwt.Parse(receivedToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return clave, nil
	})
	// mal parseo
	if err != nil {
		return "", err
	}
	// token invalido
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok != true || token.Valid == false {
		return "", fmt.Errorf("Token invalido o claims corrompidas")
	}
	return claims["username"].(string), nil
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
			log.Printf("Cannot extract token: %s\n", err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
			return
		}
		// comprobar el token
		user, err := ComprobarJWT(token)
		if err != nil {
			log.Printf("Token authentication. Token invalid (%s): %s\n", token, err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Invalid token or bad auth header format"))
			return
		}
		// guardo el username extraido para usar mas tarde en la request
		r.SetBasicAuth(user, "")
		// pasar al siguiente midleware
		next.ServeHTTP(w, r)
	})
}

// Funcion para extraer el token de una llamada http con Authorization Bearer
// como header
func extractToken(r *http.Request) (string, error) {
	// get the token from the request
	authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
	// comprobar el estado del token
	if len(authHeader) != 2 {
		return "", fmt.Errorf("Malformed Token")
	}
	// comprobar si el token es correcto
	token := strings.TrimPrefix(authHeader[1], ":") // eliminar esta parte de la request
	token = strings.TrimPrefix(token, " ")          // eliminar ese espacio raro
	return token, nil
}
