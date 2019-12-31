package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zerofelx/GameC/GameInfo"
	"github.com/zerofelx/GameC/entities"
)

var GI = GameInfo.GInfo
var Screen = GameInfo.Screen
var Player = new(entities.Player)

func main() {
	Player.Position.X = 80
	Player.Position.Y = 80
	Player.Life = 10
	Player.Speed = 5

	S := Player.Speed
	PlayerX := Player.Position.X
	PlayerY := Player.Position.Y

	rl.InitWindow(Screen.Width, Screen.Height, GI.Title)

	rl.SetTargetFPS(60)
	camera := entities.Camara(PlayerX, PlayerY)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		camera.Target = rl.NewVector2(float32(PlayerX+20), float32(PlayerY+20))
		rl.BeginMode2D(camera)

		rl.DrawText("Static", 20, 20, 20, rl.White)
		rl.DrawText("Player", int32(PlayerX), int32(PlayerY), 20, rl.Red)

		if rl.IsKeyDown(rl.KeyDown) {
			PlayerY += S
		}
		if rl.IsKeyDown(rl.KeyUp) {
			PlayerY -= S
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			PlayerX -= S
		}
		if rl.IsKeyDown(rl.KeyRight) {
			PlayerX += S
		}

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
