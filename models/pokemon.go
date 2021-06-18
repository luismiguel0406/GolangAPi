package models

type PokemonModel struct {
	Name string `json:"name"`
}

type PokemonSpecies struct {
	Name string `json:"name"`
}

type ResponsePokemon struct {
   Name string `json:"name"`

}