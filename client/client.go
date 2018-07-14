package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

type Params struct {
	Url    string
	ApiKey string
}

type client struct {
	url    string
	apiKey string
}

type requestParam struct {
	name  string
	value string
}

func Create(params Params) *client {
	return &client{
		url:    params.Url,
		apiKey: params.ApiKey,
	}
}

func (c *client) GetCountries() ([]Country, error) {
	var countries []Country

	resp, err := c.request("get_countries", []requestParam{})
	if err != nil {
		return nil, fmt.Errorf("getting the countries: %s", err)
	}

	if err = json.Unmarshal(resp, &countries); err != nil {
		return nil, fmt.Errorf("unmarshaling json: %s", err)
	}

	return countries, nil
}

func (c *client) GetLeagues(countryId int) ([]League, error) {
	var leagues []League
	var params []requestParam

	if countryId != 0 {
		params = append(params, requestParam{
			name:  "country_id",
			value: strconv.Itoa(countryId),
		})
	}
	resp, err := c.request("get_leagues", params)
	if err != nil {
		return nil, fmt.Errorf("getting the leagues: %s", err)
	}

	if err = json.Unmarshal(resp, &leagues); err != nil {
		return nil, fmt.Errorf("unmarshaling json: %s", err)
	}

	return leagues, nil
}

func (c *client) GetStandings(leagueId int) ([]Standings, error) {
	var standings []Standings
	var params []requestParam

	params = append(params, requestParam{
		name:  "league_id",
		value: strconv.Itoa(leagueId),
	})

	resp, err := c.request("get_standings", params)
	if err != nil {
		return nil, fmt.Errorf("getting the standings: %s", err)
	}

	if err = json.Unmarshal(resp, &standings); err != nil {
		return nil, fmt.Errorf("unmarshling json: %s", err)
	}

	return standings, nil
}

func (c *client) request(action string, params []requestParam) ([]byte, error) {
	httpClient := &http.Client{}
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}

	query := u.Query()
	query.Add("action", action)
	query.Add("APIkey", c.apiKey)
	for _, param := range params {
		query.Add(param.name, param.value)
	}
	u.RawQuery = query.Encode()

	resp, err := httpClient.Get(u.String())
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
