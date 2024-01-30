package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const w = 640
const h = 480

func main() {
    println("Initializing SDL")
    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        panic(err)
    }
    defer sdl.Quit()

    println("Initializing window")
    window, err := sdl.CreateWindow("SDL Window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)
    if err != nil {
        panic(err)
    }
    defer window.Destroy()

    println("Getting surface")
    surface, err := window.GetSurface()
    if err != nil {
        panic(err)
    }
    println("Creating background on surface")
    surface.FillRect(nil, 250)

    println("Creating object")
    rect := sdl.Rect{w/2-20, h/2-20, 40, 40}
    colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
    pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
    println("Adding object to surface")
    surface.FillRect(&rect, pixel)
	

    println("Initial updating of surface")
    window.UpdateSurface()

    running := true
    for running {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch et := event.(type) {
            case *sdl.QuitEvent:
                println("Quit")
                running = false
                break
            case *sdl.KeyboardEvent:
                if et.Keysym.Sym == sdl.K_LEFT {
                    rect.X += -20
                } else if et.Keysym.Sym == sdl.K_RIGHT{
                    rect.X += 20
                }
                
                surface.FillRect(nil, 250)
                surface.FillRect(&rect, pixel)
                window.UpdateSurface()
            }
        }
    }
}
