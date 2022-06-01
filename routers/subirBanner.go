package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/alexchocho/CursoGo/bd"
	"github.com/alexchocho/CursoGo/models"
)

func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, _ := r.FormFile("banner")

	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al intentar subir el banner! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar el banner! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUsuario + "." + extension

	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || !status {
		http.Error(w, "Error al grabar el banner en la BD! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
