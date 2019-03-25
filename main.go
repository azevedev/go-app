package main

import (
	//"time"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const(
	screenW = 800
	screenH = 600
	playerSize = 1
)

func newTex(renderer *sdl.Renderer, file string) *sdl.Texture {
	img, err := sdl.LoadBMP(file)
	if err != nil {
		panic(fmt.Errorf("Loading file BMP %v: %v", file, err))
	}

	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		panic(fmt.Errorf("Creating texture %v: %v", file, err))
	}

	img.Free()

	return tex
}
func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("Erro em sdl.INIT: ", err)
		return
	}

	w, err := sdl.CreateWindow("GAME", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenW, screenH, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("Erro em CreateWindow: ", err)
		return
	}
	defer w.Destroy()

	renderer, err := sdl.CreateRenderer(w, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("Erro em CreateRenderer: ", err)
		return
	}

	player1 := newPlayer(renderer)
	if err != nil {
		fmt.Println("Erro em newPlayer: ", err)
		return
	}
	var enemies []enemy
	for i := 0; i < 8; i++ {
		for j := 0; j < 6; j++ {
			x := (float64(i) / 5) * 600
			x += 42
			y := float64(j * 30)
			y += 100
			enemy := newEnemy(renderer, x, y)
			if err != nil {
				fmt.Println("erro em newEnemy: ", err)
			}

			enemies = append(enemies, enemy)
		}
	}
	//enemy, err := newEnemy(renderer, 300, 300)
	if err != nil {
		fmt.Println("Erro em newEnemy: ", err)
		return
	}

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}

		}

		renderer.Clear()
		renderer.SetDrawColor(255, 255, 255, 255)
		player1.draw(renderer)
		player1.update()
		for _, enemy := range enemies {
			enemy.draw(renderer)
			enemy.update()
		}

		renderer.Present()
	}

	return
}
