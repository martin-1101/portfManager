package domain

type Asset struct {
	Symbol         string
	Price          float64
	PriceChange    float64
	PriceChangePct float64
	MarketCap      float64
	Volume         float64
}

type AssetRepository interface {
	GetAssets() ([]Asset, error)
}
