package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	window := sdl.CreateWindow("UniSim", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	surface := window.GetSurface()

	scene := NewScene()
	scene.AddPerson(NewPerson(1, 1))
	scene.AddPerson(NewPerson(20, 1))

	rate := NewRateLimiter(30)

	done := false
	for !done {
		e := sdl.PollEvent()
		switch keyEvent := e.(type) {
		case *sdl.KeyDownEvent:
			if keyEvent.Keysym.Scancode == sdl.SCANCODE_ESCAPE {
				done = true
			}
		}

		dt := rate.BeginFrame()
		scene.Update(dt)

		rect := sdl.Rect{0, 0, 800, 600}
		surface.FillRect(&rect, 0x00000000)

		scene.Draw(surface)
		window.UpdateSurface()

		rate.EndFrame()
		if fps := rate.FPS(100); fps > 0 {
			fmt.Println("FPS =", fps)
		}
	}

	window.Destroy()
}
