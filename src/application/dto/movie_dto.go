package dto

type MovieDto struct {
	Name        string   `json:"name"`
	Synopsis    string   `json:"synopsis"`
	ReleaseTime string   `json:"releaseTime"`
	Performer   []string `json:"performer"`
	Director    string   `json:"director"`
	Cover       string   `json:"cover"`
	Address     string   `json:"address"`
}
