package common

import (
	"crypto"
	randc "crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"hash"
	"math/rand"
	"time"
)

// GetReqTime 获取请求时间
func GetReqTime() string {
	return FormatTime(time.Now())
}

// FormatTime 格式化时间
func FormatTime(t time.Time) string {
	return t.Format("20060102150405")
}

// RandomString 随机生成字符串
func RandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandomPureString 随机生成纯字符串
func RandomPureString(l int) string {
	str := "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// RandomNumber 随机生成数字字符串
func RandomNumber(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// Sign 签名
func Sign(msg []byte, key *rsa.PrivateKey, hashType crypto.Hash, hash hash.Hash) (sign string, err error) {
	hash.Write(msg)
	hashed := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(randc.Reader, key, hashType, hashed)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(signature), nil
}

// ParseCertificate 解析证书
func ParseCertificate(cart []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(cart)
	if block == nil {
		return nil, NewErrMsg(InternalCode, "failed to parse certificate PEM")
	}

	return x509.ParseCertificate(block.Bytes)
}

// ParsePrivateKey 解析私钥
func ParsePrivateKey(pkcs8Key []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pkcs8Key)
	if block == nil {
		return nil, NewErrMsg(InternalCode, "failed to parse private key")
	}

	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, NewErrMsg(InternalCode, "failed to pkcs8Key to rsa.PrivateKey")
	}
	return privateKey, nil
}
