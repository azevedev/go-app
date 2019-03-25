package main

import(
	"github.com/veandco/go-sdl2/sdl"
)

type platform struct{
	container *entity 
	speed float64

}

func newPlatform(cont *entity, sp float64) *platform {
	p := &platform{
		container: cont,
		speed: sp}
	return p
}

func(p *platform) onDraw(*sdl.Renderer) error{
	return nil
}

func (p *platform) onUpdate() error {
	cont := p.container
	
	keys := sdl.GetKeyboardState()
	if keys[sdl.SCANCODE_LEFT] == 1 {
		cont.position.x -= p.speed
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		cont.position.x += p.speed
	} else if keys[sdl.SCANCODE_UP] == 1 {
		cont.position.y -= p.speed
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		cont.position.y += p.speed
	}
	return nil
}

func(p *platform) onCollision() error{
	return nil
}


func(p *platform) onDestroy() error{
	return nil
}
