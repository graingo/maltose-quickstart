package model

type HelloInput struct {
	Name string `json:"name"`
}

type HelloOutput struct {
	Tip string `json:"tip"`
}
