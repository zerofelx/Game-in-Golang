package GUI

import (
	"github.com/zerofelx/GameC/GameInfo"
)

var Screen = GameInfo.Screen

type GElements struct {
	X    int32
	Y    int32
	Size int32
}

func LifeGUI(PlayerX float32, PlayerY float32, vida float32) *GElements {
	Life := new(GElements)
	Life.X = int32(PlayerX - (float32(Screen.Width) / 2.3))
	Life.Y = int32(PlayerY - 200)
	Life.Size = int32(vida * 20)

	return Life
}
