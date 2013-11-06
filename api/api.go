package api

import (
	"net/url"
	"github.com/rwl/endpoint"
)

const defaultURL = "http://localhost:8080"

func init() {
	u, _ := url.Parse(defaultURL)
	e := endpoint.NewEndpointsServer(u)
	e.HandleHttp(nil)
}
