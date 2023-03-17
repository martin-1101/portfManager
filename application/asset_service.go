package application

import "github.com/martin-1101/portfManager/domain"

type AssetService struct {
	repo domain.AssetRepository
}

func NewAssetService(repo domain.AssetRepository) *AssetService {
	return &AssetService{repo: repo}
}

func (s *AssetService) GetAllAssets() ([]domain.Asset, error) {
	return s.repo.GetAssets()
}
