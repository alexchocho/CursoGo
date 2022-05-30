package models

import "time"

type GraboTweet struct {
	UserID string    `bson:"userid" json:"userid,omitempty"`
	Mesaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha  time.Time `bson:"fecha" json:"fecha,omitempty"`
}
