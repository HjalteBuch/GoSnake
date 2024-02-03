package main

import (
	"Snake/snake"

	"github.com/veandco/go-sdl2/sdl"
)


const velocity = 20
const w = 640
const h = 480

var window *sdl.Window
var renderer *sdl.Renderer

func main() {
    if !sdlInit() {
        println("Something went wrong when initalizing sdl")
        closeSdl()
        return
    }

    println("Creating snake")
    snakePart := sdl.Rect{w/2-velocity/2, h/2-velocity/2, velocity, velocity}
    snake := snake.Snake{[]sdl.Rect{snakePart}, 0}

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

        clearScreen()

        renderer.SetDrawColor(135, 135, 135, 1)
        renderer.FillRect(&snake.Body[0])

        renderer.Present()

        sdl.Delay(1000 / 5)
    }

    closeSdl()
}

func clearScreen() {
        renderer.SetDrawColor(0, 0, 0, 1)
        renderer.Clear()
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

    println("Initializing renderer...")
    renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
        println(err)
        success = false
    }

    return success
}

func closeSdl() {
    println("Disposing off sdl...")
    window.Destroy()
    renderer.Destroy()
    sdl.Quit()
}
