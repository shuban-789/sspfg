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
	SpriteYOffset = -10
)

var frameCounts = map[int]int{
	0: 4,
	2: 6,
	3: 6,
	5: 6,
}

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

func (p *Player) checkCollisionX(world [][][2]int) {
	if world == nil || len(world) == 0 || len(world[0]) == 0 {
		return
	}

	left := int((p.X) / TileSize)
	right := int((p.X + FrameWidth - 1) / TileSize)
	top := int(p.Y / TileSize)
	bottom := int((p.Y + FrameHeight - 1) / TileSize)

	worldHeight := len(world)
	worldWidth := len(world[0])

	for y := top; y <= bottom; y++ {
		if y < 0 || y >= worldHeight {
			continue
		}
		for x := left; x <= right; x++ {
			if x < 0 || x >= worldWidth {
				continue
			}
			if isSolidTile(world[y][x]) {
				if p.VX > 0 {
					p.X = float64(x*TileSize - FrameWidth)
				} else if p.VX < 0 {
					p.X = float64((x + 1) * TileSize)
				}
				p.VX = 0
				return
			}
		}
	}
}

func (p *Player) checkCollisionY(world [][][2]int) {
	if world == nil || len(world) == 0 || len(world[0]) == 0 {
		return
	}

	left := int((p.X) / TileSize)
	right := int((p.X + FrameWidth - 1) / TileSize)
	top := int(p.Y / TileSize)
	bottom := int((p.Y + FrameHeight - 1) / TileSize)

	worldHeight := len(world)
	worldWidth := len(world[0])
	p.OnGround = false

	for y := top; y <= bottom; y++ {
		if y < 0 || y >= worldHeight {
			continue
		}
		for x := left; x <= right; x++ {
			if x < 0 || x >= worldWidth {
				continue
			}
			if isSolidTile(world[y][x]) {
				if p.VY > 0 {
					p.Y = float64(y*TileSize - FrameHeight)
					p.OnGround = true
				} else if p.VY < 0 {
					p.Y = float64((y + 1) * TileSize)
				}
				p.VY = 0
				return
			}
		}
	}
}

func (p *Player) Update(world [][][2]int) {
	const gravity = 0.5
	const moveSpeed = 2
	const jumpForce = -10

	p.VY += gravity
	prevRow := p.AnimRow

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

	if p.AnimRow != prevRow {
		p.frameIndex = 0
	}

	if p.OnGround && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.VY = jumpForce
		p.OnGround = false
	}

	p.X += p.VX
	p.checkCollisionX(world)

	p.Y += p.VY
	p.checkCollisionY(world)

	p.tickCount++
	if p.tickCount%5 == 0 {
		if count, ok := frameCounts[p.AnimRow]; ok {
			p.frameIndex = (p.frameIndex + 1) % count
		}
	}

	scaledWidth := FrameWidth * SpriteScale
	scaledHeight := FrameHeight * SpriteScale

	if p.X < float64(0 - TileSize) {
		p.X = float64(0 - TileSize)
	}

	if p.X+scaledWidth > float64(ScreenWidth + TileSize) {
		p.X = float64(ScreenWidth + TileSize) - scaledWidth
	}

	if p.Y < float64(0 - TileSize) {
		p.Y = float64(0 - TileSize)
	}

	if p.Y+scaledHeight > float64(ScreenHeight + TileSize) {
		p.Y = float64(ScreenHeight + TileSize) - scaledHeight
	}
}


func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	frame := getFrame(p.AnimRow, p.frameIndex)

	if p.FacingLeft {
		op.GeoM.Scale(-SpriteScale, SpriteScale)
		op.GeoM.Translate(p.X+FrameWidth*SpriteScale, p.Y+SpriteYOffset)
	} else {
		op.GeoM.Scale(SpriteScale, SpriteScale)
		op.GeoM.Translate(p.X, p.Y+SpriteYOffset)
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
