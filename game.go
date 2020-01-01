package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zerofelx/GameC/GUI"
	"github.com/zerofelx/GameC/GameInfo"
	"github.com/zerofelx/GameC/Sounds"
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
	rl.InitAudioDevice()

	music := rl.LoadMusicStream("assets/audio/Kindred.mp3")
	pause := false
	rl.PlayMusicStream(music)

	rl.SetTargetFPS(60)
	camera := entities.Camara(PlayerX, PlayerY)

	for !rl.WindowShouldClose() {
		MusicTimePlayed := Sounds.BackgroundMusic(music, pause)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		camera.Target = rl.NewVector2(float32(PlayerX+20), float32(PlayerY+20))
		rl.BeginMode2D(camera)

		rl.DrawText("Static", 20, 20, 20, rl.White)
		rl.DrawText("Player", int32(PlayerX), int32(PlayerY), 20, rl.Red)

		Life := GUI.LifeGUI(PlayerX, PlayerY, Player.Life)
		rl.DrawRectangle(Life.X, Life.Y, Life.Size, 15, rl.Maroon)
		rl.DrawRectangleLines(Life.X, Life.Y, Life.Size, 15, rl.White)

		rl.DrawRectangle(int32(PlayerX-float32(Screen.Width/3)), int32(PlayerY+50), int32(MusicTimePlayed), 12, rl.Maroon)
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
		if rl.IsKeyPressed(rl.KeyK) {
			Player.Life -= 1
		}

		rl.EndDrawing()
	}

	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
	rl.CloseWindow()
}
