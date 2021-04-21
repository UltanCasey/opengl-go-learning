package game

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

type ProgramID uint32

type Program struct {
	ProgramID ProgramID
	VertexShader Shader
	FragmentShader Shader
}

// NewShaderProgram takes a vertex and fragment shader file name and
// creates a new shader program given those. The pointer to this program is
// then returned.
func NewShaderProgram(vertPath, fragPath string) *Program {

	gl.Init()

	// Load Shaders.
	vertexShader := CreateShader(vertPath, gl.VERTEX_SHADER)
	fragmentShader := CreateShader(fragPath, gl.FRAGMENT_SHADER)

	// Build Program.
	return createShaderProgram(vertexShader, fragmentShader)
}

// createShaderProgram takes the Shader IDs and then loads the OpenGL
// program given those shaders.
func createShaderProgram(vertShader, fragShader Shader) *Program {
	program := Program{
		VertexShader:   vertShader,
		FragmentShader: fragShader,
	}
	program.load()
	return &program
}

// load handles the loading of the program into memory and the cleaning up of
// shaders from memory. An errors are also returned.
func (p *Program) load() {

	// Create program and attach both vertex and fragment shaders.
	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, uint32(p.VertexShader.ShaderID))
	gl.AttachShader(shaderProgram, uint32(p.FragmentShader.ShaderID))
	gl.LinkProgram(shaderProgram)

	// Handle response from shader linking.
	var success int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength + 1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))
		panic("Failed to link program:\n" + log)
	}
	p.ProgramID = ProgramID(shaderProgram)

	// Clean shaders from memory.
	gl.DeleteShader(uint32(p.VertexShader.ShaderID))
	gl.DeleteShader(uint32(p.FragmentShader.ShaderID))

}
