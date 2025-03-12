package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
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

// sortObjectString 对传入的 map 按键排序后进行 JSON 编码，并添加可选后缀
func sortObjectString(obj map[string]interface{}, opt string) string {
	// 获取所有键
	var keys []string
	for key := range obj {
		keys = append(keys, key)
	}
	// 对键进行排序
	sort.Strings(keys)

	// 按照排序后的键创建新的 map
	sortedObj := make(map[string]interface{})
	for _, key := range keys {
		sortedObj[key] = obj[key]
	}

	// 对排序后的 map 进行 JSON 编码
	jsonData, err := json.Marshal(sortedObj)
	if err != nil {
		// 处理 JSON 编码错误
		fmt.Println("JSON encoding error:", err)
		return ""
	}

	// 添加可选后缀
	result := string(jsonData) + opt
	return result
}
func filteredNullData(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		// 处理 map
		for key, value := range v {
			if value == nil {
				// 如果值为 nil，删除该键值对
				delete(v, key)
			} else {
				// 递归处理值
				v[key] = filteredNullData(value)
			}
		}
		return v
	case []interface{}:
		// 处理 slice
		var result []interface{}
		for _, value := range v {
			if value != nil {
				// 递归处理值
				result = append(result, filteredNullData(value))
			}
		}
		return result
	default:
		// 其他类型直接返回
		return v
	}
}

// 验证签名
func calculateSign(data map[string]interface{}, secret string) string {
	// 过滤空值
	filteredData := filteredNullData(data).(map[string]interface{})
	paramStr := sortObjectString(filteredData, secret)
	fmt.Println("paramStr:" + paramStr)
	hash := sha256.Sum256([]byte(paramStr))
	return hex.EncodeToString(hash[:])
}

// 验证签名
func verifySign(data map[string]interface{}, sign string, secret string) bool {
	backSign := calculateSign(data, secret)
	fmt.Println("backSign:" + backSign)
	fmt.Println("sign:" + strings.ToLower(sign))
	return calculateSign(data, secret) == strings.ToLower(sign)
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
func handleQueryParams(req *RequestData, c *gin.Context) {
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			req.Data[k] = v[0]
		}
	}
}

// 签名验证中间件
func SignatureMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req RequestData

		var err error
		req.Data = make(map[string]interface{})
		req.Sign = c.GetHeader("Sign")
		// 修改为从请求头获取 Timestamp
		timestampStr := c.GetHeader("Timestamp")
		req.Timestamp, _ = strconv.ParseInt(timestampStr, 10, 64)
		req.Data["timestamp"] = req.Timestamp
		// 修改为从请求头获取 Nonce
		req.Nonce = c.GetHeader("Nonce")
		fmt.Println("stringSignTemp:" + c.GetHeader("stringSignTemp"))
		req.Data["nonce"] = req.Nonce

		// 处理查询参数
		handleQueryParams(&req, c)

		// 处理不同请求方法和数据类型
		switch c.Request.Method {
		case http.MethodPost, http.MethodPut, http.MethodDelete:
			contentType := c.Request.Header.Get("Content-Type")
			if strings.Contains(contentType, "application/json") {
				// 处理 application/json
				decoder := json.NewDecoder(c.Request.Body)
				if err = decoder.Decode(&req.Data); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					c.Abort()
					return
				}
			} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
				// 处理 application/x-www-form-urlencoded
				if err = c.Request.ParseForm(); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					c.Abort()
					return
				}
				for k, v := range c.Request.PostForm {
					if len(v) > 0 {
						req.Data[k] = v[0]
					}
				}
			}
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
		if !verifySign(req.Data, req.Sign, req.Secret) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid signature"})
			c.Abort()
			return
		}
		fmt.Println("签名认证成功")
		c.Next()
	}
}
