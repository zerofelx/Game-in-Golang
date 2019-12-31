package ge

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/zerofelx/GameC/GameInfo"
	"github.com/zerofelx/GameC/entities"
)

var Screen = GameInfo.Screen

type Game struct {
	gameOver bool
	pause    bool
	player   entities.Player
}

func NewGame() (g Game) {
	g.Init()
	return
}

func (g *Game) Init() {
	// Inicializar jugador
	g.player.Position = rl.Vector2{float32(Screen.Width / 2), float32(Screen.Height * 7 / 8)}
	g.player.Size = rl.Vector2{float32(Screen.Width / 10), 20}
	g.player.Life = 10

}

func (g *Game) Update() {
	if !g.gameOver {
		if rl.IsKeyPressed(rl.KeyP) {
			g.pause = !g.pause
		}

		if !g.pause {
			if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
				fmt.Println("El usuario se desplazará hacia arriba")
				g.player.Position.Y -= 5
			}
			// if g.player.Position.X - g.player.Size.X/2 {

			// }
			if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
				g.player.Position.X += 5
			}
			if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
				g.player.Position.Y += 5
			}
			if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
				g.player.Position.X -= 5
			}
		}
	}
}
