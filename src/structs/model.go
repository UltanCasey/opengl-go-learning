package structs

type ModelValues struct {
	Model Model `json:"model"`
	Texture Texture `json:"texture"`
}

type Model struct {
	Vertices []float32
}

type Texture struct {

}
