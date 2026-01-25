package product

import (
	"regexp"
	"strings"

	dto "github.com/srv-api/merchant/dto"
)

func (b *merchantService) Update(req dto.UpdateMerchantRequest) (dto.UpdateMerchantResponse, error) {
	merchantSlug := generateMerchantSlug(req.MerchantName)

	request := dto.UpdateMerchantRequest{
		ID:           req.ID,
		MerchantName: req.MerchantName,
		MerchantSlug: merchantSlug,
		IDNumber:     req.IDNumber,
		Description:  req.Description,
		Address:      req.Address,
		Country:      req.Country,
		CurrencyID:   req.CurrencyID,
		City:         req.City,
		Zip:          req.Zip,
		Phone:        req.Phone,
		UpdatedBy:    req.UpdatedBy,
	}

	product, err := b.Repo.Update(req)
	if err != nil {
		return product, err
	}

	response := dto.UpdateMerchantResponse{
		ID:           req.ID,
		MerchantName: request.MerchantName,
		MerchantSlug: request.MerchantSlug,
		IDNumber:     request.IDNumber,
		Description:  request.Description,
		Address:      request.Address,
		Country:      request.Country,
		CurrencyID:   request.CurrencyID,
		City:         request.City,
		Zip:          request.Zip,
		Phone:        request.Phone,
		UpdatedBy:    request.UpdatedBy,
	}

	return response, nil
}

func generateMerchantSlug(name string) string {
	slug := strings.ToLower(strings.TrimSpace(name))

	reg := regexp.MustCompile(`[^a-z0-9]+`)
	slug = reg.ReplaceAllString(slug, "-")

	regDash := regexp.MustCompile(`-+`)
	slug = regDash.ReplaceAllString(slug, "-")

	return strings.Trim(slug, "-")
}
