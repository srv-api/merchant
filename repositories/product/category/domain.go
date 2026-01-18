package category

import (
	"github.com/srv-api/merchant/dto"
	"gorm.io/gorm"
)

type DomainRepository interface {
	Get(req dto.CategoryRequest) ([]dto.CategoryResponse, error)
}

type getcategoryRepository struct {
	DB *gorm.DB
}

func NewGetCategoryRepository(DB *gorm.DB) DomainRepository {
	return &getcategoryRepository{DB: DB}
}
