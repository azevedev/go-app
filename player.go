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
	
	sr := newSpriteRenderer(player, renderer, "res/slime.bmp")
	player.addComponent(sr)

	plat := newPlatform(player, pSpeed)
	player.addComponent(plat)

	return player

}

// func newPlayer(renderer *sdl.Renderer) (p player) {
// 	p.tex = newTex(renderer, "res/slime.bmp")

// 	p.w = 32
// 	p.h = 40
// 	p.frame = 1
// 	p.x = 800 / 2
// 	p.y = 600 - 100

// 	return p
// }

func (p *player) draw(renderer *sdl.Renderer) {
	}

func (p *player) update() {


}
