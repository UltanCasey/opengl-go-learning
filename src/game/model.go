package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ModelValues struct {
	Model   Model   `json:"model"`
	Texture Texture `json:"texture"`
}

type Model struct {
	Vertices []float32 `json:"vertices"`
	Indices []uint32	`json:"indices"`
}

type Texture struct {
}

// LoadModel loads in the file referenced from the /assets/models folder
// and parses it into the ModelValues struct. Return the model vertices for
// rendering by the shader program.
func LoadModel(file string) Model {

	// Open the model file.
	f, err := os.Open(fmt.Sprintf("./assets/models/%s", file))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Parse the values.
	b, err := ioutil.ReadAll(f)
	var model ModelValues
	json.Unmarshal(b, &model)

	return model.Model
}
