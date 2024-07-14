package handler

import (
	"github.com/SawitProRecruitment/UserService/service"
)

type Server struct {
	// Repository repository.RepositoryInterface
	Service service.ServiceInterface
}

type NewServerOptions struct {
	// Repository repository.RepositoryInterface
	Service service.ServiceInterface
}

func NewServer(opts NewServerOptions) *Server {
	return &Server{
		Service: opts.Service,
	}
}
