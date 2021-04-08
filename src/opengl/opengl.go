package opengl

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"io/ioutil"
	"strings"
	"unsafe"
)

type ShaderID uint32
type ProgramID uint32
type VAOID uint32
type VBOID uint32

// Return the current OpenGL version from the current context.
func GetVersion() string {
	return gl.GoStr(gl.GetString(gl.VERSION))
}

// Load shader from file of type shaderType. Returns the compiled shader ID for
// use by the shader program.
func createShaderFromFile(file string, shaderType uint32) ShaderID {

	// Read in source from file.
	source, err := ioutil.ReadFile(fmt.Sprintf("./assets/shaders/%s", file))
	if err != nil {
		panic(fmt.Sprintf("Unable to create shader from file '%s': %v", file, err))
	}

	return compileShader(string(source), shaderType)
}

// Given the shader source, append the termination character and compile the shader.
// Return the ShaderID for use by the shader program.
func compileShader(source string, shaderType uint32) ShaderID {

	// Create shader.
	shader := gl.CreateShader(shaderType)
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
		log := strings.Repeat("\x00", int(logLength + 1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		panic("Failed to compile shader:\n" + log)
	}

	return ShaderID(shader)
}

// Create an OpenGL program with a vertex and fragment shader loaded from file.
// Return the ProgramID to the user to render.
func CreateProgram(vertFile, fragFile string) ProgramID {

	// Create shaders from file.
	vertexShader := createShaderFromFile(vertFile, gl.VERTEX_SHADER)
	fragmentShader := createShaderFromFile(fragFile, gl.FRAGMENT_SHADER)

	// Create program and attach both vertex and fragment shaders.
	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, uint32(vertexShader))
	gl.AttachShader(shaderProgram, uint32(fragmentShader))
	gl.LinkProgram(shaderProgram)

	// Handle respons from shader linking.
	var success int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength + 1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(log))
		panic("Failed to link program:\n" + log)
	}

	// Clean shaders from memory.
	gl.DeleteShader(uint32(vertexShader))
	gl.DeleteShader(uint32(fragmentShader))

	return ProgramID(shaderProgram)
}

// Create a vertex buffer object.
// Return the VBOID for use by the shader program.
func CreateVertexBufferObject() VBOID {
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	return VBOID(VBO)
}

// Create a vertex array object.
// Return the VAOID for use by the shader program.
func CreateVertexArrayObject() VAOID {
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)
	return VAOID(VAO)
}

// Create buffer data for the model/data being displayed.
// Requires a target, the model data and the usage.
// Currently hard codes the data values as being of type float32.
func CreateBufferData(target uint32, data []float32, usage uint32) {
	var float32 float32
	gl.BufferData(target, len(data)*int(unsafe.Sizeof(float32)), gl.Ptr(data), usage)
}

// Create a vertex array pointer.
func CreateVertexArrayPointer(size int32, attribType uint32) {
	var float32 float32
	gl.VertexAttribPointer(0, size, attribType, false, size*int32(unsafe.Sizeof(float32)), nil)
	gl.EnableVertexAttribArray(0)
}

// Unbind the vertex array.
func DestroyVertexArray() {
	gl.BindVertexArray(0)
}

// Clear the open GL screen.
func Clear() {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// Draw the current shader program using the vao and provided data.
func Draw(shaderProgram ProgramID, vao VAOID, data []float32) {
	gl.UseProgram(uint32(shaderProgram))
	gl.BindVertexArray(uint32(vao))
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(data)/3))
}