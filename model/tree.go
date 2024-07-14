package model

import (
	"github.com/google/uuid"
)

type (
	Tree struct {
		ID       uuid.UUID `db:"id"`
		EstateID uuid.UUID `db:"estate_id"`
		X        int       `db:"x"`
		Y        int       `db:"y"`
		Height   int       `db:"height"`
	}
)
