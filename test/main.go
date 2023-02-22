package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/jpeg"
	"os"
	"time"
)

func main() {
	video, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer video.Close()

	img := gocv.NewMat()
	defer img.Close()

	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	i := 0
	ticker := time.NewTicker(3 * time.Second)
	for {
		<-ticker.C
		i++
		if i > 10 {
			break
		}

		if ok := video.Read(&img); !ok {
			return
		}

		if img.Empty() {
			return
		}

		data, err := img.ToImage()
		if err != nil {
			fmt.Println("to img err:", err)
			return
		}

		file, err := os.Create("./test.jpg")
		if err != nil {
			return
		}
		defer file.Close()

		if err = jpeg.Encode(file, data, nil); err != nil {
			fmt.Println("jpeg err:", err)
		} else {
			fmt.Println("success")
		}
	}
}
