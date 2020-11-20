package api

//go:generate mockery -name Manager -outpkg apimocks -output ./apimocks -dir .
type Manager interface {
	GetShortURLStatus(key string) (Res, error)
}

func NewManager() Manager {
	return &manager{}
}

type manager struct {
}

type Res struct {
	Key string
}

func (m *manager) GetShortURLStatus(key string) (Res, error) {
	res := Res{
		Key : key,
	}
	return res, nil
}