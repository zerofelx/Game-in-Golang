package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
	rl.InitWindow(800, 450, "Game")
	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Juego Creado!", 20, 200, 20, rl.Red)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
