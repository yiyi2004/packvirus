package utils

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestReplace(t *testing.T) {
	res, err := Replace("C:\\Users\\zhang\\workspace\\packvirus\\test\\test.go.template", "{{ .algorithm }}", "3DES")
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("../tmp/main.go", []byte(res), 0666)
	if err != nil {
		fmt.Println(err)
	}

	cmd := exec.Command("go", "run", "../tmp/main.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
}
