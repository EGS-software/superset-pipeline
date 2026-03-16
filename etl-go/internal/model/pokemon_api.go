package model

type PokemonAPI struct {
	ID    int    `json:"ID"`
	Name  string `json:"name"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
	} `json:"stats"`
}
