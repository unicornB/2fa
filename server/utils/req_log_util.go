package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RequestData struct {
	Data map[string]interface{} `json:"data"`
}

func handleQueryParams(data map[string]interface{}, c *gin.Context) {
	for k, v := range c.Request.URL.Query() {
		if len(v) > 0 {
			data[k] = v[0]
		}
	}
}
func ReqLog(c *gin.Context) {
	start := time.Now()
	reqData := make(map[string]interface{})
	// 处理查询参数
	fmt.Println("Content-Type:" + c.Request.Header.Get("Content-Type"))
	handleQueryParams(reqData, c)
	// 处理不同请求方法和数据类型
	switch c.Request.Method {
	case http.MethodPost, http.MethodPut, http.MethodDelete:
		contentType := c.Request.Header.Get("Content-Type")
		if strings.Contains(contentType, "application/json") {
			// 处理 application/json
			var req RequestData

			decoder := json.NewDecoder(c.Request.Body)
			if err := decoder.Decode(&req.Data); err != nil {
				return
			}
			for k, v := range req.Data {
				reqData[k] = v
			}
			reqDataBytes, _ := json.Marshal(reqData)
			c.Request.Body = io.NopCloser(bytes.NewBuffer(reqDataBytes))
		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			for k, v := range c.Request.PostForm {
				if len(v) > 0 {
					reqData[k] = v[0]
				}
			}
		}
	}
	dataStr, _ := json.Marshal(reqData)
	c.Next()
	// 记录请求结束时间
	end := time.Now()
	latency := end.Sub(start)
	// 记录请求信息
	clientIP := c.ClientIP()
	method := c.Request.Method
	statusCode := c.Writer.Status()
	path := c.Request.URL.Path
	// 使用 logrus 记录日志
	logrus.WithFields(logrus.Fields{
		"status_code": statusCode,
		"latency":     latency,
		"client_ip":   clientIP,
		"method":      method,
		"path":        path,
		"data":        string(dataStr),
	}).Info("Request completed")

}
