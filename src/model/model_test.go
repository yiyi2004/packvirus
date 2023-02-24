package model

import (
	"fmt"
	tg "github.com/galeone/tfgo"
	"testing"
)

func TestLoadModel(t *testing.T) {
	model := tg.LoadModel("C:\\Users\\zhang\\juypter_notebook_workplace\\model", []string{"serve"}, nil)
	fmt.Println(model)
}
