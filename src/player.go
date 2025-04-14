package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	FrameWidth   = 32
	FrameHeight  = 32
	SpriteScale  = 1.5
	RollDuration = 20
)

type Player struct {
	X, Y        float64
	VX, VY      float64
	OnGround    bool
	frameIndex  int
	tickCount   int
	AnimRow     int
	FacingLeft  bool
	Rolling     bool
	RollTick    int
}

func (p *Player) Update() {
	const gravity = 0.5
	const moveSpeed = 2
	const jumpForce = -10

	p.VY += gravity

	if ebiten.IsKeyPressed(ebiten.KeyR) && !p.Rolling {
		p.Rolling = true
		p.RollTick = 0
		p.AnimRow = 5
	}

	if p.Rolling {
		p.RollTick++
		if p.RollTick > RollDuration {
			p.Rolling = false
		}
		if p.FacingLeft {
			p.VX = -moveSpeed * 2
		} else {
			p.VX = moveSpeed * 2
		}
	} else {
		p.VX = 0

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			p.VX = -moveSpeed
			p.FacingLeft = true
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.VX = moveSpeed
			p.FacingLeft = false
		}

		if p.VX != 0 {
			if p.VX > 0 {
				p.AnimRow = 2
			} else {
				p.AnimRow = 3
			}
		} else {
			p.AnimRow = 0
		}
	}

	if p.OnGround && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.VY = jumpForce
		p.OnGround = false
	}

	p.X += p.VX
	p.Y += p.VY

	if p.Y > 400 {
		p.Y = 400
		p.VY = 0
		p.OnGround = true
	}

	p.tickCount++
	if p.tickCount%5 == 0 {
		switch p.AnimRow {
		case 0:
			p.frameIndex = (p.frameIndex + 1) % 4
		case 2, 3:
			p.frameIndex = (p.frameIndex + 1) % 6
		case 5:
			p.frameIndex = (p.frameIndex + 1) % 6
		}
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	frame := getFrame(p.AnimRow, p.frameIndex)

	if p.FacingLeft {
		op.GeoM.Scale(-SpriteScale, SpriteScale)
		op.GeoM.Translate(p.X+FrameWidth*SpriteScale, p.Y)
	} else {
		op.GeoM.Scale(SpriteScale, SpriteScale)
		op.GeoM.Translate(p.X, p.Y)
	}

	screen.DrawImage(frame, op)
}

func getFrame(row, col int) *ebiten.Image {
	x := col * FrameWidth
	y := row * FrameHeight
	return KnightSpriteSheet.SubImage(
		image.Rect(x, y, x+FrameWidth, y+FrameHeight),
	).(*ebiten.Image)
}
