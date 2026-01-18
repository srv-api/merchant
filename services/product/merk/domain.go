package dashboard

import (
	m "github.com/srv-api/middlewares/middlewares"

	"github.com/srv-api/merchant/dto"
	r "github.com/srv-api/merchant/repositories/product/merk"
)

type GetMerkService interface {
	Get(req dto.MerkRequest) ([]dto.MerkResponse, error)
}

type getMerkdashboardService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewGetMerkService(Repo r.DomainRepository, jwtS m.JWTService) GetMerkService {
	return &getMerkdashboardService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
