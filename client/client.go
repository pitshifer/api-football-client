package client

import (
	"encoding/json"
	"fmt"
	"io"
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

	result, err := c.request("get_countries")
	if err != nil {
		return nil, fmt.Errorf("getting the countries: %s", err)
	}

	b, err := ioutil.ReadAll(result)

	if err = json.Unmarshal(b, &countries); err != nil {
		return nil, err
	}

	return countries, nil
}

func (c *client) request(action string) (io.Reader, error) {
	httpClient := &http.Client{}
	u, err := url.Parse(c.url)
	if err != nil {
		return nil, err
	}
	query := u.Query()
	query.Add("action", action)
	query.Add("APIkey", c.apiKey)
	u.RawQuery = query.Encode()

	fmt.Printf("url : %s\n", u.String())

	resp, err := httpClient.Get(u.String())
	//defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}
