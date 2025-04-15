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

var BackgroundSpriteSheet *ebiten.Image

func getBackgroundTile(row, col int) *ebiten.Image {
	x := col * TileSize
	y := row * TileSize
	return BackgroundSpriteSheet.SubImage(image.Rect(x, y, x+TileSize, y+TileSize)).(*ebiten.Image)
}

func generateRepeatedRow(tile [2]int, count int) [][2]int {
	row := make([][2]int, count)
	for i := 0; i < count; i++ {
		row[i] = tile
	}
	return row
}

func generateBackgroundLayout() [][][2]int {
	var WidthTiles int
	WidthTiles = ScreenWidth / TileSize
	layout := [][][2]int{
		generateRepeatedRow([2]int{9, 0}, WidthTiles),
		generateRepeatedRow([2]int{9, 0}, WidthTiles),
		generateRepeatedRow([2]int{9, 0}, WidthTiles),
		generateRepeatedRow([2]int{10, 0}, WidthTiles),
		generateRepeatedRow([2]int{11, 0}, WidthTiles),
		generateRepeatedRow([2]int{11, 0}, WidthTiles),
		generateRepeatedRow([2]int{11, 0}, WidthTiles),
		generateRepeatedRow([2]int{11, 0}, WidthTiles),
		generateRepeatedRow([2]int{11, 0}, WidthTiles),
		generateRepeatedRow([2]int{11, 0}, WidthTiles),
		generateRepeatedRow([2]int{11, 0}, WidthTiles),
		generateRepeatedRow([2]int{11, 0}, WidthTiles),
		generateRepeatedRow([2]int{12, 0}, WidthTiles),
		generateRepeatedRow([2]int{13, 0}, WidthTiles),
		generateRepeatedRow([2]int{13, 0}, WidthTiles),
		generateRepeatedRow([2]int{13, 0}, WidthTiles),
		generateRepeatedRow([2]int{13, 0}, WidthTiles),
		generateRepeatedRow([2]int{13, 0}, WidthTiles),
		generateRepeatedRow([2]int{13, 0}, WidthTiles),
		generateRepeatedRow([2]int{13, 0}, WidthTiles),
		generateRepeatedRow([2]int{13, 0}, WidthTiles),
		generateRepeatedRow([2]int{14, 0}, WidthTiles),
		generateRepeatedRow([2]int{15, 0}, WidthTiles),
		generateRepeatedRow([2]int{15, 0}, WidthTiles),
		generateRepeatedRow([2]int{15, 0}, WidthTiles),
		generateRepeatedRow([2]int{15, 0}, WidthTiles),
		generateRepeatedRow([2]int{15, 0}, WidthTiles),
		generateRepeatedRow([2]int{15, 0}, WidthTiles),
		generateRepeatedRow([2]int{15, 0}, WidthTiles),
		generateRepeatedRow([2]int{15, 0}, WidthTiles),
	}
	return layout
}

var backgroundLayout = generateBackgroundLayout()

func DrawBackground(screen *ebiten.Image) {
	for y, row := range backgroundLayout {
		for x, tileID := range row {
			tile := getBackgroundTile(tileID[0], tileID[1])
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
			screen.DrawImage(tile, op)
		}
	}
}
