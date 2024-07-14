// This file contains the service implementation layer.
package service

import (
	"github.com/SawitProRecruitment/UserService/repository"
)

type Service struct {
	repo repository.RepositoryInterface
}

func NewService(repo repository.RepositoryInterface) *Service {
	return &Service{repo: repo}
}
