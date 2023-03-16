package infrastructure

import (
	"github.com/martin-1101/portfManager/domain"
)

type AssetData struct{}

func (ad *AssetData) GetAssets() ([]domain.Asset, error) {
	// Aquí obtendrías los datos de los activos en tiempo real, pero usaremos datos hardcodeados
	assets := []domain.Asset{
		{"AAPL", 150.50, 0.35, 0.23, 2.5e12, 5e6},
		{"GOOGL", 1050.25, -5.10, -0.48, 1.5e12, 2e6},
		{"TSLA", 900.85, 10.25, 1.15, 1e12, 8e6},
	}
	return assets, nil
}
