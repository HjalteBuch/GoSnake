package main

import (
	"Snake/snake"
    "math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

const gridSize = 20
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
    snake := snake.NewSnake(sdl.Point{w/2, h/2}, gridSize)
    snake.AddPart(gridSize)
    snake.AddPart(gridSize)
    snake.AddPart(gridSize)
    snake.AddPart(gridSize)
    snake.AddPart(gridSize)

    println("Creating food")
    food := spawnFood()

    running := true
    for running {
        if snake.Body[0].HasIntersection(&food) {
            snake.AddPart(gridSize)
            food = spawnFood()
        }
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
        snake.Move(gridSize)
        if collision(snake) {
            death()
            break
        }

        clearScreen()

        renderer.SetDrawColor(135, 135, 135, 1)
        for i := 0; i < len(snake.Body); i++ {
            renderer.FillRect(&snake.Body[i])
        }
        renderer.SetDrawColor(200, 30, 80, 1)
        renderer.FillRect(&food)

        renderer.Present()

        sdl.Delay(1000 / 5)
    }

    closeSdl()
}

func spawnFood() sdl.Rect{
    x := rand.Int31n(w/20) * 20
    y := rand.Int31n(h/20) * 20
    food := sdl.Rect{
        X: x+2,
        Y: y+2,
        W: gridSize-4,
        H: gridSize-4,
    }
    return food
}

func collision(snake snake.Snake) bool {
    if snake.Body[0].X < 0 || snake.Body[0].X > w || snake.Body[0].Y < 0 || snake.Body[0].Y > h {
        return true
    }
    if snake.Collision() {
        return true
    }
    return false
}

func death(){
    println("Death")
    renderer.SetDrawColor(255, 0, 0, 1)
    renderer.Clear()
    renderer.Present()
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
