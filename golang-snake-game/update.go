package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func Start(game *Game, app *App) {
	ResetGame(game)
}

func Update(game *Game, app *App) {

	UpdateHighScore(game)
	switch game.State {
	case Playing:
		game.TimeSinceLastMove = game.TimeSinceLastMove + game.Dt

		if game.TimeSinceLastMove >= game.MoveTime {
			game.MoveNow = true
			game.TimeSinceLastMove = 0
		} else {
			game.MoveNow = false
		}

		UpdatePlayerVelocity(game)

		PauseIfSpacePressed(game)

		if game.MoveNow {

			TickTailSegments(game)
			AddNewTailSegment(game)
			MovePlayer(game)
			KillIfHitTail(game)
			KillIfOutOfBounds(game)
			CollectAppleIfPlayerOnApple(game)

		}

	case Paused:
		UnpauseIfSpacePressed(game)
	case GameOver, WinScreen:
		GoToHomeScreenIfSpacePressed(game)
	case HomeScreen:
		SwitchToPlayingIfSpacePressed(game)
	}
}

func UpdateHighScore(game *Game) {
	if game.Score > game.HighScore {
		game.HighScore = game.Score
	}
}

func AddNewTailSegment(game *Game) {
	game.Tail = append(game.Tail, TailSegment{Position: game.Position, Lifetime: game.Score + game.StartingTailSegments})
}

func UpdatePlayerVelocity(game *Game) {

	vel := Vector2Int{0, 0}

	if rl.IsKeyPressed(rl.KeyD) {
		vel = vel.Add(Vector2Int{1, 0})
	} else if rl.IsKeyPressed(rl.KeyA) {
		vel = vel.Add(Vector2Int{-1, 0})
	} else if rl.IsKeyPressed(rl.KeyW) {
		vel = vel.Add(Vector2Int{0, -1})
	} else if rl.IsKeyPressed(rl.KeyS) {
		vel = vel.Add(Vector2Int{0, 1})
	}

	if len(game.Tail) > 0 && game.Position.Add(vel).Equals(game.Tail[len(game.Tail)-1].Position) {
		vel = Vector2Int{0, 0}
	}

	if !vel.Equals(Vector2Int{0, 0}) {
		game.InputVelocity = vel
	}

}

func MovePlayer(game *Game) {

	if !game.InputVelocity.Equals(Vector2Int{0, 0}) {
		game.Velocity = game.InputVelocity
	}

	game.Position = game.Position.Add(game.Velocity)
	game.LastPosition = game.Position
}

func KillIfOutOfBounds(game *Game) {
	if game.Position.X < 0 || game.Position.X >= game.WorldSize.X || game.Position.Y < 0 || game.Position.Y >= game.WorldSize.Y {
		EndGame(game)
	}
}

func KillIfHitTail(game *Game) {
	for i := 0; i < len(game.Tail); i++ {
		if game.Position.Equals(game.Tail[i].Position) {
			EndGame(game)
		}
	}
}

func CollectAppleIfPlayerOnApple(game *Game) {
	if game.Position.Equals(game.ApplePosition) {
		noSpace := game.NewApple()
		if noSpace {
			WinGame(game)
		}
		game.Score += 1

		// Increase all tail segment lifetimes
		for i := 0; i < len(game.Tail); i++ {
			game.Tail[i].Lifetime += 1
		}
	}

}

func WinGame(game *Game) {
	ChangeState(game, WinScreen)
	ResetGame(game)
}

func TickTailSegments(game *Game) {
	newTail := make([]TailSegment, 0, len(game.Tail))
	for i := 0; i < len(game.Tail); i++ {
		game.Tail[i].Lifetime -= 1

		if game.Tail[i].Lifetime > 0 {
			newTail = append(newTail, game.Tail[i])
		}
	}

	game.Tail = newTail
}

func SwitchToPlayingIfSpacePressed(game *Game) {
	if rl.IsKeyPressed(rl.KeySpace) {
		ChangeState(game, Playing)
	}
}

func PauseIfSpacePressed(game *Game) {
	if rl.IsKeyPressed(rl.KeySpace) {
		ChangeState(game, Paused)
	}
}

func UnpauseIfSpacePressed(game *Game) {
	if rl.IsKeyPressed(rl.KeySpace) {
		ChangeState(game, Playing)
		game.TimeSinceLastMove = 0
	}
}

func GoToHomeScreenIfSpacePressed(game *Game) {
	if rl.IsKeyPressed(rl.KeySpace) {
		ChangeState(game, HomeScreen)
	}
}

func EndGame(g *Game) {
	ResetGame(g)
	ChangeState(g, GameOver)
}

func ResetGame(g *Game) {
	g.Tail = make([]TailSegment, 0, 100)
	g.Position = Vector2Int{g.WorldSize.X / 2, g.WorldSize.Y / 2}
	g.Velocity = Vector2Int{0, 1}
	g.InputVelocity = Vector2Int{0, 1}
	g.NewApple()
	g.Score = 0
}

func ChangeState(game *Game, desiredState GameState) {
	game.State = desiredState
}
