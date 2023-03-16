package main

import (
	"github.com/martin-1101/portfManager/application"
	"github.com/martin-1101/portfManager/infrastructure"
	"github.com/martin-1101/portfManager/presentation"
)

func main() {
	assetData := &infrastructure.AssetData{}
	assetService := &application.AssetService{Repo: assetData}
	assets, _ := assetService.GetAssets()

	presentation.DisplayAssets(assets)
}
