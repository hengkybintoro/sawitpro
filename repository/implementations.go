package repository

import (
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
)

// func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
// 	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

func (r *Repository) AddEstate(width, length int) (uuid.UUID, error) {
	id := uuid.New()
	_, err := r.Db.Exec("INSERT INTO estates (id, width, length) VALUES ($1, $2, $3)", id, width, length)
	return id, err
}

func (r *Repository) GetEstate(id uuid.UUID) (*model.Estate, error) {
	var estate model.Estate
	err := r.Db.Get(&estate, "SELECT * FROM estates WHERE id=$1", id)
	return &estate, err
}

func (r *Repository) AddTree(estateID uuid.UUID, x, y, height int) (uuid.UUID, error) {
	id := uuid.New()
	_, err := r.Db.Exec("INSERT INTO trees (id, estate_id, x, y, height) VALUES ($1, $2, $3, $4, $5)", id, estateID, x, y, height)
	return id, err
}

func (r *Repository) GetTreesByEstate(estateID uuid.UUID) ([]model.Tree, error) {
	var trees []model.Tree
	err := r.Db.Select(&trees, "SELECT * FROM trees WHERE estate_id=$1", estateID)
	return trees, err
}
