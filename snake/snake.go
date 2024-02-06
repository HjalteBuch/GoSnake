package snake

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Snake struct {
    Body []sdl.Rect
    Direction []int
}

const (
    RIGHT = iota
    LEFT
    UP
    DOWN
)

func NewSnake(startingPoint sdl.Point, size int32) Snake{
    head := sdl.Rect{
        X: startingPoint.X,
        Y: startingPoint.Y,
        W: size,
        H: size,
    }
    return Snake{[]sdl.Rect{head}, []int{0}}
}

func (s *Snake) AddPart(startingPoint sdl.Point, size int32) {
    part := sdl.Rect{
        X: startingPoint.X,
        Y: startingPoint.Y,
        W: size,
        H: size,
    }
    s.Body = append(s.Body, part)
}

func (s *Snake) Move(velocity int32) {
    for i := 0; i < len(s.Body); i++ {
        switch s.Direction[i] {
        case RIGHT:
            s.Body[i].X += velocity
        case LEFT:
            s.Body[i].X -= velocity
        case DOWN:
            s.Body[i].Y += velocity
        case UP:
            s.Body[i].Y -= velocity
        }
        println(s.Body[i].X)
    }
}

func (s *Snake) ChangeDirection(key sdl.Keycode) {
    switch key {
    case sdl.K_LEFT:
        s.Direction[0] = LEFT
    case sdl.K_RIGHT:
        s.Direction[0] = RIGHT
    case sdl.K_UP:
        s.Direction[0] = UP
    case sdl.K_DOWN:
        s.Direction[0] = DOWN
    }
}
