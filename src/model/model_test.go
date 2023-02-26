package model

import (
	"fmt"
	"testing"

	tf "github.com/galeone/tensorflow/tensorflow/go"
)

//func TestLoadModel(t *testing.T) {
//	model := tg.LoadModel("C:\\Users\\zhang\\juypter_notebook_workplace\\model", []string{"serve"}, nil)
//	fmt.Println(model)
//}

func TestMnistModel(t *testing.T) {
	modelPath := "C:\\Users\\zhang\\workspace\\packvirus\\src\\model\\mnist_2"
	labelPath := "C:\\Users\\zhang\\workspace\\packvirus\\src\\model\\mnist\\label.txt"
	model, err := LoadModel(modelPath, labelPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(tf.Version())

	fakeInput, _ := tf.NewTensor([1][28][28][1]float32{})
	results := model.Net.Exec([]tf.Output{model.Net.Op("StatefulPartitionedCall", 0)}, map[tf.Output]*tf.Tensor{
		model.Net.Op("serving_default_input_1", 0): fakeInput,
	})

	fmt.Printf("%T\n", results[0].Value())
}
