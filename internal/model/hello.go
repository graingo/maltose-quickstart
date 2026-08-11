package model

// HelloInput contains the data required to build a greeting.
type HelloInput struct {
	Name string `json:"name"`
}

// HelloOutput contains the generated greeting.
type HelloOutput struct {
	Greeting string `json:"greeting"`
}
