package teams

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func TeamsHandlers(r chi.Router) {
	r.Get("/teams", func(w http.ResponseWriter, res *http.Request) {
		credential := res.Header["Authentication"][0] // <-- cuidar el estandar http
		// TODO: anadir el paso de autenticacion para la api
		if credential == "123" {
			w.Write([]byte("exitoso"))
		} else {
			w.Write([]byte("Credenciales invalidas"))
			res.Response.Status = "405"
		}
	})
}
