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
    body := sdl.Rect{
        X: startingPoint.X - size+2,
        Y: startingPoint.Y+2,
        W: size-4,
        H: size-4,
    }
    return Snake{[]sdl.Rect{head, body}, []int{0, 0}}
}

func (s *Snake) AddPart(size int32) {
    x := s.Body[len(s.Body)-1].X
    y := s.Body[len(s.Body)-1].Y
    switch s.Direction[len(s.Direction)-1] {
    case RIGHT:
        x -= size
    case LEFT:
        x += size
    case DOWN:
        y -= size
    case UP:
        y += size
    }
    part := sdl.Rect{
        X: x,
        Y: y,
        W: size-4,
        H: size-4,
    }
    s.Body = append(s.Body, part)
    s.Direction = append(s.Direction, s.Direction[len(s.Direction)-1])
}

func (s *Snake) Collision() bool {
    head := s.Body[0]
    for i := 1; i < len(s.Body); i++ {
        if head.HasIntersection(&s.Body[i]) {
            return true
        }
    }
    return false
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
    }
    var temp []int
    for i := 0; i < len(s.Direction); i++ {
        temp = append(temp, s.Direction[i])
    }
    for i := 1; i < len(s.Direction); i++ {
        s.Direction[i] = temp[i-1]
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
