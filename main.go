package main

import (
	"github.com/martin-1101/portfManager/application"
	"github.com/martin-1101/portfManager/infrastructure"
	"github.com/martin-1101/portfManager/presentation"
)

func main() {
	assetRepo := &infrastructure.AssetData{}
	assetService := application.NewAssetService(assetRepo)
	cli := presentation.NewCLI(assetService)

	cli.Run()
}
