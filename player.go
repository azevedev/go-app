package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gravity = 0.2
	pSpeed  = 0.5
)

type player struct {
	tex         *sdl.Texture
	x, y        float64
	h, w, frame int32
}

func newPlayer(renderer *sdl.Renderer) *entity{
	player := &entity{}
	player.position.x = screenW/2.0
	player.position.y = screenH/2.0
	player.active = true
	
	
}

func newPlayer(renderer *sdl.Renderer) (p player) {
	p.tex = newTex(renderer, "res/slime.bmp")

	p.w = 32
	p.h = 40
	p.frame = 1
	p.x = 800 / 2
	p.y = 600 - 100

	return p
}

func (p *player) draw(renderer *sdl.Renderer) {
	x := p.x - 16
	y := p.y - 20
	renderer.Copy(p.tex, &sdl.Rect{0 + ((p.frame - 1) * p.w), 0 + ((p.frame - 1) * p.h), p.w, p.h}, &sdl.Rect{int32(x), int32(y), 2 * p.w, 2 * p.h})
}

func (p *player) update() {

	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		p.x -= pSpeed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		p.x += pSpeed
	} else if keys[sdl.SCANCODE_UP] == 1 {
		p.y -= pSpeed
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		p.y += pSpeed
	}

}
