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
	
    snake.draw()

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
                snake.changeDirection(et.Keysym.Sym)
            }
        }
        snake.move()
        snake.draw()
        window.UpdateSurface()
        time.Sleep(time.Second/2)
    }

    closeSdl()
}

func (s *Snake) move() {
    switch s.Direction {
    case RIGHT:
        s.Body[0].X += velocity
    case LEFT:
        s.Body[0].X -= velocity
    case DOWN:
        s.Body[0].Y += velocity
    case UP:
        s.Body[0].Y -= velocity
    }
}

func (s *Snake) changeDirection(key sdl.Keycode) {
    switch key {
    case sdl.K_LEFT:
        s.Direction = LEFT
    case sdl.K_RIGHT:
        s.Direction = RIGHT
    case sdl.K_UP:
        s.Direction = UP
    case sdl.K_DOWN:
        s.Direction = DOWN
    }
}

func (s *Snake) draw() {
    surface.FillRect(nil, 0)
    for _, body := range(s.Body) {
        println(s.Direction)
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
