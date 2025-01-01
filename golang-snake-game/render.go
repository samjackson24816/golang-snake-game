package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func StartRender(game *Game, app *App) {

	app.Texture = rl.LoadTexture("assets/spritesheet.png")

	// Initialize background offsets
	TS := float32(16.0)

	app.BackgroundTexOffsets = []rl.Rectangle{
		rl.NewRectangle(0, 0, TS, TS),
		rl.NewRectangle(TS, 0, TS, TS),
		rl.NewRectangle(TS*2, 0, TS, TS),
		rl.NewRectangle(TS*3, 0, TS, TS),
	}

	app.AppleTexOffset = rl.NewRectangle(TS*4, 0, TS, TS)

	app.SnakeHeadOffset = rl.NewRectangle(TS*5, 0, TS, TS)

	app.SnakeTailStraightOffset = rl.NewRectangle(TS*6, 0, TS, TS)

	app.SnakeTailCornerOffsets = []rl.Rectangle{
		rl.NewRectangle(TS*7, 0, TS, TS),
		rl.NewRectangle(TS*8, 0, TS, TS),
	}

	app.SnakeTailEndOffset = rl.NewRectangle(TS*9, 0, TS, TS)

	app.BackgroundTileOffsets = make([]int32, game.WorldSize.X*game.WorldSize.Y)

	for i := int32(0); i < game.WorldSize.X; i++ {
		for j := int32(0); j < game.WorldSize.Y; j++ {
			index := (i*43 + (j * 61)) % int32(len(app.BackgroundTexOffsets))
			app.BackgroundTileOffsets[i+j*game.WorldSize.X] = index
		}
	}

}

func Render(game *Game, app *App) {

	rl.BeginDrawing()

	RenderBackground(game, app)
	switch game.State {
	case Playing:
		RenderApple(game, app)
		RenderSnake(game, app)
		RenderScoreUI(game, app)
	case Paused:
		RenderApple(game, app)
		RenderSnake(game, app)
		RenderScoreUI(game, app)
		RenderPaused(game, app)
	case GameOver:
		RenderGameOver(game, app)
	case HomeScreen:
		RenderHomeScreen(game, app)
	case WinScreen:
		RenderWinScreen(game, app)
	}

	rl.EndDrawing()
}

func RenderBackground(game *Game, app *App) {
	screenPerWorldX := float32(app.ScreenSize.X / game.WorldSize.X)
	screenPerWorldY := float32(app.ScreenSize.Y / game.WorldSize.Y)

	// Render the background
	for i := int32(0); i < game.WorldSize.X; i++ {
		for j := int32(0); j < game.WorldSize.Y; j++ {
			offset := app.BackgroundTexOffsets[app.BackgroundTileOffsets[i+j*game.WorldSize.X]]
			rl.DrawTexturePro(app.Texture, offset, rl.NewRectangle(float32(i)*screenPerWorldX, float32(j)*screenPerWorldY, screenPerWorldX, screenPerWorldY), rl.NewVector2(0, 0), 0, rl.White)
		}
	}
}

func RenderWinScreen(game *Game, app *App) {

	textWidth := rl.MeasureText("You Win!", 40)

	rl.DrawText("You Win!", app.ScreenSize.X/2-textWidth/2, app.ScreenSize.Y/2, 40, rl.White)
}

func RenderApple(game *Game, app *App) {
	// Render the apple
	screenPerWorldX := float32(app.ScreenSize.X / game.WorldSize.X)
	screenPerWorldY := float32(app.ScreenSize.Y / game.WorldSize.Y)

	rl.DrawTexturePro(app.Texture, app.AppleTexOffset, rl.NewRectangle(float32(game.ApplePosition.X)*screenPerWorldX, float32(game.ApplePosition.Y)*screenPerWorldY, screenPerWorldX, screenPerWorldY), rl.NewVector2(0, 0), 0, rl.White)

}

func RenderScoreUI(game *Game, app *App) {

	// Write the score at the top right corner

	scoreText := fmt.Sprintf("%d", game.Score)
	textWidth := rl.MeasureText(scoreText, 40)

	rl.DrawText(scoreText, app.ScreenSize.X-textWidth-5, 5, 40, rl.White)

	highScoreText := fmt.Sprintf("%d", game.HighScore)

	rl.DrawText(highScoreText, 5, 5, 40, rl.White)
}

func RenderSnake(game *Game, app *App) {
	screenPerWorldX := float32(app.ScreenSize.X / game.WorldSize.X)
	screenPerWorldY := float32(app.ScreenSize.Y / game.WorldSize.Y)

	// Render the head

	headAngle := game.Velocity.AngleDeg() + 90

	position := rl.NewVector2((float32(game.Position.X)+0.5)*screenPerWorldX, (float32(game.Position.Y)+0.5)*screenPerWorldY)

	rotationCenter := rl.NewVector2(screenPerWorldX/2, screenPerWorldY/2)

	rl.DrawTexturePro(app.Texture, app.SnakeHeadOffset, rl.NewRectangle(position.X, position.Y, screenPerWorldX, screenPerWorldY), rotationCenter, float32(headAngle), rl.White)

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
			// Last segment
			// 0 is up, 1 is right, 2 is down, 3 is left
			angle := dirNext.AngleDeg() + 90

			position := rl.NewVector2((float32(segment.Position.X)+0.5)*screenPerWorldX, (float32(segment.Position.Y)+0.5)*screenPerWorldY)

			rotationCenter := rl.NewVector2(screenPerWorldX/2, screenPerWorldY/2)

			rl.DrawTexturePro(app.Texture, app.SnakeTailEndOffset, rl.NewRectangle(position.X, position.Y, screenPerWorldX, screenPerWorldY), rotationCenter, float32(angle), rl.White)
			continue

		} else {
			prev = game.Tail[i-1].Position
		}

		dirPrev := segment.Position.Sub(prev)

		// If the directions aren't equal, the tail segment is a corner
		isCorner := dirNext != dirPrev

		// Render the tail segment, making it a rectangle that is thinner side to side

		if isCorner {

			// 0 is up, 1 is right, 2 is down, 3 is left
			angleNext := int32(dirNext.AngleDeg()) + 90
			anglePrev := int32(dirPrev.AngleDeg()) + 90

			// If the tail is turning right, the corner is the first corner in the sprite sheet
			// If the tail is turning left, the corner is the second corner in the sprite sheet
			// This is because the corners are in clockwise order in the sprite sheet

			index := 0

			if angleNext == anglePrev+90 || angleNext == anglePrev-270 {
				index = 0
			} else {
				index = 1
			}

			position := rl.NewVector2((float32(segment.Position.X)+0.5)*screenPerWorldX, (float32(segment.Position.Y)+0.5)*screenPerWorldY)

			rotationCenter := rl.NewVector2(screenPerWorldX/2, screenPerWorldY/2)

			rl.DrawTexturePro(app.Texture, app.SnakeTailCornerOffsets[index], rl.NewRectangle(position.X, position.Y, screenPerWorldX, screenPerWorldY), rotationCenter, float32(angleNext), rl.White)

		} else {

			// 0 is up, 1 is right, 2 is down, 3 is left
			angle := dirNext.AngleDeg() + 90

			position := rl.NewVector2((float32(segment.Position.X)+0.5)*screenPerWorldX, (float32(segment.Position.Y)+0.5)*screenPerWorldY)

			rotationCenter := rl.NewVector2(screenPerWorldX/2, screenPerWorldY/2)

			rl.DrawTexturePro(app.Texture, app.SnakeTailStraightOffset, rl.NewRectangle(position.X, position.Y, screenPerWorldX, screenPerWorldY), rotationCenter, float32(angle), rl.White)

		}

	}

}

func RenderPaused(game *Game, app *App) {

	rl.DrawRectangle(0, 0, app.ScreenSize.X, app.ScreenSize.Y, rl.Fade(rl.RayWhite, 0.5))

	textWidth := rl.MeasureText("Paused", 40)

	rl.DrawText("Paused", app.ScreenSize.X/2-textWidth/2, app.ScreenSize.Y/2, 40, rl.White)

}

func RenderGameOver(game *Game, app *App) {

	textWidth := rl.MeasureText("Game Over", 40)

	rl.DrawText("Game Over", app.ScreenSize.X/2-textWidth/2, app.ScreenSize.Y/2, 40, rl.White)

	RenderScores(game, app)

}

func RenderScores(game *Game, app *App) {
	scoreWidth := rl.MeasureText(fmt.Sprintf("Score: %d", game.Score), 30)

	rl.DrawText(fmt.Sprintf("Score: %d", game.Score), app.ScreenSize.X/2-scoreWidth/2, app.ScreenSize.Y/2+50, 30, rl.White)

	highScoreWidth := rl.MeasureText(fmt.Sprintf("High Score: %d", game.HighScore), 30)

	rl.DrawText(fmt.Sprintf("High Score: %d", game.HighScore), app.ScreenSize.X/2-highScoreWidth/2, app.ScreenSize.Y/2+80, 30, rl.White)

}

func RenderHomeScreen(game *Game, app *App) {

	rl.ClearBackground(rl.RayWhite)

	textWidth := rl.MeasureText("Press Space to Start", 40)

	rl.DrawText("Press Space to Start", app.ScreenSize.X/2-textWidth/2, app.ScreenSize.Y/2, 40, rl.White)

	highScoreWidth := rl.MeasureText(fmt.Sprintf("High Score: %d", game.HighScore), 30)

	rl.DrawText(fmt.Sprintf("High Score: %d", game.HighScore), app.ScreenSize.X/2-highScoreWidth/2, app.ScreenSize.Y/2+80, 30, rl.White)
}
