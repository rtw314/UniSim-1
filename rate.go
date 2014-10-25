package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type RateLimiter struct {
	fps      uint32
	lastTime uint32
	count    uint32
}

func NewRateLimiter(fps uint32) *RateLimiter {
	return &RateLimiter{fps, 0, 0}
}

func (d *RateLimiter) BeginFrame() uint32 {
	ret := sdl.GetTicks() - d.lastTime
	if ret <= 0 {
		return 1
	}
	return ret
}

func (d *RateLimiter) EndFrame() {
	curTime := sdl.GetTicks()
	dt := curTime - d.lastTime
	d.lastTime = curTime

	delay := int(1000/30) - int(dt)
	if delay < 0 {
		return
	}
	sdl.Delay(uint32(delay))
}

func (d *RateLimiter) FPS(rate uint32) uint32 {
	d.count += 1
	if d.count%100 == 0 {
		fps := 1000/10/d.BeginFrame() + 9*d.fps/10
		return fps
	}
	return 0
}
