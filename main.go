// Run go mod init internal to initialize the internal module. Then, run go mod
// tidy to add the dependencies to go.mod. Finally, run go build to build the
// application.
package main

import (
	"os"
	"strconv"
	"fmt"
	"github.com/BigPhatNerd/introductory-go-project/internal"
)



func main(){

	// Load the Pokedex from file
	repo, err := internal.LoadPokedexFromFile("database.json")
	if err != nil{
		fmt.Printf("Error loading Pokedex: %v\n", err)
		os.Exit(1)
	}

	// Create a new PokemonService with the repository
	service := internal.NewPokemonService(repo)

	// Parse the command line arguments

	args := os.Args
	if len(args) < 2{
		printHelp()
		os.Exit(1)
	}

	command := args[1]
	switch command {
	case "add":
		if len(args) != 5 {
			printHelp()
			os.Exit(1)
		}

		id, _ := strconv.Atoi(args[2])
		name := args[3]
		pokemonType := args[4]
		pokemon := internal.Pokemon{
			ID: id,
			Name: name,
			Type: pokemonType,
		}

		err := service.AddPokemon(pokemon)
		if err != nil {
			fmt.Printf("Error adding Pokemon %v\n", err)
		} else {
			fmt.Prinln("Pokemon added successfully")
		}
	case "remove":
		if len(args) != 3 {
			printHelp()
			os.Exit(1)
		}
		id, _ := strconv.Atoi(args[2])
		err := service.RemovePdokemon(id)
		if err != nil{
			fmt.Printf("Error removing Pokemon: %v\n", err)
		} else {
			fmt.Printf("Pokemon removed successfully")
		}
	case "search":
		if len(args) != 4 {
			printHelp()
			os.Exit(1)
		}
		field := args[2]
		value := args[3]
		results := service.SearchPokemon(field, value)
		fmt.Println("Search results:")
		for _, pokemon := range results {
			fmt.Printf("ID: %d, Name: %s, Type: %s\n", pokemon.ID, pokemon.Name, pokemon.Type)
		}
	default:
		printHelp()
		os.Exit(1)
	}

	//Save the updated Pokedex to file
	err = internal.SavePokedexToFile("database.json", repo)
	if err != nil {
		fmt.Pringf("Error saving Pokedex: %v\n", err)
		os.Exit(1)
	}
}

func printHelp(){
	fmt.Println("Usage: pokedex <command> [<arguments>]")
	fmt.Println("Commands:")
	fmt.Println("  add <id> <name> < type>    Add a new Pokemon to the Pokedex")
	fmt.Println("  remove <id>                Remove a Pokemon from the Pokedex")
	fmt.Println("  search <field> <value>     Search for a Pokemon in the Pokedex")
}