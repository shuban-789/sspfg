package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	worlds = []string{"default", "desert", "prairie", "ice"}
	currentWorld = worlds[0]
)

func DrawWorld(screen *ebiten.Image) {
	var worldLayout [][][2]int

	initWorldTemplates()

	switch currentWorld {
		case "default":
			worldLayout = defaultWorldTemplate
		case "desert":
			worldLayout = desertWorldTemplate
		case "prairie":
			worldLayout = prairieWorldTemplate
		case "ice":
			worldLayout = iceWorldTemplate
		default:
			worldLayout = defaultWorldTemplate
	} 

	for y, row := range worldLayout {
		for x, tileID := range row {
			tile := getBackgroundTile(tileID[0], tileID[1])
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
			screen.DrawImage(tile, op)
		}
	}
}
