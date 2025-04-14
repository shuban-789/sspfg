package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var SpriteSheet *ebiten.Image

type Game struct {
	Player Player
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{135, 206, 235, 255})
	ebitenutil.DrawRect(screen, 0, 432, 640, 48, color.RGBA{34, 139, 34, 255})
	g.Player.Draw(screen)
	ebitenutil.DebugPrint(screen, "Use ← → to move, space to jump")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	var err error
	SpriteSheet, _, err = ebitenutil.NewImageFromFile("../assets/sprites/knight.png")
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Super Simple Platformer Fighting Game")

	game := &Game{
		Player: Player{X: 100, Y: 100},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
