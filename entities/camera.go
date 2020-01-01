package entities

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zerofelx/GameC/GameInfo"
)

var Screen = GameInfo.Screen

func Camara(PlayerX float32, PlayerY float32) rl.Camera2D {
	camera := rl.Camera2D{}
	camera.Target = rl.NewVector2(float32(PlayerX), float32(PlayerY))
	camera.Offset = rl.NewVector2(float32(Screen.Width/2), float32(Screen.Height/2))
	camera.Rotation = 0.0
	camera.Zoom = 1.0

	return camera
}
