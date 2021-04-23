package game

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"unsafe"
)

// CreateVertexBufferObject creates a vertex buffer object.
// Returns the VBOID for use by the shader program.
func CreateVertexBufferObject() VBOID {
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	return VBOID(VBO)
}

// CreateVertexArrayObject creates a vertex array object.
// Returns the VAOID for use by the shader program.
func CreateVertexArrayObject() VAOID {
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	gl.BindVertexArray(VAO)
	return VAOID(VAO)
}

// CreateElementBufferObject creates an element buffer object.
// Returns a EBOID for use by the shader program.
func CreateElementBufferObject() EBOID {
	var EBO uint32
	gl.GenBuffers(1, &EBO)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, EBO)
	return EBOID(EBO)
}

// CreateBufferDataFloat creates a buffer for the model/data being displayed.
// Requires a target, the model data and the usage.
// Currently hard codes the data values as being of type float32.
func CreateBufferDataFloat(target uint32, data []float32, usage uint32) {
	var float32 float32
	gl.BufferData(target, len(data)*int(unsafe.Sizeof(float32)), gl.Ptr(data), usage)
}

// CreateBufferDataInt creates a buffer for the model/data being displayed.
// Requires a target, the model data and the usage.
// Currently hard codes the data values as being of type float32.
func CreateBufferDataInt(target uint32, data []uint32, usage uint32) {
	var uint32 uint32
	gl.BufferData(target, len(data)*int(unsafe.Sizeof(uint32)), gl.Ptr(data), usage)
}

// CreateVertexAttributePointer sets the pointer for the OpenGL vertex array.
func CreateVertexAttributePointer(index uint32, size int32, attribType uint32, stride int32, pointer unsafe.Pointer) {
	var float32 float32
	gl.VertexAttribPointer(index, size, attribType, false, stride*int32(unsafe.Sizeof(float32)), pointer)
}

// DestroyVertexArray uses the index value of 0 to remove the
// vertex array from memory.
func DestroyVertexArray() {
	gl.BindVertexArray(0)
}

// Clear the open GL screen.
func Clear() {
	gl.ClearColor(0.0, 0.0, 0.0, 0.0)
	gl.Clear(gl.COLOR_BUFFER_BIT)
}

// Draw the current shader program using the vao and provided vertices
// and indices.
func Draw(program *Program, vao VAOID, count int) {
	gl.UseProgram(uint32(program.ProgramID))
	gl.BindVertexArray(uint32(vao))
	gl.DrawElements(gl.TRIANGLES, int32(count), gl.UNSIGNED_INT, gl.PtrOffset(0))
}
