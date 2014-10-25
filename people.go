package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Person struct {
	locX, locY float32
}

func NewPerson(x, y float32) *Person {
	return &Person{x, y}
}

func (p *Person) Update(dt uint32) {
	p.locX += .01
	p.locY += .01
}

func (p *Person) Draw(surface *sdl.Surface) {
	rect := sdl.Rect{int32(p.locX * 10), int32(p.locY * 10), 10, 10}
	surface.FillRect(&rect, 0xffff0000)
}
