package service

import (
	"github.com/EGS-software/superset-pipeline/etl-go/internal/model"
)

func transformPokemon(apiData model.PokemonAPI) model.PokemonDB {
	bst := 0
	for _, s := range apiData.Stats {
		bst += s.BaseStat
	}

	gen := 1
	if apiData.ID > 151 {
		gen = 2
	} // Lógica simplificada para teste

	return model.PokemonDB{
		ID:         apiData.ID,
		Name:       apiData.Name,
		Generation: gen,
		TotalStats: bst,
	}
}
