package main

import(
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type spriteRenderer struct{
	container *entity
	tex *sdl.Texture

}

func newSpriteRenderer(cont *entity, renderer *sdl.Renderer, filename string) *spriteRenderer{
	return &spriteRenderer{
		container: cont,
		tex: newTex(renderer, filename)}
}
func (sr *spriteRenderer) onDraw(renderer *sdl.Renderer) error{
	_, _, width, height, err := sr.tex.Query()
	if err != nil{
		return err
	}
	x := sr.container.position.x - float64(width/2.0)
	y := sr.container.position.y - float64(height/2.0)
	renderer.CopyEx(
		sr.tex, 
		nil, 
		&sdl.Rect{int32(x), int32(y), width, height}, 
		sr.container.rotation,
		&sdl.Point{width/2, height/2},
		sdl.FLIP_NONE)

	return nil
}

func (sr *spriteRenderer) onUpdate() error{
	return nil
}
func (sr *spriteRenderer) onCollision() error{
	return nil
}
func (sr *spriteRenderer) onDestroy() error{
	return nil
}

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