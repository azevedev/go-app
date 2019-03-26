package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func newBullet(renderer *sdl.Renderer) *entity {
	bul := &entity{}
	sr := newSpriteRenderer(bul, renderer, "res/bu1llet.bmp")
	bul.addComponent(sr)

	bm := newBulletMove(bul, 0.2)
	bul.addComponent(bm)
	bul.active = false
	return bul
}

var bullets []*entity

func InitBullet(renderer *sdl.Renderer) {
	for i := 0; i < 20; i++ {
		bul := newBullet(renderer)
		bullets = append(bullets, bul)
	}
}

func getBullet() (*entity, bool) {
	for _, bul := range bullets {
		if !bul.active {
			return bul, true
		}
	}
	return nil, false
}
