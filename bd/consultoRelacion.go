package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/alexchocho/CursoGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()
	db := MongoCN.Database("Twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)

	if err != nil {
		fmt.Print(err.Error())
		return false, err
	}

	return true, nil
}
