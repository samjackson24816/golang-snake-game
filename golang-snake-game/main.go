package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// Enum
type GameState int

func (g GameState) String() string {
	return [...]string{"Playing", "Paused", "GameOver", "HomeScreen"}[g]
}

const (
	Playing GameState = iota
	Paused
	GameOver
	HomeScreen
)

type TailSegment struct {
	Position Vector2Int
	Lifetime uint
}

// The state of the pure game
type Game struct {
	State                GameState
	Dt                   float32
	MoveTime             float32
	TimeSinceLastMove    float32
	MoveNow              bool
	WorldSize            Vector2Int
	LastPosition         Vector2Int
	Position             Vector2Int
	Velocity             Vector2Int
	ApplePosition        Vector2Int
	Score                uint
	Tail                 []TailSegment
	StartingTailSegments uint
}

func (g *Game) NewApple() {
	// Make a list of all the valid positions
	validPositions := make([]Vector2Int, 0, g.WorldSize.X*g.WorldSize.Y)

	for x := int32(0); x < g.WorldSize.X; x++ {
		for y := int32(0); y < g.WorldSize.Y; y++ {
			if x == g.Position.X && y == g.Position.Y {
				continue
			}

			stop := false

			for _, segment := range g.Tail {
				if (Vector2Int{x, y}).Equals(segment.Position) {
					stop = true
					continue
				}
			}

			if stop {
				continue
			}

			validPositions = append(validPositions, Vector2Int{x, y})
		}
	}

	// Remove the positions that are occupied by the snake
	for _, segment := range g.Tail {
		for i, pos := range validPositions {
			if pos.Equals(segment.Position) {
				validPositions = append(validPositions[:i], validPositions[i+1:]...)
			}
		}
	}

	// Choose a random position from the valid positions
	g.ApplePosition = validPositions[rl.GetRandomValue(0, int32(len(validPositions)-1))]

}

func (g *Game) Init(worldSize Vector2Int) {
	g.WorldSize = worldSize
}

func (g *Game) Log() {
	fmt.Println("Position: ", g.Position)
	fmt.Println("Velocity: ", g.Velocity)
	fmt.Println("MoveNow: ", g.MoveNow)
}

func NewGame(worldSize Vector2Int) (g Game) {
	g.Init(worldSize)
	return g
}

// The state of the user interface
type App struct {
	ScreenSize Vector2Int
}

func (a *App) Init(screenSize Vector2Int) {
	a.ScreenSize = screenSize
}

func NewApp(screenSize Vector2Int) (a App) {
	a.Init(screenSize)
	return a
}

func main() {

	game := NewGame(Vector2Int{8, 8})
	app := NewApp(Vector2Int{800, 800})

	game.MoveTime = 0.5
	game.StartingTailSegments = 3

	rl.InitWindow(int32(app.ScreenSize.X), int32(app.ScreenSize.Y), "Snake")

	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	Start(&game, &app)

	for !rl.WindowShouldClose() {
		game.Dt = rl.GetFrameTime()
		Update(&game, &app)
		Render(&game, &app)
	}
}
