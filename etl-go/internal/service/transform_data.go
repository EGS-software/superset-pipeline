package service

import (
	"github.com/EGS-software/superset-pipeline/etl-go/internal/model"
)

func transformPokemon(apiData model.PokemonAPI) model.PokemonDB {
	bst := 0
	for _, s := range apiData.Stats {
		bst += s.BaseStat
	}

}
