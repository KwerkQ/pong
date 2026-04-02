package main

import (
	//"fmt"

	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	rec   rl.Rectangle
	color rl.Color
}
type Ball struct {
	rec      rl.Rectangle
	color    rl.Color
	velocity rl.Vector2
}
type Game struct {
	player1 Player
	player2 Player
	ball    Ball
	started bool
	score1  int
	score2  int
	winner  int
}

func DrawGame(game *Game) {
	rl.DrawRectangleRec(game.ball.rec, game.ball.color)
	rl.DrawRectangleRec(game.player1.rec, game.player1.color)
	rl.DrawRectangleRec(game.player2.rec, game.player2.color)
}
func UpdateGame(game *Game) {
	var dir int32
	if game.player1.rec.Y != 450 {
		if rl.IsKeyDown(rl.KeyS) {
			game.player1.rec.Y += 5
		}
	}
	if game.player1.rec.Y != 0 {
		if rl.IsKeyDown(rl.KeyW) {
			game.player1.rec.Y -= 5
		}
	}
	if game.player2.rec.Y != 450 {
		if rl.IsKeyDown(rl.KeyDown) {
			game.player2.rec.Y += 5
		}
	}
	if game.player2.rec.Y != 0 {
		if rl.IsKeyDown(rl.KeyUp) {
			game.player2.rec.Y -= 5
		}
	}
	if rl.IsKeyDown(rl.KeySpace) && game.started != true {
		dir = rl.GetRandomValue(0, 1)
		if dir == 0 {
			game.ball.velocity.X = 5
		} else {
			game.ball.velocity.X = -5
		}
		game.ball.velocity.Y = float32(rl.GetRandomValue(-4, 4))
		game.started = true
	}
	if rl.CheckCollisionRecs(game.player1.rec, game.ball.rec) {
		game.ball.velocity.X *= -1
		if rl.IsKeyDown(rl.KeyW) {
			game.ball.velocity.Y += float32(rl.GetRandomValue(3, 5))
		}
		if rl.IsKeyDown(rl.KeyS) {
			game.ball.velocity.Y += float32(rl.GetRandomValue(-3, -5))
		}
	}
	if rl.CheckCollisionRecs(game.player2.rec, game.ball.rec) {
		game.ball.velocity.X *= -1
		if rl.IsKeyDown(rl.KeyUp) {
			game.ball.velocity.Y += float32(rl.GetRandomValue(3, 5))
		}
		if rl.IsKeyDown(rl.KeyDown) {
			game.ball.velocity.Y += float32(rl.GetRandomValue(-3, -5))
		}
	}
	if game.ball.rec.Y <= 0 {
		game.ball.velocity.Y *= -1
	}
	if game.ball.rec.Y >= 580 {
		game.ball.velocity.Y *= -1
	}
	if game.started {
		if game.ball.velocity.X > 0 {
			game.ball.velocity.X += 0.01
		}
		if game.ball.velocity.X < 0 {
			game.ball.velocity.X -= 0.01
		}
	}
	if game.ball.rec.X <= -20 {
		game.score2 += 1
		game.started = false
		game.ball.velocity.X = 0
		game.ball.velocity.Y = 0
		game.ball.rec.X = 390
		game.ball.rec.Y = 290
	}
	if game.ball.rec.X >= 800 {
		game.score1 += 1
		game.started = false
		game.ball.velocity.X = 0
		game.ball.velocity.Y = 0
		game.ball.rec.X = 390
		game.ball.rec.Y = 290
	}
	game.ball.rec.X += game.ball.velocity.X
	game.ball.rec.Y += game.ball.velocity.Y
	if game.score1 == 10 {
		game.winner = 1
	}
	if game.score2 == 10 {
		game.winner = 2
	}
}

func main() {
	velocity := rl.Vector2{
		X: 0,
		Y: 0,
	}
	player1 := Player{
		rec:   rl.NewRectangle(10, 225, 10, 150),
		color: rl.Beige,
	}
	player2 := Player{
		rec:   rl.NewRectangle(780, 225, 10, 150),
		color: rl.Beige,
	}
	ball := Ball{
		rec:      rl.NewRectangle(390, 290, 20, 20),
		color:    rl.Red,
		velocity: velocity,
	}
	game := Game{
		player1: player1,
		player2: player2,
		ball:    ball,
		started: false,
		score1:  0,
		score2:  0,
		winner:  0,
	}
	rl.InitWindow(800, 600, "ping pong pang")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		if game.winner == 0 {
			rl.DrawText("PING PONG PANG", 225, 200, 40, rl.LightGray)
			rl.DrawText(strconv.Itoa(game.score1), 100, 20, 50, rl.Gold)
			rl.DrawText(strconv.Itoa(game.score2), 675, 20, 50, rl.Gold)
			DrawGame(&game)
			UpdateGame(&game)
		}
		if game.winner == 1 {
			rl.DrawText("Player 1 won\nPress B to play again", 225, 200, 40, rl.LightGray)

		}
		if game.winner == 2 {
			rl.DrawText("Player 2 won\nPress B to play again", 225, 200, 40, rl.LightGray)
		}
		if game.winner != 0 {
			if rl.IsKeyDown(rl.KeyB) {
				game.started = false
				game.ball.velocity.X = 0
				game.ball.velocity.Y = 0
				game.ball.rec.X = 390
				game.ball.rec.Y = 290
				game.score1 = 0
				game.score2 = 0
				game.winner = 0
				game.player1.rec.Y = 225
				game.player2.rec.Y = 225
			}
		}
		rl.EndDrawing()
	}
}
