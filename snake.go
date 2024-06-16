package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Direction string

const (
	UP    Direction = "up"
	DOWN  Direction = "down"
	LEFT  Direction = "left"
	RIGHT Direction = "right"
)

type Snake interface {
	drawSnake()
	moveSnake()
	snakeGrow()
	outOfZone() bool
}

type snake struct {
	posX   float32
	posY   float32
	width  int32
	height int32
	col    rl.Color
}

type snakeSlice struct {
	items     []snake
	direction Direction
}

const (
	SPEED           = 200
	FOLLOW_DISTANCE = 10
)

func (s *snakeSlice) drawSnake() {
	for i := 0; i < len(s.items); i++ {
		rl.DrawRectangle(int32(s.items[i].posX), int32(s.items[i].posY), s.items[i].width, s.items[i].height, s.items[i].col)
	}
}

func (s *snakeSlice) outOfZone() bool {
	isOutOfZone := false

	head := s.items[0]

	if head.posX < 0 {
		isOutOfZone = true
	} else if head.posX > SCREEN_WIDTH {
		isOutOfZone = true
	} else if head.posY < 0 {
		isOutOfZone = true
	} else if head.posY > SCREEN_HEIGHT {
		isOutOfZone = true
	}

	return isOutOfZone
}

func (s *snakeSlice) moveSnake() {
	d := rl.GetFrameTime()

	if rl.IsKeyPressed(rl.KeyW) {
		s.direction = UP
	} else if rl.IsKeyPressed(rl.KeyS) {
		s.direction = DOWN
	} else if rl.IsKeyPressed(rl.KeyD) {
		s.direction = RIGHT
	} else if rl.IsKeyPressed(rl.KeyA) {
		s.direction = LEFT
	}

	switch s.direction {
	case UP:
		s.items[0].posY -= SPEED * d
	case RIGHT:
		s.items[0].posX += SPEED * d
	case LEFT:
		s.items[0].posX -= SPEED * d
	default:
		s.items[0].posY += SPEED * d
	}

	for i := 1; i < len(s.items); i++ {
		child := s.items[i]
		parent := s.items[i-1]
		dx := parent.posX - child.posX
		dy := parent.posY - child.posY
		distance := float32(rl.Vector2Length(rl.NewVector2(dx, dy)))

		if distance > FOLLOW_DISTANCE {
			moveX := (dx / distance) * (distance - FOLLOW_DISTANCE)
			moveY := (dy / distance) * (distance - FOLLOW_DISTANCE)

			s.items[i].posY += moveY
			s.items[i].posX += moveX
		}
	}
}

func (s *snakeSlice) snakeGrow(snkHead snake) {
	tail := s.items[len(s.items)-1]
	s.items = append(s.items, snake{
		posX:   tail.posX,
		posY:   tail.posY,
		width:  tail.width,
		height: tail.width,
		col:    tail.col,
	})
}
