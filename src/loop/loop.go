package loop

import (
	"fmt"
	"game-dev-opengl/src/game"
	"game-dev-opengl/src/structs"
	"github.com/veandco/go-sdl2/sdl"
)

// NewGameLoop generates a new game loop. This creates a game Object which stores
// the window, the shader program, model and game state. When an event occurs it
// is handled and the shader program is updated. Shader hot loading is also supported.
func NewGameLoop() {

	// Create required resources and channels.
	g := game.NewGame()
	events := make(chan structs.Event)

	// Handle events which occur.
	go handleEvents(&g, events)

	// Main loop.
	for g.Running {

		// Poll for new events.
		pollEvent(&g, events)

		// Reload program.
		if g.Reloading {
			g.Program = game.NewShaderProgram(g.Program.VertexShader.Path, g.Program.FragmentShader.Path)
			g.Reloading = false
		}

		// Draw game.
		game.Clear()
		game.Draw(g.Program, g.VAOID, g.Model)
		g.Window.Window.GLSwap()
	}
}

// pollEvent retrieves latest user interactions and send those as type structs.Event to
// the events channel to be dealt with by the event handler. If the event is
// a reload keyboard event the shader program is reloaded, allowing the hot
// loading of shaders.
func pollEvent(g *game.Game, events chan structs.Event) {

	// Retrieve latest event.
	event := sdl.PollEvent()
	if event == nil {
		return
	}
	switch t := event.(type) {

	// Send new events to be handled.
	case *sdl.QuitEvent:
		sendQuitEvent(events, t)
	case *sdl.KeyboardEvent:
		sendKeyboardEvent(events, t)
	}
}

// sendKeyboardEvent handles keyboard key strokes and sends the relevant event
// into the global event channel to be handled by the main game loop.
func sendKeyboardEvent(events chan structs.Event, event *sdl.KeyboardEvent) {
	switch event.Keysym.Sym {
	case sdl.K_r:
		events <- structs.Event{
			Source: "Player",
			Type:   structs.Reload,
		}
	}
}

// sendQuitEvent handles the main quit event occuring and when sent to the main
// game loop terminates the game execution.
func sendQuitEvent(events chan structs.Event, event *sdl.QuitEvent) {
	events <- structs.Event{
		Source: "Player",
		Type:   structs.Quit,
	}
}


// Asynchronous loop to handle the events being transmitted due to user
// and game actions occurring.
func handleEvents(g *game.Game, events chan structs.Event) {
	for  {
		event := <-events
		switch event.Type {
		case structs.Quit:
			g.Running = false
		case structs.Reload:
			fmt.Println("Reloading program.")
			g.Reloading = true
		}
	}
}