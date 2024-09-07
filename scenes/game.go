package game

import (
	"fmt"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/byrdbrandon6/astro/assets"
	"github.com/byrdbrandon6/astro/entities"
	"github.com/byrdbrandon6/astro/utils"
	"github.com/byrdbrandon6/astro/vars"
)

type Game struct {
	player *entities.Player
	world  *ebiten.Image
	camera utils.Camera
	layers [][]int
}

const worldSizeX = vars.ScreenWidth / 16

func NewGame() *Game {
	return &Game{
		player: entities.NewPlayer(),
		world:  ebiten.NewImage(vars.ScreenWidth, vars.ScreenHeight),
		camera: utils.Camera{ViewPort: utils.Vec2{X: vars.ScreenWidth, Y: vars.ScreenHeight}},
		layers: [][]int{
			{
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243,
				243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 219, 243, 243, 243, 219, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,

				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 218, 243, 243, 243,
				243, 218, 243, 243, 243, 243, 243, 243, 243, 243, 243, 244, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
				243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243, 243,
			},
			{
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 26, 27, 28, 29, 30, 31, 0, 0, 0, 0, 0, 0, 0, 0, 0, 63, 64, 65, 66, 67, 68, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 51, 52, 53, 54, 55, 56, 0, 0, 0, 0, 0, 0, 0, 0, 0, 88, 89, 90, 91, 92, 93, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 76, 77, 78, 79, 80, 81, 0, 0, 0, 0, 0, 0, 0, 0, 0, 113, 114, 115, 116, 117, 118, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 101, 102, 103, 104, 105, 106, 0, 0, 0, 0, 0, 0, 0, 0, 0, 138, 139, 140, 141, 142, 143, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 126, 127, 128, 129, 130, 131, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 288, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 303, 303, 245, 242, 303, 303, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0,

				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 45, 46, 47, 48,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 70, 71, 72, 73,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 95, 96, 97, 98,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 120, 121, 122, 123,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 145, 146, 147, 148,

				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 267, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 268, 270, 242, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 192, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 193, 222, 0, 0, 0, 0, 0,
				0, 0, 0, 0, 0, 0, 0, 245, 242, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
	}
}

func (g *Game) Update() error {
	mouseX, mouseY := g.camera.ScreenToWorld(ebiten.CursorPosition())

	g.player.Update(mouseX, mouseY)
	g.camera.Position.X = g.player.Position.X - vars.ScreenWidth/2
	g.camera.Position.Y = g.player.Position.Y // - vars.ScreenHeight/2

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Clear()
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%worldSizeX)*16), float64((i/worldSizeX)*16))

			sx := (t % 25) * 16
			sy := (t / 25) * 16
			g.world.DrawImage(
				assets.Tiles.SubImage(image.Rect(sx, sy, sx+16, sy+16)).(*ebiten.Image),
				op,
			)
		}
	}

	g.player.Draw(g.world)

	g.camera.Render(g.world, screen)

	worldX, worldY := g.camera.ScreenToWorld(ebiten.CursorPosition())

	// ebitenutil.DebugPrint(
	// 	screen,
	// 	fmt.Sprintf(
	// 		"TPS: %0.2f\nMove (WASD/Arrows)\nZoom (QE)\nRotate (R)\nReset (Space)",
	// 		ebiten.ActualTPS(),
	// 	),
	// )
	//
	ebitenutil.DebugPrintAt(
		screen,
		fmt.Sprintf("%s\nCursor World Pos: %.2f,%.2f",
			g.camera.String(),
			worldX, worldY),
		0, vars.ScreenHeight-32,
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return vars.ScreenWidth, vars.ScreenHeight
}
