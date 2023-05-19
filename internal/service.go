package internal

import "fmt"

type Service interface {
	AddPokemon(pokemon Pokemon) error
	RemovePokemon(id int) error
	SearchPokemon(field, value string) []Pokemon
}

type PokemonService struct {
	repo Repository
}

func NewPokemonService(repo Repository) *PokemonService {
	return &PokemonService{
		repo: repo,
	}
}

func (s *PokemonService) AddPokemon(pokemon Pokemon) error {
	existingPokemon := s.repo.GetPokemonById(pokemon.ID)
	if existingPokemon != nil {
		return fmt.Errorf("Pokemon with ID %d already exists", pokemon.ID)
	}

	s.repo.AddPokemon(pokemon)
	return nil
}

func (s *PokemonService) RemovePokemon(id int) error {
	existingPokemon := s.repo.GetPokemonById(id)

	if existingPokemon == nil {
		return fmt.Errorf("Pokemon with ID %d not found", id)
	}

	s.repo.RemovePokemon(id)
	return nil
}

func (s *PokemonService) SearchPokemon(field, value string) []Pokemon {
	return s.repo.SearchPokemon(field, value)
}
