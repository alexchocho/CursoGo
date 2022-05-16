package main

import (
	"log"

	"github.com/alexchocho/CursoGo/bd"
	"github.com/alexchocho/CursoGo/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin Conexion a la Base de Datos")
		return
	}
	handlers.Manejadores()

}
