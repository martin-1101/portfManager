package infrastructure

import "github.com/martin-1101/portfManager/domain"

type AssetData struct{}

func (ad *AssetData) GetAssets() ([]domain.Asset, error) {
	// Aquí es donde buscarías los datos de los activos en tiempo real
	// desde una fuente externa o una base de datos.
	// Por ahora, simplemente devolveremos algunos activos de ejemplo.

	assets := []domain.Asset{
		{
			Symbol:         "AAPL",
			Price:          135.20,
			PriceChange:    -0.75,
			PriceChangePct: -0.55,
			MarketCap:      2.3e12,
			Volume:         9.2e7,
		},
		{
			Symbol:         "GOOGL",
			Price:          2100.0,
			PriceChange:    5.25,
			PriceChangePct: 0.25,
			MarketCap:      1.4e12,
			Volume:         1.3e6,
		},
	}

	return assets, nil
}
