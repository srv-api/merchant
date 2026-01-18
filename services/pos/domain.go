package pos

import (
	dto "github.com/srv-api/merchant/dto"
	m "github.com/srv-api/middlewares/middlewares"

	r "github.com/srv-api/merchant/repositories/pos"
)

type PosService interface {
	Create(req dto.PosRequest) (dto.PosResponse, error)
}

type posService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewPosService(Repo r.DomainRepository, jwtS m.JWTService) PosService {
	return &posService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
