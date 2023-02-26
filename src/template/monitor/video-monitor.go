package monitor

import (
	"fmt"
	tf "github.com/galeone/tensorflow/tensorflow/go"
	"gocv.io/x/gocv"
	"image"
	"packvirus/model"
	"packvirus/utils"
	"strconv"
	"sync"
	"time"
)

// VideoMonitor -
func VideoMonitor(wg *sync.WaitGroup, ticker *time.Ticker, modelPath string, labelPath string, keyChan chan<- string, exitChan <-chan bool) {
	defer wg.Done()

	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println("Error opening video capture device")
		return
	}
	defer webcam.Close()

	vmodel, err := model.LoadModel(modelPath, labelPath)
	if err != nil {
		fmt.Println("Error loading model")
	}

	fmt.Println(vmodel)
	img := gocv.NewMat()
	defer img.Close()

	// test
	window := gocv.NewWindow("Mnist Classifier")
	defer window.Close()
	//i := 0

	for {

		//if i > 1000 {
		//	exitChan <- true
		//	close(exitChan)
		//	break
		//}
		//i++

		select {
		case <-ticker.C:
			if ok := webcam.Read(&img); !ok {
				fmt.Println("device closed")
				return
			}
			if img.Empty() {
				continue
			}

			gocv.CvtColor(img, &img, gocv.ColorBGRAToGray)

			imgConv := img.Clone()
			//imgConv.ConvertTo()
			defer imgConv.Close()

			blob := gocv.BlobFromImage(imgConv, 1/255.0, image.Pt(28, 28), gocv.NewScalar(0, 0, 0, 0), true, false)
			defer blob.Close()

			m, err := model.LoadModel(modelPath, labelPath)
			if err != nil {
				fmt.Println("load model failed")
				return
			}

			fmt.Printf("blob Size %v\n", blob.Size())
			fmt.Printf("blob Channels %v\n", blob.Channels())
			fmt.Printf("blob Cols %v\n", blob.Cols())
			fmt.Printf("blob IsContinuous %v\n", blob.IsContinuous())

			ptrFloat32, err := blob.DataPtrFloat32()
			if err != nil {
				fmt.Println(err)
				return
			}

			input, err := tf.NewTensor(ptrFloat32)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("input shape %v\n", input.Shape())
			fmt.Printf("input DataType %v\n", input.DataType())
			fmt.Printf("input Value %v\n", input.Value())

			// unable to convert shape [2352] (num_elements: 2352) into shape [-1 28 28 1] (num_elements: -784)
			//err = input.Reshape([]int64{-1, 28, 28, 1})
			//if err != nil {
			//	fmt.Println(err)
			//	return
			//}

			fmt.Printf("input2 shape %v\n", input.Shape())
			fmt.Printf("input2 DataType %v\n", input.DataType())
			fmt.Printf("input2 Value %v\n", input.Value())

			// A model exported with tf.saved_model.save()
			// automatically comes with the "serve" tag because the SavedModel
			// file format is designed for serving.
			// This tag contains the various functions exported. Among these, there is
			// always present the "serving_default" signature_def. This signature def
			// works exactly like the TF 1.x graph. Get the input tensor and the output tensor,
			// and use them as placeholder to feed and output to get, respectively.

			// To get info inside a SavedModel the best tool is saved_model_cli
			// that comes with the TensorFlow Python package.

			// e.g. saved_model_cli show --all --dir output/keras
			// gives, among the others, this info:

			// signature_def['serving_default']:
			// The given SavedModel SignatureDef contains the following input(s):
			//   inputs['inputs_input'] tensor_info:
			//       dtype: DT_FLOAT
			//       shape: (-1, 28, 28, 1)
			//       name: serving_default_inputs_input:0
			// The given SavedModel SignatureDef contaicons the following output(s):
			//   outputs['logits'] tensor_info:
			//       dtype: DT_FLOAT
			//       shape: (-1, 10)
			//       name: StatefulPartitionedCall:0
			// Method name is: tensorflow/serving/predict

			//	MetaGraphDef with tag-set: 'serve' contains the following SignatureDefs:
			//
			//	signature_def['__saved_model_init_op']:
			//	The given SavedModel SignatureDef contains the following input(s):
			//	The given SavedModel SignatureDef contains the following output(s):
			//	outputs['__saved_model_init_op'] tensor_info:
			//dtype: DT_INVALID
			//shape: unknown_rank
			//name: NoOp
			//	Method name is:
			//
			//	signature_def['serving_default']:
			//	The given SavedModel SignatureDef contains the following input(s):
			//	inputs['input_1'] tensor_info:
			//dtype: DT_FLOAT
			//shape: (-1, 28, 28)
			//name: serving_default_input_1:0
			//	The given SavedModel SignatureDef contains the following output(s):
			//	outputs['output_1'] tensor_info:
			//dtype: DT_FLOAT
			//shape: (-1, 10)
			//name: StatefulPartitionedCall:0

			err = input.Reshape([]int64{28, 28, 1})
			if err != nil {
				fmt.Println(err)
				return
			}
			//fakeInput, _ := tf.NewTensor([1][28][28][1]float32{})
			//fmt.Println(fakeInput)
			results := m.Net.Exec([]tf.Output{m.Net.Op("StatefulPartitionedCall", 0)}, map[tf.Output]*tf.Tensor{m.Net.Op("serving_default_input_1", 0): input})

			// Close
			fmt.Printf("res %v\n", results[0].Value())
			fmt.Printf("res Type %T\n", results[0].Value())

			pro := results[0].Value().([][]float32)

			loc, _, err := m.GetMaxProLocation(pro[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			//binaryString := utils.Convert2BinaryString(loc)
			//padding the binary key to 16 bits
			//hash the malware
			fmt.Printf("loc %v\n", loc)
			fmt.Printf("value %v\n", m.Labels[loc])
			result := procMnistResult(m.Labels[loc])
			keyChan <- result

			window.IMShow(img)
			if window.WaitKey(1) >= 0 {
				break
			}
		case <-exitChan:
			close(keyChan)
			return
		}
	}
}

// procResult handles the model output
func procMnistResult(res string) string {
	intRes, err := strconv.Atoi(res)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	binaryString := utils.Convert2BinaryString(intRes)
	binaryString = binaryString + binaryString + binaryString + binaryString
	fmt.Printf("key value %s\n", binaryString)
	fmt.Printf("key len %d\n", len(binaryString))

	return binaryString
}
