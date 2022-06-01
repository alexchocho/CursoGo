package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/alexchocho/CursoGo/bd"
)

func ObtenerBanner(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no Encontrado", http.StatusBadRequest)
		return
	}
	OpenFile, err := os.Open("uploads/banners/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no Encontrado", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error al copiar Imagen", http.StatusBadRequest)
		return
	}

}
