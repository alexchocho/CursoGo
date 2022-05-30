package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/alexchocho/CursoGo/bd"
	"github.com/alexchocho/CursoGo/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Error en parametro recibido "+err.Error(), 400)
		return
	}
	registro := models.GraboTweet{
		UserID: IDUsuario,
		Mesaje: mensaje.Mensaje,
		Fecha:  time.Now(),
	}

	var status bool
	_, status, err = bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Error al intentar insertar el registro: "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se logrado grabar el Tweer", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
