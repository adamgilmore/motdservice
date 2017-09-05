package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	httpTransport "github.com/go-kit/kit/transport/http"

	"context"
)

var messages = [...]string{"What you want fat boy?", "Howdy Doody!!!"}

func main() {
	svc := service{}

	messageHandler := httpTransport.NewServer(
		makeMessageEndpoint(svc),
		decodeMessageRequest,
		encodeResponse,
	)

	http.Handle("/", messageHandler)

	port := os.Getenv("MOTDSERVICE_PORT")
	if port == "" {
		port = "8091"
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// MOTDService returns messages
type MOTDService interface {
	Message(context.Context) (string, error)
}

type service struct{}

func (service) Message(_ context.Context) (string, error) {
	return messages[rand.Intn(len(messages))], nil
}

type messageResponse struct {
	Msg string `json:"msg"`
	Err string `json:"err,omitempty"`
}

func decodeMessageRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, res interface{}) error {
	return json.NewEncoder(w).Encode(res)
}

func makeMessageEndpoint(svc MOTDService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		msg, err := svc.Message(ctx)
		if err != nil {
			return messageResponse{msg, err.Error()}, nil
		}

		return messageResponse{msg, ""}, nil
	}
}
