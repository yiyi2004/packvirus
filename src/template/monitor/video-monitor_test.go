package monitor

import (
	"sync"
	"testing"
	"time"
)

var wg *sync.WaitGroup

func TestVideoMonitor(t *testing.T) {
	ticker := time.NewTicker(500 * time.Millisecond)
	modelPath := "C:\\Users\\zhang\\workspace\\packvirus\\src\\model\\mnist"
	labelPath := "C:\\Users\\zhang\\workspace\\packvirus\\src\\model\\mnist\\label.txt"
	keyChan := make(chan string)
	exitChan := make(chan bool)

	wg.Add(1)
	go VideoMonitor(wg, ticker, modelPath, labelPath, keyChan, exitChan)
	wg.Wait()

}
