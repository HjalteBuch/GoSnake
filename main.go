package main

import (
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type Snake struct {
    Body []sdl.Rect
    Direction int
}

const (
    RIGHT = iota
    LEFT
    UP
    DOWN
)
const velocity = 20
const w = 640
const h = 480

var window *sdl.Window
var surface *sdl.Surface
var snake Snake
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

    snake = Snake{[]sdl.Rect{head}, 0}
	
    drawSnake()

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
                changeSnakeDirection(et.Keysym.Sym)
            }
        }
        moveSnake()
        drawSnake()
        window.UpdateSurface()
        time.Sleep(time.Second/2)
    }

    closeSdl()
}

func moveSnake() {
    switch snake.Direction {
    case RIGHT:
        snake.Body[0].X += velocity
    case LEFT:
        snake.Body[0].X -= velocity
    case DOWN:
        snake.Body[0].Y += velocity
    case UP:
        snake.Body[0].Y -= velocity
    }
}

func changeSnakeDirection(key sdl.Keycode) {
    switch key {
    case sdl.K_LEFT:
        snake.Direction = LEFT
    case sdl.K_RIGHT:
        snake.Direction = RIGHT
    case sdl.K_UP:
        snake.Direction = UP
    case sdl.K_DOWN:
        snake.Direction = DOWN
    }
}

func drawSnake() {
    surface.FillRect(nil, 0)
    for i := 0; i < len(snake.Body); i++ {
        surface.FillRect(&snake.Body[i], pixel)
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
