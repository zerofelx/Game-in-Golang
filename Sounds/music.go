package Sounds

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func BackgroundMusic(music rl.Music, pause bool) float32 {
	rl.UpdateMusicStream(music)

	if rl.IsKeyPressed(rl.KeySpace) {
		rl.StopMusicStream(music)
		rl.PlayMusicStream(music)
	}

	if rl.IsKeyPressed(rl.KeyP) {
		pause = !pause

		if pause {
			rl.PauseMusicStream(music)
		} else {
			rl.ResumeMusicStream(music)
		}
	}

	timePlayed := rl.GetMusicTimePlayed(music) / rl.GetMusicTimeLength(music) * 100 * 4
	return timePlayed
}
