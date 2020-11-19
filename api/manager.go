package api

type Manager interface {
	GetShortURLStatus(ctx context.Context, key string) (interface{}, error)
}

func NewManager() Manager {
	return &manager{}
}

type manager struct {
}

func (m *manager) GetShortURLStatus(ctx context.Context, key string) (interface{}, error) {
	type Res struct {
		Key string
	}
	res := Res{
		Key : key,
	}
	return res, nil
}