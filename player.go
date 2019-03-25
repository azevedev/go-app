package main

import (
	"time"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	gravity = 0.2
	pSpeed  = 0.5
)

func newPlayer(renderer *sdl.Renderer) *entity{
	
	player := &entity{}
	player.position.x = screenW/2.0
	player.position.y = screenH/2.0
	player.active = true
	
	sr := newSpriteRenderer(player, renderer, "res/slime.bmp")
	player.addComponent(sr)

	plat := newPlatform(player, pSpeed)
	player.addComponent(plat)

	shooter := newShooter(player, (time.Millisecond * 300))
	player.addComponent(shooter)
	return player

}
