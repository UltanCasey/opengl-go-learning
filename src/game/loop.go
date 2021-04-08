package game

import (
	"game-dev-opengl/src/opengl"
	"github.com/veandco/go-sdl2/sdl"
)

// Generate a new game loop for handling user input on both keyboard and mouse.
// Requires a pointer to the current SDL window, a shader program to be rendered, a
// vertex array and a model to be rendered. Renders as fast as permitted by the GPU
// and CPU.
//
// TODO: Add ability to set the render delay to allow framerate to be capped.
// TODO: Only rerender on updates occuring, otherwise ignore? Not sure if this will work or not.
func NewGameLoop(w *sdl.Window, shaderProgram opengl.ProgramID, vao opengl.VAOID, model []float32) {
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
			opengl.Clear()
			opengl.Draw(shaderProgram, vao, model)
			w.GLSwap()
		}
	}
}