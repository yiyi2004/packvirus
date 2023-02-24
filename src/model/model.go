package model

type Model interface {
	Exec(data interface{}) (interface{}, error)
	GetKey() string
}

type PredictM2MModel struct {
	originalKey string
	key         string
}

var _ Model = &PredictM2MModel{}

func (m *PredictM2MModel) Exec(data interface{}) (interface{}, error) {

	return "originalKey", nil
}

func (m *PredictM2MModel) GetKey() string {

	return m.key
}

func generateKey(originalKey string) (string, error) {

	return "key", nil
}
