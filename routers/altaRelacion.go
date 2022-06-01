package routers

import (
	"net/http"

	"github.com/alexchocho/CursoGo/bd"
	"github.com/alexchocho/CursoGo/models"
)

func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar relacion", http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "no se ha podido grabar la relacion: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
