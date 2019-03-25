package main

import(
	"fmt"
	"time"
	"github.com/veandco/go-sdl2/sdl"
)


type shooter struct{
	container *entity
	cooldown time.Duration
	lastShot time.Time

}


func newShooter(cont *entity, cd time.Duration) *shooter{
	return &shooter{
		container: cont,
		cooldown: cd}
}

func (s *shooter) onDraw(*sdl.Renderer) error{
	return nil
}

func (s *shooter) onUpdate() error{
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_SPACE] == 1 {
		if time.Since(s.lastShot) >= s.cooldown{
			s.shoot(s.container.position.x+25, s.container.position.y-20)
			s.shoot(s.container.position.x-25, s.container.position.y-20)
		}
	} 


	return nil
}

func (s *shooter) onCollision() error{
	return nil
}

func (s *shooter) onDestroy() error{
	return nil
}

func (s *shooter) shoot(x float64, y float64){
	fmt.Println("Shooting!")
	s.lastShot = time.Now()
	
}
