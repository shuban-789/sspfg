package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TileSize     = 16
	ScreenWidth  = 640
	ScreenHeight = 480
)

var (
	BackgroundSpriteSheet *ebiten.Image
	bgTypes             = []string{"bluesky", "orangeksy", "purplesky", "cloudysky"}
	currentBgType   	= bgTypes[0]
)

func getBackgroundTile(row, col int) *ebiten.Image {
	x := col * TileSize
	y := row * TileSize
	return BackgroundSpriteSheet.SubImage(image.Rect(x, y, x+TileSize, y+TileSize)).(*ebiten.Image)
}

func generateUniformRow(tile [2]int, count int) [][2]int {
	row := make([][2]int, count)
	for i := 0; i < count; i++ {
		row[i] = tile
	}
	return row
}

func repeatRow(tile [2]int, width, times int) [][][2]int {
	rows := make([][][2]int, times)
	repeated := generateUniformRow(tile, width)
	for i := 0; i < times; i++ {
		rows[i] = repeated
	}
	return rows
}

func DrawBackground(screen *ebiten.Image) {
	var backgroundLayout [][][2]int

	initTemplates()

	switch currentBgType {
		case "bluesky":
			backgroundLayout = blueSkyTemplate
		case "orangeksy":
			backgroundLayout = orangeSkyTemplate
		case "purplesky":
			backgroundLayout = purpleSkyTemplate
		case "cloudysky":
			backgroundLayout = cloudySkyTemplate
		default:
			backgroundLayout = blueSkyTemplate
	} 

	for y, row := range backgroundLayout {
		for x, tileID := range row {
			tile := getBackgroundTile(tileID[0], tileID[1])
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
			screen.DrawImage(tile, op)
		}
	}
}
