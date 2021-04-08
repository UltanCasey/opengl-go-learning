package main

import (
	"game-dev-opengl/src/game"
	"game-dev-opengl/src/opengl"
	"game-dev-opengl/src/window"
	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	windowHeight = 720
	windowWidth = 1280
	windowName = "Learning OpenGL"
	vertexShader = "vertex.glsl"
	fragmentShader = "fragment.glsl"
	modelFile = "square.json"
)

func main()  {

	// Create Window.
	w := window.NewWindow(windowWidth, windowHeight, windowName)
	defer w.Destroy()

	// Initialise Open GL.
	gl.Init()

	// Generate OpenGL program by compiling and linking vertex and fragment shaders.
	shaderProgram := opengl.CreateProgram(vertexShader, fragmentShader)
	model := opengl.LoadModel(modelFile)

	_ = opengl.CreateVertexBufferObject()
	VAO := opengl.CreateVertexArrayObject()

	opengl.CreateBufferData(gl.ARRAY_BUFFER, model, gl.STATIC_DRAW)
	opengl.CreateVertexArrayPointer(3, gl.FLOAT)
	opengl.DestroyVertexArray()

	// Create new game loop.
	game.NewGameLoop(w.Window, shaderProgram, VAO, model)
}
