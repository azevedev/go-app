package main

import (
	//"time"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenW    = 800
	screenH    = 600
	playerSize = 1
)

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
	var enteties []*entity
	for i := 0; i < 8; i++ {
		for j := 0; j < 6; j++ {
			x := (float64(i) / 5) * 600
			x += 42
			y := float64(j * 30)
			y += 100
			enemy := newEnemy(renderer, x, y, 0)
			if err != nil {
				fmt.Println("erro em newEnemy: ", err)
			}

			enteties = append(enteties, enemy)
		}
	}
	//enemy := newEnemy(renderer, 300, 300)
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
		err := player1.draw(renderer)
		if err != nil {
			fmt.Println("Drawning player: ", err)
			return
		}

		err = player1.update()
		if err != nil {
			fmt.Println("Updating player: ", err)
			return
		}

		for _, ent := range enteties {
			ent.draw(renderer)
			ent.update()
		}

		renderer.Present()
	}

	return
}
