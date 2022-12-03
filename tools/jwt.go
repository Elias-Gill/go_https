package tools

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

// generar un nuevo jwt para la autenticacion del usuario
func GenerateJWT(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userName": username,
	})
	res, err := token.SignedString([]byte("mi clave super secreta"))
	if err != nil {
		return "", err
	}
	return res, nil
}

// comprueba el estado del jwt para revisar
func ComprobarJWT(receivedToken string) error {
	token, err := jwt.Parse(receivedToken, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tk.Header["alg"])
		}
		return tk, nil
	})

	// comprobar el estado del jwt
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	} else {
		return err
	}
}

// TODO: terminar
// jwt midleware to protect authentication
func JwtMidleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/user/" {
			print("\n" + r.URL.Path)
			next.ServeHTTP(w, r)
			return
		}
		// get the token from the request
		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer")
		// comprobar el estado del token
		if len(authHeader) != 2 {
            fmt.Println("Malformed token: "+ authHeader[1])
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
			return
		}
		// comprobar si el token es correcto
		token := authHeader[1]
		if err := ComprobarJWT(token); err != nil {
			fmt.Println("Unauthorized: " + err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		// pasar al siguiente midleware
		next.ServeHTTP(w, r)
	})
}
