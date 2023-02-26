package bypass

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"unsafe"

	"golang.org/x/sys/windows"
	"syscall"
	"time"
)

func CheckSandBox() bool {
	return CheckFile() || CheckMemory() || CheckNumberOfTempFiles() || CheckCoreNum() || CheckUILanguage() || CheckBootTime() || CheckVirtual()
}

func TimeSleep(t time.Duration) bool {
	startTime := time.Now()
	time.Sleep(t)
	period := time.Since(startTime)
	if period >= t {
		return true
	} else {
		return false
	}
}

// CheckUILanguage -
func CheckUILanguage() bool {
	languages, err := windows.GetUserPreferredUILanguages(windows.MUI_LANGUAGE_NAME)
	if err != nil {
		fmt.Println(err)
		return true
	}
	return languages[0] != "zh-CN"
}

// CheckBootTime check whether the boot time is greater than 30 minutes
func CheckBootTime() bool {
	var kernel = syscall.NewLazyDLL("Kernel32.dll")
	GetTickCount := kernel.NewProc("GetTickCount") // GetTickCount
	r, _, _ := GetTickCount.Call()
	if r == 0 {
		return true
	}
	ms := time.Duration(r * 1000 * 1000)
	tm := time.Duration(30 * time.Minute)
	if ms < tm {
		return true
	} else {
		return false
	}
}

// CheckMemory check whether the memory of host is greater than 4G
func CheckMemory() bool {
	var mod = syscall.NewLazyDLL("kernel32.dll")
	var proc = mod.NewProc("GetPhysicallyInstalledSystemMemory")
	var mem uint64

	proc.Call(uintptr(unsafe.Pointer(&mem))) // ret, _, err := proc.Call(uintptr(unsafe.Pointer(&mem)))

	mem = mem / 1048576
	if mem < 4 {
		return true
	}

	return false
}

// CheckCoreNum check the number of cpu cores is greater than 4
func CheckCoreNum() bool {
	coreNum := runtime.NumCPU()
	if coreNum < 4 {
		return true
	} else {
		return false
	}
}

// CheckNumberOfTempFiles check whether the number of temporary files is greater than 30
func CheckNumberOfTempFiles() bool {
	conn := os.Getenv("temp") // 通过环境变量读取temp文件夹路径
	var k int
	if conn == "" {
		//fmt.Println("未找到temp文件夹，或temp文件夹不存在")
		return true
	} else {
		local_dir := conn
		err := filepath.Walk(local_dir, func(filename string, fi os.FileInfo, err error) error {
			if fi.IsDir() {
				return nil
			}
			k++
			// fmt.Println("filename:", filename)  // 输出文件名字
			return nil
		})
		//fmt.Println("Temp总共文件数量:", k)
		if err != nil {
			// fmt.Println("路径获取错误")
			return true
		}
	}
	if k < 30 {
		return true
	}
	return false
}

// CheckVirtual -
func CheckVirtual() bool {
	model := ""
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/C", "wmic path Win32_ComputerSystem get Model")

	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		return false
	}

	model = strings.ToLower(string(stdout))
	if strings.Contains(model, "VirtualBox") || strings.Contains(model, "virtual") || strings.Contains(model, "VMware") ||
		strings.Contains(model, "KVM") || strings.Contains(model, "Bochs") || strings.Contains(model, "HVM domU") || strings.Contains(model, "Parallels") {
		return true
	}

	return false
}

func CheckFile() bool {
	if fack("C:\\windows\\System32\\Drivers\\Vmmouse.sys") ||
		fack("C:\\windows\\System32\\Drivers\\vmtray.dll") ||
		fack("C:\\windows\\System32\\Drivers\\VMToolsHook.dll") ||
		fack("C:\\windows\\System32\\Drivers\\vmmousever.dll") ||
		fack("C:\\windows\\System32\\Drivers\\vmhgfs.dll") ||
		fack("C:\\windows\\System32\\Drivers\\vmGuestLib.dll") ||
		fack("C:\\windows\\System32\\Drivers\\VBoxMouse.sys") ||
		fack("C:\\windows\\System32\\Drivers\\VBoxGuest.sys") ||
		fack("C:\\windows\\System32\\Drivers\\VBoxSF.sys") ||
		fack("C:\\windows\\System32\\Drivers\\VBoxVideo.sys") ||
		fack("C:\\windows\\System32\\vboxdisp.dll") ||
		fack("C:\\windows\\System32\\vboxhook.dll") ||
		fack("C:\\windows\\System32\\vboxoglerrorspu.dll") ||
		fack("C:\\windows\\System32\\vboxoglpassthroughspu.dll") ||
		fack("C:\\windows\\System32\\vboxservice.exe") ||
		fack("C:\\windows\\System32\\vboxtray.exe") ||
		fack("C:\\windows\\System32\\VBoxControl.exe") {
		return true
	}
	return false
}

func fack(path string) bool {
	b, _ := pathExists(path)
	if b {
		return true
	}
	return false
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
