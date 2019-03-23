package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type bullet struct {
	tex  *sdl.Texture
	x, y float64
	size int32
}

func newBullet(renderer *sdl.Renderer) (bul bullet) {
	bul.tex = newTex(renderer, "res/bullet.bmp")
	bul.size = 14
	return bul
}

func (b *bullet) draw(renderer *sdl.Renderer) {
	x := b.x - 16
	y := b.y - 20
	renderer.Copy(b.tex, &sdl.Rect{0, 0, b.size, b.size}, &sdl.Rect{int32(x), int32(y), b.size, b.size})
}
