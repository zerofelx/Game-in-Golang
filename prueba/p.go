package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position rl.Vector2
}
type PlayerPosition struct {
	X float32
	Y float32
}

var Rectangle = new(Player)
var PlayerP = PlayerPosition{20, 20}

var screenWidth = int32(800)
var screenHeight = int32(450)

func main() {
	Rectangle.Position.X = PlayerP.X
	Rectangle.Position.Y = PlayerP.Y
	PlayerX := Rectangle.Position.X
	PlayerY := Rectangle.Position.Y

	rl.InitWindow(screenWidth, screenHeight, "Prueba")

	rl.SetTargetFPS(60)

	camera := rl.Camera2D{}
	camera.Target = rl.NewVector2(float32(PlayerX+20), float32(PlayerY+20))
	camera.Offset = rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2))
	camera.Rotation = 0.0
	camera.Zoom = 1.0

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		camera.Target = rl.NewVector2(float32(PlayerX+20), float32(PlayerY+20))
		rl.BeginMode2D(camera)

		rl.DrawText("Player", int32(PlayerX), int32(PlayerY), 20, rl.Red)
		rl.DrawText("Static", 20, 20, 20, rl.White)

		if rl.IsKeyDown(rl.KeyDown) {
			PlayerY += 5
		}
		if rl.IsKeyDown(rl.KeyUp) {
			PlayerY -= 5
		}
		if rl.IsKeyDown(rl.KeyLeft) {
			PlayerX -= 5
		}
		if rl.IsKeyDown(rl.KeyRight) {
			PlayerX += 5
		}
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
