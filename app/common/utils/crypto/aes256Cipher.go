package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/eduwill/jarvis-api/app/common"
)

var (
	initialVector = "186409919b2f4311"
	secretKey     = []byte("186409919b2f4311bc592d11936a0397")
)

func Test() {
	var plainText = "usrNo=1072306|usrNm=안치순|usrId=chris83|progress=G|brnchNo=100|prgrNo=0|instNo=0|loginDt=20190905134723|systemCd=|loginIP=10.10.13.16|hisSeq=115535250"
	common.Logger.Debug("plainText : ", plainText)

	encryptedData, _ := AESEncrypt(plainText)
	common.Logger.Debug("encryptedData : ", encryptedData)

	decryptedText, _ := AESDecrypt(encryptedData)
	common.Logger.Debug("decryptedText : ", decryptedText)
}

func AESEncrypt(src string) (string, error) {
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		common.Logger.Error("key error1", err)
		return "", err
	}
	if src == "" {
		common.Logger.Warn("plain content empty")
		return "", nil
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	//base64 Encoding
	encriptText := base64.StdEncoding.EncodeToString(crypted)

	return encriptText, nil
}

func AESDecrypt(encriptText string) (string, error) {
	// Base64 Decoding
	crypt, _ := base64.StdEncoding.DecodeString(encriptText)

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		common.Logger.Error("key error1", err)
		return "", err
	}
	if len(crypt) == 0 {
		common.Logger.Warn("plain content empty")
		return "", nil
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return string(PKCS5Trimming(decrypted)), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
