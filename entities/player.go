package entities

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"github.com/byrdbrandon6/astro/assets"
	"github.com/byrdbrandon6/astro/utils"
	"github.com/byrdbrandon6/astro/vars"
)

type Player struct {
	Position utils.Vec2
	rotation float64
	sprite   *ebiten.Image
}

func NewPlayer() *Player {
	sprite := assets.PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := utils.Vec2{
		X: vars.ScreenWidth/2 - halfW,
		Y: vars.ScreenHeight/2 - halfH,
	}

	return &Player{
		Position: pos,
		sprite:   sprite,
	}
}

func (p *Player) Update(mX, mY float64) {
	speed := float64(300 / ebiten.TPS())

	// mX, mY := ebiten.CursorPosition()

	p.rotation = p.getRotation(float64(mX), float64(mY))

	var delta utils.Vec2

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		delta.Y = speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		delta.Y = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		delta.X = -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		delta.X = speed
	}

	// Check for diagonal movement
	if delta.X != 0 && delta.Y != 0 {
		factor := speed / math.Sqrt(delta.X*delta.X+delta.Y*delta.Y)
		delta.X *= factor
		delta.Y *= factor
	}

	p.Position.X += delta.X
	p.Position.Y += delta.Y
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()

	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-halfW, -halfH)

	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.Position.X, p.Position.Y)

	// Debug Stuff for me
	s := fmt.Sprintf(
		"Position: (%f, %f) Rotation: %f",
		p.Position.X,
		p.Position.X,
		p.rotation*180/math.Pi,
	)
	mX, mY := ebiten.CursorPosition()
	vector.StrokeLine(
		screen,
		float32(p.Position.X+halfW),
		float32(p.Position.Y+halfH),
		float32(mX),
		float32(mY),
		1.0,
		color.RGBA{255, 0, 0, 255},
		false,
	)

	ebitenutil.DebugPrintAt(
		screen,
		s,
		int(p.Position.X)-int(halfW),
		int(p.Position.Y)+int(halfH*2)+5,
	)

	screen.DrawImage(p.sprite, op)
}

func (p *Player) getRotation(mX, mY float64) float64 {
	distanceX := mX - p.Position.X - float64(p.sprite.Bounds().Dx()/2)
	distanceY := mY - p.Position.Y - float64(p.sprite.Bounds().Dy()/2)
	degrees := math.Atan2(distanceY, distanceX) * 180 / math.Pi
	degrees += 90

	return degrees * math.Pi / 180
}
