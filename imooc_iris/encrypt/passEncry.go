package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// 高级加密标准

//16 24 32位加密标准，分别对应AES-128 AES-192 AES-256的加密方法
var PwdKey = []byte("DIS**#KKKDJJSKDI")

func PKCS7Padding(ciphertext []byte, blocsize int) []byte {
	padding := blocsize - len(ciphertext)%blocsize
	// Repeat函数的功能是吧切片[]byte{byte(padding)}复制padding个，然后合并成新的字节切片返回
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

//填充的反向操作，删除填充字符串
func PKCS7UnPadding(origData []byte) ([]byte, error) {
	//获取数据的长度
	length := len(origData)
	if length == 0 {
		return nil, errors.New("bad")

	} else {
		unpadding := int(origData[length-1])
		//截取切片，删除填充的字节，并且返回明文
		return origData[:(length - unpadding)], nil
	}

}

func AesEcrypt(origData []byte, key []byte) ([]byte, error) {
	//创建加密算法
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//获取块大小
	blocksize := block.BlockSize()
	//对数据进行填充
	origData = PKCS7Padding(origData, blocksize)
	//采用AES加密方法中的CBC加密
	blockMode := cipher.NewCBCDecrypter(block, key[:blocksize])
	crypted := make([]byte, len(origData))
	//执行加密
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//解密
func AesDecrypt(cypted []byte, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blocksize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blocksize])
	origData := make([]byte, len(cypted))
	blockMode.CryptBlocks(origData, cypted)
	origData, err = PKCS7UnPadding(origData)
	if err != nil {
		return nil, err
	}
	return origData, nil
}

//加密base64
func EnPwdCode(pwd []byte) (string, error) {
	result, err := AesEcrypt(pwd, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(result), err

}

//解密
func DePwdCode(pwd string) ([]byte, error) {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return nil, err
	}
	//执行AES解密
	return AesDecrypt(pwdByte, PwdKey)
}
