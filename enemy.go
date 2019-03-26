package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	eSpeed = 0.5
)

func newEnemy(renderer *sdl.Renderer, x float64, y float64, r float64) *entity {
	enemy := &entity{}

	enemy.position.x = x
	enemy.position.y = y
	enemy.rotation = r
	enemy.active = true

	sr := newSpriteRenderer(enemy, renderer, "res/enemy.bmp")
	enemy.addComponent(sr)

	return enemy
}
