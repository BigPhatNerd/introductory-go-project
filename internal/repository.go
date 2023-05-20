package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Repository interface {
	GetPokemonById(id int) *Pokemon
	AddPokemon(pokemon Pokemon)
	RemovePokemon(id int)
	SearchPokemon(field, value string) []Pokemon
}

type PokemonRepository struct {
	pokedex []Pokemon
}

func NewPokemonRepository() *PokemonRepository {
	return &PokemonRepository{
		pokedex: []Pokemon{},
	}
}

func (r *PokemonRepository) GetPokemonById(id int) *Pokemon {
	for _, pokemon := range r.pokedex {
		if pokemon.ID == id {
			return &pokemon
		}
	}
	return nil
}

func (r *PokemonRepository) AddPokemon(pokemon Pokemon) {
	r.pokedex = append(r.pokedex, pokemon)
}

func (r *PokemonRepository) RemovePokemon(id int) {
	for i, pokemon := range r.pokedex {
		if pokemon.ID == id {
			r.pokedex = append(r.pokedex[:i], r.pokedex[i+1:]...)
			break
		}
	}
}

func (r *PokemonRepository) SearchPokemon(field, value string) []Pokemon {
	var results []Pokemon

	for _, pokemon := range r.pokedex {
		switch field {
		case "name":
			if pokemon.Name == value {
				results = append(results, pokemon)
			}
		case "type":
			if pokemon.Type == value {
				results = append(results, pokemon)
			}
		}
	}
	return results
}

func LoadPokedexFromFile(filename string) (*PokemonRepository, error) {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read pokedex file: %w", err)
	}

	var pokedex []Pokemon
	err = json.Unmarshal(fileData, &pokedex)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal pokedex data: %w", err)
	}

	return &PokemonRepository{
		pokedex: pokedex,
	}, nil
}

func SavePokedexToFile(filename string, repo *PokemonRepository) error {
	pokedexData, err := json.Marshal(repo.pokedex)
	if err != nil {
		return fmt.Errorf("failed to marshal pokedex data: %w", err)
	}
	err = ioutil.WriteFile(filename, pokedexData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write pokedex file: %w", err)
	}

	return nil
}
