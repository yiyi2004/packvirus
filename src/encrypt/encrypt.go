package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

var EncryptFunctions map[string]func(data []byte, key []byte) ([]byte, error)
var DecryptFunctions map[string]func(data []byte, key []byte) ([]byte, error)

func init() {
	EncryptFunctions = make(map[string]func(data []byte, key []byte) ([]byte, error))
	EncryptFunctions["aes"] = EncryptAES

	DecryptFunctions = make(map[string]func(data []byte, key []byte) ([]byte, error))
	DecryptFunctions["aes"] = DecryptAES
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func PaddingText(str []byte, blockSize int) []byte {
	paddingCount := blockSize - len(str)%blockSize
	paddingStr := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	newPaddingStr := append(str, paddingStr...)

	return newPaddingStr
}

// EncryptAES -
func EncryptAES(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil, err
	}

	src = PaddingText(src, block.BlockSize())

	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)

	return src, nil
}

func UnPaddingText(str []byte) []byte {
	n := len(str)
	count := int(str[n-1])
	newPaddingText := str[:n-count]

	return newPaddingText
}

func DecryptAES(src, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(nil)
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = UnPaddingText(src)
	return src, nil
}
