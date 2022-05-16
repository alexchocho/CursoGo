package routers

import (
	"encoding/json"
	"net/http"

	"github.com/alexchocho/CursoGo/bd"
	"github.com/alexchocho/CursoGo/models"
)

func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar una contraseña de almenos 6 caracteres", 400)
		return
	}
	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "Email ya registrado", 400)
		return
	}
	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el Registro del Usuario"+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del Usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
