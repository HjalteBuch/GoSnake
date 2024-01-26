package main 

import (
    "fmt"
    "github.com/veandco/go-sdl2/sdl"
)

func main() {
    fmt.Println("Starting snake")
    if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
        panic(err)
    }
    defer sdl.Quit()

    window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

    surface, err := window.GetSurface()
    if err != nil {
        panic(err)
    }
    surface.FillRect(nil, 0)

    rect := sdl.Rect{0, 0, 200, 200}
    colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
    pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
    surface.FillRect(&rect, pixel)
    window.UpdateSurface()

    running := true
    for running {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch event.(type) {
            case *sdl.QuitEvent:
                println("Quit")
                running = false
                break
            }
        }
    }
}
