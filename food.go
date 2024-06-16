package main

import (
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	FOOD_BUFFER = 100
)

type Food interface {
	spawnFood()
	foodCollision(rect rl.Rectangle) bool
	respawnFood()
}

type food struct {
	centerX int32
	centerY int32
	radius  float32
	col     rl.Color
}

func (f *food) spawnFood() {
	rl.DrawCircle(
		f.centerX,
		f.centerY,
		f.radius,
		f.col,
	)
}

func (f *food) foodCollision(rect rl.Rectangle) bool {
	return rl.CheckCollisionCircleRec(
		rl.Vector2{
			X: float32(f.centerX),
			Y: float32(f.centerY),
		},
		f.radius,
		rect,
	)
}

func (f *food) respawnFood() {
	f.centerX = rand.Int31n(SCREEN_WIDTH-FOOD_BUFFER) + FOOD_BUFFER
	f.centerY = rand.Int31n(SCREEN_HEIGHT-FOOD_BUFFER) + FOOD_BUFFER
}
