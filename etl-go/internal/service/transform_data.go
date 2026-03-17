package service

import (
	"github.com/EGS-software/superset-pipeline/etl-go/internal/model"
)

func TransformPokemon(apiData model.PokemonAPI) model.PokemonDB {
	bst := 0
	for _, s := range apiData.Stats {
		bst += s.BaseStat
	}

	gen := 0

	switch {
	case apiData.ID <= 151:
		gen = 1
	case apiData.ID <= 251:
		gen = 2
	case apiData.ID <= 386:
		gen = 3
	case apiData.ID <= 493:
		gen = 4
	case apiData.ID <= 649:
		gen = 5
	default:
		gen = 6
	}

	typeOne := ""
	if len(apiData.Types) > 0 {
		typeOne = apiData.Types[0].Type.Name
	}

	typeTwo := ""
	if len(apiData.Types) > 1 {
		typeTwo = apiData.Types[1].Type.Name
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
