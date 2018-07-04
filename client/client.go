package client

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
