package application

import "github.com/martin-1101/portfManager/domain"

type AssetService struct {
	Repo domain.AssetRepository
}

func (as *AssetService) GetAssets() ([]domain.Asset, error) {
	return as.Repo.GetAssets()
}
