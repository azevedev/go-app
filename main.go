package main

import (
	//"time"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println("Erro em sdl.INIT: ", err)
		return
	}

	w, err := sdl.CreateWindow("GAME", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_OPENGL)
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

	player1, err := newPlayer(renderer)
	if err != nil {
		fmt.Println("Erro em newPlayer: ", err)
		return
	}

	enemy, err := newEnemy(renderer, 300, 300)
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
		enemy.draw(renderer)
		enemy.update()
		renderer.Present()
	}

	return
}
