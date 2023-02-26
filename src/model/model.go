package model

import (
	"bufio"
	"errors"
	tg "github.com/galeone/tfgo"
	"os"
)

type Model struct {
	Net    *tg.Model
	Labels []string
}

func LoadModel(path2model string, path2label string) (*Model, error) {
	m := new(Model)
	err := m.loadLabels(path2label)
	if err != nil {
		return nil, err
	}
	m.Net = tg.LoadModel(path2model, []string{"serve"}, nil)
	// set the default backend and target
	return m, nil
}

// PreprocessImage -
// The image preprocessing method should be consistent
func PreprocessImage() {}

func (m *Model) GetMaxProLocation(pro []float32) (maxLoc int, maxValue float32, err error) {
	if len(pro) < 0 || pro == nil {
		return -1, -1.0, errors.New("pro is nil")
	}

	maxLoc = 0
	maxValue = pro[0]

	for i := 0; i < len(pro); i++ {
		if pro[i] > maxValue {
			maxValue = pro[i]
			maxLoc = i
		}
	}

	return maxLoc, maxValue, nil
}

// loadLabels loads labels from path
func (m *Model) loadLabels(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	m.Labels = lines
	return scanner.Err()
}
