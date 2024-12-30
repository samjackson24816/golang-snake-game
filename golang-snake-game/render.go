package main

import (
	"fmt"
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Render(game *Game, app *App) {
	switch game.State {
	case Playing:
		RenderPlaying(game, app)
	case Paused:
		// Render Paused in the middle of the screen
		RenderPaused(game, app)
	case GameOver:
		// Render Game Over in the middle of the screen
		RenderGameOver(game, app)
	case HomeScreen:
		RenderHomeScreen(game, app)
	}

}

func RenderGame(game *Game, app *App) {

	rl.ClearBackground(rl.RayWhite)

	screenPerWorldX := float32(app.ScreenSize.X / game.WorldSize.X)
	screenPerWorldY := float32(app.ScreenSize.Y / game.WorldSize.Y)

	// Render the head

	rl.DrawRectangle(game.Position.X*int32(screenPerWorldX), game.Position.Y*int32(screenPerWorldY), int32(screenPerWorldX), int32(screenPerWorldY), rl.Black)

	// Render the tail
	for i, segment := range game.Tail {
		// Get the direction of the tail segment
		var next Vector2Int
		if i == len(game.Tail)-1 {
			next = game.Position
		} else {
			next = game.Tail[i+1].Position
		}

		// Calculate the direction of the tail segment
		dirNext := next.Sub(segment.Position)

		var prev Vector2Int
		if i == 0 {
			prev = next.Scale(-1)
		} else {
			prev = game.Tail[i-1].Position
		}

		dirPrev := segment.Position.Sub(prev)

		// If the directions aren't equal, the tail segment is a corner
		isCorner := dirNext != dirPrev

		// Render the tail segment, making it a rectangle that is thinner side to side

		if isCorner {

			// Render the corner as two intersecting rectangles

			// Previous segment
			var x, y, w, h int32

			if dirPrev.X == 0 && dirPrev.Y == 1 {
				x = int32(math.Round(float64((float32(segment.Position.X) + 0.25) * screenPerWorldX)))
				y = int32(math.Round(float64(float32(segment.Position.Y) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY * 0.75)))
				w = int32(math.Round(float64(screenPerWorldX / 2)))
			} else if dirPrev.X == 0 && dirPrev.Y == -1 {
				x = int32(math.Round(float64((float32(segment.Position.X) + 0.25) * screenPerWorldX)))
				y = int32(math.Round(float64((float32(segment.Position.Y) + 0.25) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY * 0.75)))
				w = int32(math.Round(float64(screenPerWorldX / 2)))
			} else if dirPrev.X == 1 && dirPrev.Y == 0 {
				x = int32(math.Round(float64(float32(segment.Position.X) * screenPerWorldX)))
				y = int32(math.Round(float64((float32(segment.Position.Y) + 0.25) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY / 2)))
				w = int32(math.Round(float64(screenPerWorldX * 0.75)))
			} else if dirPrev.X == -1 && dirPrev.Y == 0 {
				x = int32(math.Round(float64((float32(segment.Position.X) + 0.25) * screenPerWorldX)))
				y = int32(math.Round(float64((float32(segment.Position.Y) + 0.25) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY / 2)))
				w = int32(math.Round(float64(screenPerWorldX * 0.75)))
			}

			rl.DrawRectangle(x, y, w, h, rl.Black)

			// Next segment
			if dirNext.X == 0 && dirNext.Y == 1 {
				x = int32(math.Round(float64((float32(segment.Position.X) + 0.25) * screenPerWorldX)))
				y = int32(math.Round(float64((float32(segment.Position.Y) + 0.25) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY * 0.75)))
				w = int32(math.Round(float64(screenPerWorldX / 2)))
			} else if dirNext.X == 0 && dirNext.Y == -1 {
				x = int32(math.Round(float64((float32(segment.Position.X) + 0.25) * screenPerWorldX)))
				y = int32(math.Round(float64(float32(segment.Position.Y) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY * 0.75)))
				w = int32(math.Round(float64(screenPerWorldX / 2)))
			} else if dirNext.X == 1 && dirNext.Y == 0 {
				x = int32(math.Round(float64((float32(segment.Position.X) + 0.25) * screenPerWorldX)))
				y = int32(math.Round(float64((float32(segment.Position.Y) + 0.25) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY / 2)))
				w = int32(math.Round(float64(screenPerWorldX * 0.75)))
			} else if dirNext.X == -1 && dirNext.Y == 0 {
				x = int32(math.Round(float64(float32(segment.Position.X) * screenPerWorldX)))
				y = int32(math.Round(float64((float32(segment.Position.Y) + 0.25) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY / 2)))
				w = int32(math.Round(float64(screenPerWorldX * 0.75)))
			}

			rl.DrawRectangle(x, y, w, h, rl.Black)

		} else {

			// Render the straight segment

			var x, y, w, h int32

			if dirNext.X == 0 && dirNext.Y == 1 {
				x = int32(math.Round(float64((float32(segment.Position.X) + 0.25) * screenPerWorldX)))
				y = int32(math.Round(float64(float32(segment.Position.Y) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY)))
				w = int32(math.Round(float64(screenPerWorldX / 2)))
			} else if dirNext.X == 0 && dirNext.Y == -1 {
				x = int32(math.Round(float64((float32(segment.Position.X) + 0.25) * screenPerWorldX)))
				y = int32(math.Round(float64((float32(segment.Position.Y) + 0) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY)))
				w = int32(math.Round(float64(screenPerWorldX / 2)))
			} else if dirNext.X == 1 && dirNext.Y == 0 {
				x = int32(math.Round(float64(float32(segment.Position.X) * screenPerWorldX)))
				y = int32(math.Round(float64((float32(segment.Position.Y) + 0.25) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY / 2)))
				w = int32(math.Round(float64(screenPerWorldX)))
			} else if dirNext.X == -1 && dirNext.Y == 0 {
				x = int32(math.Round(float64((float32(segment.Position.X) + 0) * screenPerWorldX)))
				y = int32(math.Round(float64((float32(segment.Position.Y) + 0.25) * screenPerWorldY)))
				h = int32(math.Round(float64(screenPerWorldY / 2)))
				w = int32(math.Round(float64(screenPerWorldX)))
			}

			rl.DrawRectangle(x, y, w, h, rl.Black)

		}

	}

	// Render the apple as a red elipse

	rl.DrawEllipse(int32(float32(game.ApplePosition.X)*screenPerWorldX+screenPerWorldX/2), int32(float32(game.ApplePosition.Y)*screenPerWorldY+screenPerWorldY/2), screenPerWorldX/2, screenPerWorldY/2, rl.Red)

	// Write the score at the top right corner

	scoreText := fmt.Sprintf("Score: %d", game.Score)
	textWidth := rl.MeasureText(scoreText, 20)

	rl.DrawText(scoreText, app.ScreenSize.X-textWidth-5, 5, 20, rl.Black)

}

func RenderPlaying(game *Game, app *App) {

	rl.BeginDrawing()
	RenderGame(game, app)
	rl.EndDrawing()
}

func RenderPaused(game *Game, app *App) {

	rl.BeginDrawing()

	RenderGame(game, app)

	rl.DrawRectangle(0, 0, app.ScreenSize.X, app.ScreenSize.Y, rl.Fade(rl.RayWhite, 0.5))

	textWidth := rl.MeasureText("Paused", 20)

	rl.DrawText("Paused", app.ScreenSize.X/2-textWidth/2, app.ScreenSize.Y/2, 20, rl.Black)

	rl.EndDrawing()
}

func RenderGameOver(game *Game, app *App) {

	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	textWidth := rl.MeasureText("Game Over", 20)

	rl.DrawText("Game Over", app.ScreenSize.X/2-textWidth/2, app.ScreenSize.Y/2, 20, rl.Black)

	rl.EndDrawing()
}

func RenderHomeScreen(game *Game, app *App) {

	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	textWidth := rl.MeasureText("Press Space to Start", 20)

	rl.DrawText("Press Space to Start", app.ScreenSize.X/2-textWidth/2, app.ScreenSize.Y/2, 20, rl.Black)

	rl.EndDrawing()
}
