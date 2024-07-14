// This file contains the interfaces for the service layer.
// The service layer is responsible for interacting with the repository layer.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package service

import (
	"github.com/SawitProRecruitment/UserService/model"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	AddEstate(width, length int) (uuid.UUID, error)
	AddTree(estateID uuid.UUID, x, y, height int) (uuid.UUID, error)
	GetEstateStats(estateID uuid.UUID) (int, int, int, int, error)
	GetDronePlanDistance(estateID uuid.UUID) (int, error)
	GetDronePlanMaxDistance(estateID uuid.UUID, maxDistance int) (int, model.Coordinate, error)
}
