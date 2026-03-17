package service

import (
	"github.com/EGS-software/superset-pipeline/etl-go/internal/model"
)

func TransformPokemon(apiData model.PokemonAPI) model.PokemonDB {
	bst := 0
	for _, s := range apiData.Stats {
		bst += s.BaseStat
	}

	gen := 1
	if apiData.ID > 151 {
		gen = 2
	}

	typeOne := ""
	if len(apiData.Types) > 0 {
		typeOne := apiData.Types[0].Type.Name
	}

	typeTwo := ""
	if len(apiData.Types) > 0 {
		typeTwo := apiData.Types[1].Type.Name
	}

	return model.PokemonDB{
		ID:         apiData.ID,
		Name:       apiData.Name,
		Generation: gen,
		TotalStats: bst,
		TypeOne:    typeOne,
		TypeTwo:    typeTwo,
	}
}
