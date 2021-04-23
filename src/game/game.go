package game

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type VAOID uint32
type VBOID uint32
type EBOID uint32

type Game struct {
	Window    Window
	Program   *Program
	Model     Model
	VAOID     VAOID
	Running   bool
	Reloading bool
}

// NewGame is the main function to create a new instance of the Game object.
// This method creates a window, initialises the OpenGL context and buffers
// and then returns the new Game object.
func NewGame() Game {

	// Create Window
	window := NewWindow(1280, 720, "Learning OpenGL")

	// Initialise Open GL.
	gl.Init()

	// Generate OpenGL program by compiling and linking vertex and fragment shaders.
	program := NewShaderProgram("vertex.glsl", "fragment.glsl")
	model := LoadModel("square.json")

	_ = CreateVertexBufferObject()
	vao := CreateVertexArrayObject()
	_ = CreateElementBufferObject()

	CreateBufferDataFloat(gl.ARRAY_BUFFER, model.Vertices, gl.STATIC_DRAW)
	CreateBufferDataInt(gl.ELEMENT_ARRAY_BUFFER, model.Indices, gl.STATIC_DRAW)

	CreateVertexAttributePointer(0,3, gl.FLOAT, 5, nil)
	gl.EnableVertexAttribArray(0)
	CreateVertexAttributePointer(2, 2, gl.FLOAT, 5, gl.PtrOffset(3*4))

	DestroyVertexArray()

	return Game{
		Window:    window,
		Program:   program,
		Model:     model,
		VAOID:     VAOID(vao),
		Running:   true,
		Reloading: false,
	}
}
