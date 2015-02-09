package utility

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"io"
)

var commonIV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}

func GetUniHashString(content string) string {
	h := md5.New()
	io.WriteString(h, content)

	hString := hex.EncodeToString(h.Sum(nil))
	return hString
}

func GetSecretKey() string {
	return "astaxie12798akljzmknm.agggkljl;k"
}

func CryAse(origin string) (string, error) {
	plaintext := []byte(origin)
	c, err := aes.NewCipher([]byte(GetSecretKey()))
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(c, commonIV)
	ciphertext := make([]byte, len(plaintext))
	cfb.XORKeyStream(ciphertext, plaintext)
	return string(ciphertext), nil
}

func DeCryAse(cyedmsg string) (string, error) {
	cybyte := []byte(cyedmsg)
	c, err := aes.NewCipher([]byte(GetSecretKey()))
	if err != nil {
		return "", err
	}
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(cybyte))
	cfbdec.XORKeyStream(plaintextCopy, cybyte)
	return string(plaintextCopy), nil
}
