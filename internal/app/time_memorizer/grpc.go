package service

import (
	"context"
	"strconv"

	"github.com/radiance822/time_memorizer/internal/app/model"
)

type timeMemorizerServer struct {
	UnimplementedTimeMemorizerServer
	Storage *model.CategoryStorage
}

func (s timeMemorizerServer) AddTime(ctx context.Context, in *CategoryAndTime) (*Message, error) {
	s.Storage.Add(in.Category, int(in.TimeSpent))
	return &Message{
		Body: "Time added",
	}, nil
}

func (s timeMemorizerServer) SubstractTime(ctx context.Context, in *CategoryAndTime) (*Message, error) {
	s.Storage.Subtract(in.Category, int(in.TimeSpent))
	return &Message{
		Body: "Time substracted",
	}, nil
}

func (s timeMemorizerServer) ShowTime(ctx context.Context, in *Message) (*Message, error) {
	numb := s.Storage.Load(in.Body)
	return &Message{
		Body: strconv.Itoa(numb),
	}, nil
}
