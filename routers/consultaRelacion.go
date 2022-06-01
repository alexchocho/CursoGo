package routers

import (
	"encoding/json"
	"net/http"

	"github.com/alexchocho/CursoGo/bd"
	"github.com/alexchocho/CursoGo/models"
)

func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	var resp models.RespuestaCosultaRelacion

	status, err := bd.ConsultoRelacion(t)

	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
