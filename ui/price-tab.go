package ui

import (
	"bytes"
	"errors"
	"fmt"
	"goldwatcher/assets"
	"goldwatcher/prices"
	"image"
	"image/png"
	"io"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) pricesTab() *fyne.Container {
	chart := app.getChart()
	chartContainer := container.NewVBox(chart)
	app.PriceChartContainer = chartContainer

	return chartContainer
}

func (app *Config) getChart() *canvas.Image {
	apiURL := fmt.Sprintf("https://charts.goldprice.org/charts/gold_3d_b_o_%s_x.png", strings.ToLower(prices.DefaultCurrency))
	var image *canvas.Image

	err := app.downloadFile(apiURL, "gold.png")
	if err != nil {
		image = canvas.NewImageFromResource(assets.ResourceNetworkDisconnectSvgrepoComSvg)
	} else {
		image = canvas.NewImageFromFile("gold.png")
	}

	image.SetMinSize(fyne.Size{
		Width:  770,
		Height: 410,
	})

	image.FillMode = canvas.ImageFillOriginal
	return image
}

func (app *Config) downloadFile(URL, fileName string) error {
	// get the response bytes from calling a url
	response, err := app.HTTPClient.Get(URL)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("received wrong response code when downloading image")
	}

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	out, err := os.Create(fmt.Sprintf("./%s", fileName))
	if err != nil {
		return err
	}

	err = png.Encode(out, img)
	if err != nil {
		return err
	}

	return nil
}
