package api


func NewManager(stg storage.Client, db db.Client, idGen utils.IDGen) Manager {
	return &manager{}
}

type manager struct {
}

func (m *manager) GetShortURLStatus(ctx context.Context, key string) (types.Job, error) {
	type Res struct {
		Key string
	}
	res := Res{
		Key : key,
	}
	return res, nil
}