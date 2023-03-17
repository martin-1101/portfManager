package presentation

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/martin-1101/portfManager/application"
	"github.com/martin-1101/portfManager/domain"
	"github.com/olekukonko/tablewriter"
)

type CLI struct {
	assetService *application.AssetService
}

func NewCLI(assetService *application.AssetService) *CLI {
	return &CLI{assetService: assetService}
}

func (c *CLI) Run() {
	for {
		assets, err := c.assetService.GetAllAssets()
		if err != nil {
			fmt.Println("Error al obtener los activos:", err)
			os.Exit(1)
		}

		c.printAssets(assets)
		time.Sleep(30 * time.Second)
	}
}

func (c *CLI) printAssets(assets []domain.Asset) {
	// Limpiar pantalla
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	// Crear tabla
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Symbol", "Price", "Change", "Change%", "Market Cap", "Volume"})
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold}, tablewriter.Colors{tablewriter.Bold})
	table.SetColumnColor(tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{}, tablewriter.Colors{})
	table.SetBorder(true)
	table.SetColumnSeparator("+")
	table.SetRowLine(true)
	table.SetAutoWrapText(false)

	// Rellenar tabla con datos
	for _, asset := range assets {
		priceChangeSign := "+"
		if asset.PriceChange < 0 {
			priceChangeSign = ""
		}
		row := []string{
			asset.Symbol,
			fmt.Sprintf("%.2f", asset.Price),
			fmt.Sprintf("%s%.2f", priceChangeSign, asset.PriceChange),
			fmt.Sprintf("%.2f%%", asset.PriceChangePct),
			fmt.Sprintf("%.0f", asset.MarketCap),
			fmt.Sprintf("%.0f", asset.Volume),
		}
		table.Append(row)
	}

	// Imprimir tabla
	table.Render()
}
