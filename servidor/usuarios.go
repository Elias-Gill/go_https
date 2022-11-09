package servidor

import "errors"

func IniciarSesion(usuario string, contrasena string) (string, error) {
	if usuario == "elias" && contrasena == "123" {
		return "123", nil
	}
	return "", errors.New("Usuario o contrasena invalido")
}
