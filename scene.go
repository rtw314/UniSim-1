package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	people []*Person
}

func NewScene() *Scene {
	return &Scene{}
}

func (s *Scene) AddPerson(p *Person) {
	s.people = append(s.people, p)
}

func (s *Scene) Update(dt uint32) {
	for _, person := range s.people {
		person.Update(dt)
	}
}

func (s *Scene) Draw(surface *sdl.Surface) {
	for _, person := range s.people {
		person.Draw(surface)
	}
}
