package api

import(
	"context"
)

type Manager interface {
	GetShortURLStatus(ctx context.Context, key string) (Res, error)
}

func NewManager() Manager {
	return &manager{}
}

type manager struct {
}

type Res struct {
	Key string
}

func (m *manager) GetShortURLStatus(ctx context.Context, key string) (Res, error) {
	res := Res{
		Key : key,
	}
	return res, nil
}