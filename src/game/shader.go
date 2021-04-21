package game

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"io/ioutil"
	"strings"
)

type ShaderID uint32

type Shader struct {
	ShaderID ShaderID
	Type uint32
	Path string
}

// CreateShader returns a shader given its type and file path.
func CreateShader(path string, shaderType uint32) Shader {
	s := Shader{
		Type:     shaderType,
		Path:     path,
	}
	s.compile()
	return s
}

// compile reads a shader from file of type shaderType. Returns the
// compiled shader ID for use by the shader program.
func (s *Shader) compile() {

	// Read in source from file.
	source, err := ioutil.ReadFile(fmt.Sprintf("./assets/shaders/%s", s.Path))
	if err != nil {
		panic(fmt.Sprintf("Unable to create shader from file '%s': %v", s.Path, err))
	}

	s.compileShader(string(source))
}

// compileShader takes the shader source, appends the termination character, and
// compiles the shader. Returns the ShaderID for use by the shader program.
func (s *Shader) compileShader(source string) {

	// Create shader.
	shader := gl.CreateShader(s.Type)
	csource, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csource, nil)

	// Clean shader from memory.
	free()

	// Compile shader.
	gl.CompileShader(shader)

	// Handle status of shader compilation.
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
	}
	s.ShaderID = ShaderID(shader)

}