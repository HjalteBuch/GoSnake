package snake

import (
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

func (s *Snake) Move(velocity int32) {
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

func (s *Snake) ChangeDirection(key sdl.Keycode) {
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
