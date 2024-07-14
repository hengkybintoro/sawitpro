// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
)

type RepositoryInterface interface {
	// GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	AddEstate(width, length int) (uuid.UUID, error)
	GetEstate(id uuid.UUID) (*model.Estate, error)
	AddTree(estateID uuid.UUID, x, y, height int) (uuid.UUID, error)
	GetTreesByEstate(estateID uuid.UUID) ([]model.Tree, error)
}
