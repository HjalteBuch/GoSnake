package main

import (
	"Snake/snake"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)


const velocity = 20
const w = 640
const h = 480

var window *sdl.Window
var surface *sdl.Surface
var pixel uint32

func main() {
    if !sdlInit() {
        println("Something went wrong when initalizing sdl")
        closeSdl()
        return
    }

    println("Creating background on surface")
    surface.FillRect(nil, 0)


    println("Creating object")
    head := sdl.Rect{w/2-velocity/2, h/2-velocity/2, velocity, velocity}
    colour := sdl.Color{R: 200, G: 200, B: 200, A: 200} // purple pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
    pixel = sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
    println("Adding object to surface")

    snake := snake.Snake{[]sdl.Rect{head}, 0}
	
    surface.FillRect(nil, 0)
    for _, body := range(snake.Body) {
        surface.FillRect(&body, pixel)
    }

    println("Initial updating of surface")
    window.UpdateSurface()

    running := true
    for running {
        for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
            switch et := event.(type) {
            case *sdl.QuitEvent:
                println("Quit event")
                running = false
                break
            case *sdl.KeyboardEvent:
                snake.ChangeDirection(et.Keysym.Sym)
            }
        }
        snake.Move(velocity)
        draw(snake)
        window.UpdateSurface()
        time.Sleep(time.Second/2)
    }

    closeSdl()
}

func draw(snake snake.Snake) {
    surface.FillRect(nil, 0)
    for _, body := range(snake.Body) {
        surface.FillRect(&body, pixel)
    }
}

func sdlInit() bool {
    success := true
    println("Initializing SDL...")
    err := sdl.Init(sdl.INIT_EVERYTHING)
    if err != nil {
        println(err)
        success = false
    }

    println("Initializing window...")
    window, err = sdl.CreateWindow("SDL Window", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, w, h, sdl.WINDOW_SHOWN)
    if err != nil {
        println(err)
        success = false
    }

    println("Initializing surface...")
    surface, err = window.GetSurface()
    if err != nil {
        println(err)
        success = false
    }

    return success
}

func closeSdl() {
    println("Disposing off sdl...")
    window.Destroy()
    sdl.Quit()
}
