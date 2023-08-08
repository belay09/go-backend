package config

import (
	"net/http"
	"os"
	"github.com/hasura/go-graphql-client"
)

func GraphqlClient() *graphql.Client {
	
	adminSecret := os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET")
	hasuraEndPoint := os.Getenv("HASURA_GRAPHQL_URL")
	headers := http.Header{}
	headers.Set("x-hasura-admin-secret", adminSecret)

	client := graphql.NewClient(hasuraEndPoint, &http.Client{
		Transport: &headerTransport{Base: http.DefaultTransport, Headers: headers},
	})

	return client
}

type headerTransport struct {
	Base    http.RoundTripper
	Headers http.Header
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, values := range t.Headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	return t.Base.RoundTrip(req)
}
