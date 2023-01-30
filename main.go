package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Circle struct {
	xc, yc int
	r      int
	c      color.Color
}

type game struct {
	c Circle
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error                             { return nil }
func (g *game) Draw(screen *ebiten.Image) {
	DrawCircle(screen, g.c.xc, g.c.yc, g.c.r, g.c.c)
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	g := &game{Circle{screenWidth / 2, screenHeight / 2, 200, color.RGBA{100, 100, 255, 255}}}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func DrawCircle(img *ebiten.Image, xc, yc, r int, c color.Color) {
	img.Set(xc, yc, c)
	img.Set(xc+r, yc, c)
	img.Set(xc-r, yc, c)
	img.Set(xc, yc+r, c)
	img.Set(xc, yc-r, c)
}
