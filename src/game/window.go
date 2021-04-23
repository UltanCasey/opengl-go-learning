package game

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type Window struct {
	Height int32
	Width  int32
	Name   string
	Window *sdl.Window
}

// NewWindow creates a new game window given the name, width and height.
func NewWindow(width, height int32, name string) Window {
	w := Window{
		Width:  width,
		Height: height,
		Name:   name,
	}
	err := w.create()
	if err != nil {
		fmt.Println("error while creating window: ", err)
		return Window{}
	}
	return w
}

// Destroy the game window.
func (w *Window) Destroy() {
	sdl.Quit()
	w.Window.Destroy()
}

// create a new game window.
func (w *Window) create() error {

	// Create SDL Context.
	w.createSdlContext()

	// Create Window.
	var err error
	w.createSdlWindow()
	if err != nil {
		fmt.Println("error initialising window: ", err)
		return err
	}

	// Create GL context.
	w.Window.GLCreateContext()

	return nil

}

// createSdlContext sets up the SDL context values by initialising everything and setting the
// OpenGL values to version 4.1.
func (w *Window) createSdlContext() error {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("error initialising sdl: ", err)
		return err
	}
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MAJOR_VERSION, 4)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_MINOR_VERSION, 1)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_PROFILE_MASK, sdl.GL_CONTEXT_PROFILE_CORE)
	sdl.GLSetAttribute(sdl.GL_CONTEXT_FORWARD_COMPATIBLE_FLAG, 1)
	return nil
}

// createSdlWindow creates an SDL window using the windows defined values.
func (w *Window) createSdlWindow() error {
	var err error
	w.Window, err = sdl.CreateWindow(
		w.Name,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		w.Width, w.Height,
		sdl.WINDOW_OPENGL)
	if err != nil {
		return err
	}
	return nil
}
