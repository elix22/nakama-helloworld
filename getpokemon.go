package main

// get_pokemon {"PokemonName": "raticate"}
// get_pokemon {"PokemonName": "pidgeotto"}
// other names  kakuna , metapod , wartortle , charmeleon , ivysaur

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/heroiclabs/nakama-common/runtime"
)

const apiBaseUrl = "https://pokeapi.co/api/v2"

func LookupPokemon(logger runtime.Logger, name string) (map[string]interface{}, error) {
	resp, err := http.Get(apiBaseUrl + "/pokemon/" + name)
	if err != nil {
		logger.Error("Failed request %v", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("Failed to read body %v", err.Error())
		return nil, err
	}
	if resp.StatusCode >= 400 {
		logger.Error("Failed request %v %v", resp.StatusCode, body)
		return nil, errors.New(string(body))
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)

	return result, err
}

func GetPokemon(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	// We'll assume payload was sent as JSON and decode it.
	var input map[string]string
	err := json.Unmarshal([]byte(payload), &input)
	if err != nil {
		return "", err
	}

	result, err := LookupPokemon(logger, input["PokemonName"])
	if err != nil {
		return "", err
	}

	response, err := json.Marshal(result)
	if err != nil {
		return "", err
	}
	return string(response), nil
}
