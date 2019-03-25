package main

import(
	"fmt"
	"reflect"
	"github.com/veandco/go-sdl2/sdl"
)
type position struct{
	x,t float64
}

type component interface{
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
	onCollision() error
	onDestroy() error
}

type entity struct{
	pos position
	rotation float64
	active bool
	components []component
}

func (ent *entity) addComponent(new component){
	for _, e := range ent.components{
		if reflect.TypeOf(new) == reflect.TypeOf(e){
			panic(fmt.Sprintf("Erro em addComponent!"))
		}
	}
}
