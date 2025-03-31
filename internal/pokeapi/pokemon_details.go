package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// InvestigatePokemon -
func (c *Client) InvestigatePokemon(pokemon string) (PokemonInfo, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		pokeResp := PokemonInfo{}

		err := json.Unmarshal(val, &pokeResp)
		if err != nil {
			return PokemonInfo{}, err
		}

		return pokeResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonInfo{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonInfo{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonInfo{}, err
	}

	pokeResp := PokemonInfo{}
	err = json.Unmarshal(dat, &pokeResp)
	if err != nil {
		return PokemonInfo{}, err
	}

	c.cache.Add(url, dat)
	return pokeResp, nil
}
