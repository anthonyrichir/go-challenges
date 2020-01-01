package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Response struct {
	Height int `json: "height"`
}

type Pokemon struct {
	Name string `json: "name"`
	Height int `json: "height"`
}

func main() {
	fmt.Println("Starting Challenge 04")
	http.HandleFunc("/", handlePokemons)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func handlePokemons(writer http.ResponseWriter, request *http.Request) {
	pokemons := make(map[string]Pokemon)
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go getPokemonHeight("ditto", &waitgroup, pokemons)
	waitgroup.Add(1)
	go getPokemonHeight("charizard", &waitgroup, pokemons)
	waitgroup.Add(1)
	go getPokemonHeight("weedle", &waitgroup, pokemons)
	waitgroup.Add(1)
	go getPokemonHeight("mew", &waitgroup, pokemons)
	waitgroup.Add(1)
	go getPokemonHeight("bulbasaur", &waitgroup, pokemons)
	waitgroup.Wait()

	fmt.Println("Finished execution")
	fmt.Println(pokemons)
	json.NewEncoder(writer).Encode(pokemons)
}

func getPokemonHeight(name string, w *sync.WaitGroup, pokemons map[string]Pokemon) {
	url := fmt.Sprintf("http://pokeapi.co/api/v2/pokemon/%s", name)
	response, _ := http.Get(url)
	responseData, _ := ioutil.ReadAll(response.Body)
	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	pokemon := Pokemon{
		Name:   name,
		Height: responseObject.Height,
	}
	pokemons[name] = pokemon
	w.Done()
}
