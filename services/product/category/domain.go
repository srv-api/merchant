package dashboard

import (
	m "github.com/srv-api/middlewares/middlewares"

	"github.com/srv-api/merchant/dto"
	r "github.com/srv-api/merchant/repositories/product/category"
)

type GetCategoryService interface {
	Get(req dto.CategoryRequest) ([]dto.CategoryResponse, error)
}

type getCategorydashboardService struct {
	Repo r.DomainRepository
	jwt  m.JWTService
}

func NewGetCategoryService(Repo r.DomainRepository, jwtS m.JWTService) GetCategoryService {
	return &getCategorydashboardService{
		Repo: Repo,
		jwt:  jwtS,
	}
}
