package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	SCREEN_HEIGHT = 400
	SCREEN_WIDTH  = 400
	FOOD_SIZE     = 5
	SCORE_INC     = 100
)

var (
	snk snakeSlice
	f   = food{
		centerX: SCREEN_WIDTH/2 + 50,
		centerY: SCREEN_HEIGHT/2 + 50,
		radius:  FOOD_SIZE,
		col:     rl.Gold,
	}
	fd    Food
	score = 0
)

func main() {

	snk.items = append(snk.items, snake{
		posX:   SCREEN_WIDTH / 2,
		posY:   SCREEN_HEIGHT / 2,
		width:  20,
		height: 20,
		col:    rl.Red,
	})

	fd = &f

	rl.InitWindow(SCREEN_WIDTH, SCREEN_HEIGHT, "snake")
	rl.SetConfigFlags(rl.FlagMsaa4xHint)

	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		rl.ClearBackground(rl.Black)

		snk.moveSnake()

		snakeHead := snk.items[0]

		collision := fd.foodCollision(rl.Rectangle{
			X:      snakeHead.posX,
			Y:      snakeHead.posY,
			Width:  float32(snakeHead.width),
			Height: float32(snakeHead.height),
		})

		if collision {
			fd.respawnFood()
			snk.snakeGrow(snakeHead)
			score += SCORE_INC
		}

		rl.BeginDrawing()
		{
			snk.drawSnake()

			fd.spawnFood()
			scoreText := fmt.Sprintf("Score: %d", score)
			rl.DrawText(scoreText, 10, 10, 15, rl.RayWhite)

			if snk.outOfZone() {
				snk.items = []snake{}
				snk.items = append(snk.items, snake{
					posX:   SCREEN_WIDTH / 2,
					posY:   SCREEN_HEIGHT / 2,
					width:  20,
					height: 20,
					col:    rl.Red,
				})
				score = 0
			}

		}
		rl.EndDrawing()
	}
}
