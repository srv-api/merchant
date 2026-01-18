package export_data

import (
	"context"

	dto "github.com/srv-api/merchant/dto"
	"github.com/srv-api/product/entity"
	"gorm.io/gorm"
)

type DomainRepository interface {
	FindAllByFilter(ctx context.Context, req dto.ExportFilter) ([]entity.Product, error)
}

type exportRepository struct {
	DB *gorm.DB
}

func NewExportRepository(DB *gorm.DB) DomainRepository {
	return &exportRepository{
		DB: DB,
	}
}
