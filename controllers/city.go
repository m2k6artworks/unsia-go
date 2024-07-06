package controllers

import (
	"context"
	"unsia/pb/cities"
)

// City struct
type City struct{}

// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.City) (*cities.City, error) {
	return &cities.City{Id: 8, Name: "Solo/Surakarta"}, nil
}
