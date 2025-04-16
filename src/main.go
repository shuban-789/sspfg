package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/go-mp3"
)

const SampleRate = 44100

var (
	audioContext      *audio.Context
	bgMusicPlayer     *audio.Player
	KnightSpriteSheet *ebiten.Image
)

type Game struct {
	Player Player
}

func (g *Game) Update() error {
	g.Player.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawBackground(screen)
	DrawWorld(screen)
	g.Player.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func initAudio() {
	audioContext = audio.NewContext(SampleRate)

	f, err := os.Open("../assets/music/time_for_adventure.mp3")
	if err != nil {
		log.Fatal(err)
	}

	decoded, err := mp3.NewDecoder(f)
	if err != nil {
		log.Fatal(err)
	}

	stream := audio.NewInfiniteLoop(decoded, decoded.Length())

	bgMusicPlayer, err = audio.NewPlayer(audioContext, stream)
	if err != nil {
		log.Fatal(err)
	}

	bgMusicPlayer.SetVolume(0.5)
	bgMusicPlayer.Play()
}

func main() {
	var err error
	KnightSpriteSheet, _, err = ebitenutil.NewImageFromFile("../assets/sprites/knight.png")
	BackgroundSpriteSheet, _, err = ebitenutil.NewImageFromFile("../assets/sprites/world_tileset.png")
	if err != nil {
		log.Fatal(err)
	}

	initAudio()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Super Simple Platformer Fighting Game")

	game := &Game{
		Player: Player{X: 100, Y: 100},
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}