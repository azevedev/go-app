package main

import(

	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	
)

const(
	eSpeed = 0.5 
)
type enemy struct{
	tex *sdl.Texture
	x, y float64
}

func newEnemy(renderer *sdl.Renderer, x int32, y int32) (e enemy, err error){
	e.x = float64(x)
	e.y = float64(y)
	img, err := sdl.LoadBMP("res/enemy.bmp")
	if err != nil {
		fmt.Println("Failed to load PNG: ", err)
		return enemy{}, fmt.Errorf("Erro em LoadBMP: %v", err)
	}


	e.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Println("Failed to create texture: ", err)
		return enemy{}, fmt.Errorf("Erro em CreateTexture: %v", err)
	}

	img.Free()

	return e, nil
}

func (e *enemy) draw(renderer *sdl.Renderer) {
	x := e.x - 10
	y := e.y - 10
	renderer.Copy(e.tex, &sdl.Rect{0, 0, 20, 21},  &sdl.Rect{int32(x), int32(y), 20,21})
}

func (e *enemy) update() {	
	
	
	
}