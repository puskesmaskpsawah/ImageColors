package models

type InputImage struct {
	Pixels      []string `json:"pixels"`
	Base64Image string   `json:"base64image"`
}
