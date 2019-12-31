package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zerofelx/GameC/GameInfo"
	"github.com/zerofelx/GameC/ge"
)

var GI = GameInfo.GInfo
var Screen = GameInfo.Screen

func main() {
	rl.InitWindow(Screen.Width, Screen.Height, GI.Title)

	// Iniciar juego
	game := ge.NewGame()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)

		rl.DrawText("Juego Creado!", 20, 200, 20, rl.Red)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}
