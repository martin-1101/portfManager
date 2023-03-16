package presentation

import (
	"fmt"

	"github.com/martin-1101/portfManager/domain"
)

func DisplayAssets(assets []domain.Asset) {
	// Aquí mostrarías la tabla en la CLI
	for _, asset := range assets {
		fmt.Printf("%s\t%.2f\t%.2f\t%.2f%%\t%.0f\t%.0f\n",
			asset.Symbol, asset.Price, asset.PriceChange, asset.PriceChangePct,
			asset.MarketCap, asset.Volume)
	}
}
