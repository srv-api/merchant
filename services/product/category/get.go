package dashboard

import "github.com/srv-api/merchant/dto"

func (s *getCategorydashboardService) Get(req dto.CategoryRequest) ([]dto.CategoryResponse, error) {
	category, err := s.Repo.Get(req)
	if err != nil {
		return nil, err
	}

	return category, nil
}
