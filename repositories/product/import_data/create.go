package import_data

import (
	"context"

	"github.com/srv-api/product/entity"
)

func (r *importRepository) Create(ctx context.Context, product *entity.Product) error {
	return r.DB.WithContext(ctx).Create(product).Error
}
