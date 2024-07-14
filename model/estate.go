package model

import (
	"github.com/google/uuid"
)

type (
	Estate struct {
		ID     uuid.UUID `db:"id"`
		Width  int       `db:"width"`
		Length int       `db:"length"`
	}
)
