package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position rl.Vector2
	Size     rl.Vector2
	Life     float32
	Speed    float32
}
