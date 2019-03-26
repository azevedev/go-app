package main

import (
	"fmt"
	"reflect"

	"github.com/veandco/go-sdl2/sdl"
)

type pos struct {
	x, y float64
}

type component interface {
	onUpdate() error
	onDraw(renderer *sdl.Renderer) error
	onCollision() error
	onDestroy() error
}

type entity struct {
	position   pos
	rotation   float64
	active     bool
	components []component
}

func (ent *entity) draw(renderer *sdl.Renderer) error {
	for _, comp := range ent.components {
		if err := comp.onDraw(renderer); err != nil {
			return err
		}
	}
	return nil
}

func (ent *entity) update() error {
	for _, comp := range ent.components {
		if err := comp.onUpdate(); err != nil {
			return err
		}
	}
	return nil
}

func (ent *entity) addComponent(new component) {
	for _, c := range ent.components {
		if reflect.TypeOf(new) == reflect.TypeOf(c) {
			panic(fmt.Sprintf("Erro em addComponent!"))
		}
	}
	ent.components = append(ent.components, new)
}

func (ent *entity) getComponent(c component) component {
	typeC := reflect.TypeOf(c)
	for _, comp := range ent.components {
		if reflect.TypeOf(comp) == typeC {
			return c
		}
	}
	panic(fmt.Sprintf("Erro em getComponent!"))
}

var enteties []*entity
