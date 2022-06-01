package bd

import (
	"context"
	"time"

	"github.com/alexchocho/CursoGo/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("Twittor")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)

	if err != nil {
		return false, err
	}

	return true, nil
}
