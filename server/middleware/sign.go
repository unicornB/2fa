package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// 用于存储已使用的随机数
var usedNonces = make(map[string]bool)
var nonceMutex sync.Mutex

// 定义请求结构体
type RequestData struct {
	Data      map[string]interface{} `json:"data"`
	Sign      string                 `json:"sign"`
	Timestamp int64                  `json:"timestamp"`
	Nonce     string                 `json:"nonce"`
	Secret    string                 `json:"-"` // 不参与JSON解析
}

// 计算签名
func calculateSign(data map[string]interface{}, timestamp int64, nonce, secret string) string {
	keys := make([]string, 0, len(data))
	for k := range data {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var paramStr string
	for _, k := range keys {
		paramStr += fmt.Sprintf("%s=%v&", k, data[k])
	}
	paramStr += fmt.Sprintf("timestamp=%d&nonce=%s&secret=%s", timestamp, nonce, secret)

	hash := sha256.Sum256([]byte(paramStr))
	return hex.EncodeToString(hash[:])
}

// 验证签名
func verifySign(data map[string]interface{}, sign string, timestamp int64, nonce, secret string) bool {
	return calculateSign(data, timestamp, nonce, secret) == sign
}

// 检查时间戳是否在有效范围内
func isValidTimestamp(timestamp int64) bool {
	now := time.Now().Unix()
	// 假设有效时间范围为 60 秒
	return now-timestamp < 60 && now-timestamp > -60
}

// 检查随机数是否已使用
func isNonceUsed(nonce string) bool {
	nonceMutex.Lock()
	defer nonceMutex.Unlock()
	if usedNonces[nonce] {
		return true
	}
	usedNonces[nonce] = true
	return false
}

// 签名验证中间件
func SignatureMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RequestData
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		req.Secret = secret

		// 检查时间戳是否有效
		if !isValidTimestamp(req.Timestamp) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid timestamp"})
			c.Abort()
			return
		}

		// 检查随机数是否已使用
		if isNonceUsed(req.Nonce) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Nonce already used"})
			c.Abort()
			return
		}

		// 验证签名
		if !verifySign(req.Data, req.Sign, req.Timestamp, req.Nonce, req.Secret) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
			c.Abort()
			return
		}
		c.Next()
	}
}
