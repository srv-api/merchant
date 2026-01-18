package dashboard

import (
	m "github.com/srv-api/middlewares/middlewares"

	"github.com/srv-api/merchant/dto"
	r "github.com/srv-api/merchant/repositories/dashboard"
)

type DashboardService interface {
	Get(req dto.GetDashboardRequest) (dto.GetDashboardResponse, error)
}

type dashboardService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewDashboardService(Repo r.DomainRepository, jwtS m.JWTService) *dashboardService {
	s := &dashboardService{
		Repo: Repo,
	}

	return s
}
