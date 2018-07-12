package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Params struct {
	Url    string
	ApiKey string
}

type client struct {
	url    string
	apiKey string
}

func Create(params Params) *client {
	return &client{
		url:    params.Url,
		apiKey: params.ApiKey,
	}
}

func (c *client) GetCountries() ([]Country, error) {
	var countries []Country

	resp, err := c.request("get_countries")
	if err != nil {
		return nil, fmt.Errorf("getting the countries: %s", err)
	}

	if err = json.Unmarshal(resp, &countries); err != nil {
		return nil, fmt.Errorf("unmarshaling json: %s", err)
	}

	return countries, nil
}

func (c *client) GetLeagues() ([]League, error) {
	var leagues []League

	resp, err := c.request("get_leagues")
	if err != nil {
		return nil, fmt.Errorf("getting the leagues: %s", err)
	}

	if err = json.Unmarshal(resp, &leagues); err != nil {
		return nil, fmt.Errorf("unmarshaling json: %s", err)
	}

	return leagues, nil
}

func (c *client) request(action string) ([]byte, error) {
	httpClient := &http.Client{}
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("action", action)
	query.Add("APIkey", c.apiKey)
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
