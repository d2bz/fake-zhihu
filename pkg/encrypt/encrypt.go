package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/bcrypt"
)

var mobileKey = []byte("this-is-a-dev-key-32-bytes-long!")

// 密码永远不用明文存储，单向哈希即可
func EnPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err

	}
	return string(hashedPassword), nil
}

func CheckPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// 手机号加密
// 明文手机号 → AES-GCM加密（随机Nonce） → 合并Nonce+密文 → Base64编码 → 安全传输
func EnMobile(mobile string) (string, error) {
	// 创建AES加密块
	block, err := aes.NewCipher(mobileKey)
	if err != nil {
		return "", err
	}
	// 创建GCM模式，包装加密块，返回一个AEAD接口，用于加密和解密，同时保证数据机密性和完整性；自动生成认证标签
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	// 生成随机的Nonce，一次性随机数确保相同的明文加密后的密文不同
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}
	// 使用GCM模式进行加密，​输出：密文 + 认证标签
	cipherText := gcm.Seal(nil, nonce, []byte(mobile), nil)
	// 将Nonce和密文组合在一起，便于传输
	combined := append(nonce, cipherText...)
	// 返回Base64编码的加密结果，将二进制数据转为可安全传输的字符串（避免二进制处理问题）
	return base64.StdEncoding.EncodeToString(combined), nil
}

// 手机号解密
// Base64编码 → 解码Nonce+密文 → AES-GCM解密 → 明文手机号
func DeMobile(encryptedMobile string) (string, error) {
	// 解码Base64编码的加密数据
	combined, err := base64.StdEncoding.DecodeString(encryptedMobile)
	if err != nil {
		return "", err
	}
	// 创建AES加密块
	block, err := aes.NewCipher(mobileKey)
	if err != nil {
		return "", err
	}
	// 创建GCM模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	// 检查加密数据的长度是否有效
	if len(combined) < nonceSize {
		return "", errors.New("invalid encrypted data")
	}
	// 分离Nonce和密文
	nonce, ciphertext := combined[:nonceSize], combined[nonceSize:]
	// 使用GCM模式进行解密
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	// 返回解密后的明文
	return string(plaintext), nil
}
