package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	eSpeed = 0.5
)

type enemy struct {
	tex  *sdl.Texture
	x, y float64
}

func newEnemy(renderer *sdl.Renderer, x float64, y float64) (e enemy) {
	e.tex = newTex(renderer, "res/enemy.bmp")

	e.x = float64(x)
	e.y = float64(y)

	return e
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	x := e.x - 10
	y := e.y - 10
	renderer.Copy(e.tex, &sdl.Rect{0, 0, 20, 21}, &sdl.Rect{int32(x), int32(y), 20, 21})
}

func (e *enemy) update() {

}
