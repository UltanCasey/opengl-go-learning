package opengl

import (
	"encoding/json"
	"fmt"
	"game-dev-opengl/src/structs"
	"io/ioutil"
	"os"
)

// Given a file name load in the file referenced from the /assets/models folder
// and parse it into the ModelValues struct. Return the model vertices for
// rendering by the shader program.
func LoadModel(file string) []float32 {

	// Open the model file.
	f, err := os.Open(fmt.Sprintf("./assets/models/%s", file))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Parse the values.
	b, err := ioutil.ReadAll(f)
	var model structs.ModelValues
	json.Unmarshal(b, &model)

	return model.Model.Vertices
}
