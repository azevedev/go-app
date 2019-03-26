package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type bullet_move struct {
	container *entity
	speed     float64
}

func newBulletMove(cont *entity, sp float64) *bullet_move {
	return &bullet_move{
		container: cont,
		speed:     sp,
	}
}

func (bm *bullet_move) onDraw(*sdl.Renderer) error {
	return nil
}

func (bm *bullet_move) onUpdate() error {
	cont := bm.container.position
	cont.x += bm.speed * math.Cos(bm.container.rotation)
	cont.y += bm.speed * math.Sin(bm.container.rotation)

	if cont.x > screenW || cont.x < 0 || cont.y > screenH || cont.y < 0 {
		bm.container.active = false
	}

	return nil
}

func (bm *bullet_move) onCollision() error {
	return nil
}

func (bm *bullet_move) onDestroy() error {
	return nil
}
