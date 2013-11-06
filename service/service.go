package service

import (
	"net/http"
	"time"
	"github.com/rwl/go-endpoints/endpoints"
)

func init() {
	greetService := &GreetingService{}
	api, err := endpoints.RegisterService(greetService,
	"greeting", "v1", "Greetings API", true)
	if err != nil {
		panic(err.Error())
	}

	info := api.MethodByName("List").Info()
	info.Name, info.HttpMethod, info.Path, info.Desc =
	"greets.list", "GET", "greetings", "List most recent greetings."

	endpoints.HandleHttp()
}

// Greeting is a datastore entity that represents a single greeting.
// It also serves as (a part of) a response of GreetingService.
type Greeting struct {
	Key     string		`json:"id"`
	Author  string		`json:"author"`
	Content string		`json:"content"`
	Date    time.Time	`json:"date"`
}

// GreetingsList is a response type of GreetingService.List method
type GreetingsList struct {
	Items []*Greeting `json:"items"`
}

// Request type for GreetingService.List
type GreetingsListReq struct {
	Limit int `json:"limit" endpoints:"d=10"`
}

// GreetingService can sign the guesbook, list all greetings and delete
// a greeting from the guestbook.
type GreetingService struct {
}

// List responds with a list of all greetings ordered by Date field.
// Most recent greets come first.
func (gs *GreetingService) List(
	r *http.Request, req *GreetingsListReq, resp *GreetingsList) error {

		if req.Limit <= 0 {
			req.Limit = 10
		}

		greets := make([]*Greeting, 0, req.Limit)

		resp.Items = greets
		return nil
	}
