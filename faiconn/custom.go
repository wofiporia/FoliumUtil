package faiconn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// 自定义连接器OPENAI API兼容API
type CustomConn struct {
	config AIConfig
	client *http.Client
}

// 创建自定义连接
func NewCustomConn(config AIConfig) (*CustomConn, error) {
	return &CustomConn{
		config: config,
		client: &http.Client{Timeout: 30 * time.Second},
	}, nil
}

// 发送消息 - 直接使用OpenAI格式
func (c *CustomConn) SendMessage(message string) (string, error) {
	req := map[string]interface{}{
		"model": c.config.Model,
		"messages": []map[string]string{
			{"role": "user", "content": message},
		},
	}

	return c.sendRequest(req)
}

// 发送请求
func (c *CustomConn) sendRequest(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 构建完整URL
	fullURL := c.buildFullURL()

	httpReq, err := http.NewRequest("POST", fullURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	httpReq.Header.Set("Content-Type", "application/json")
	if c.config.APIKey != "" {
		httpReq.Header.Set("Authorization", "Bearer "+c.config.APIKey)
	}

	resp, err := c.client.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("HTTP %d: %s", resp.StatusCode, string(body))
	}

	// 解析OpenAI格式的响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return string(body), nil
	}

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := choice["message"].(map[string]interface{}); ok {
				if content, ok := message["content"].(string); ok {
					return content, nil
				}
			}
		}
	}

	return string(body), nil
}

// 获取模型信息
func (c *CustomConn) GetModel() string {
	return c.config.Model
}

// 获取提供商信息
func (c *CustomConn) GetProvider() string {
	return c.config.Provider
}

// 关闭连接
func (c *CustomConn) Close() error {
	return nil
}

// 构建API URL - 默认添加OpenAI兼容端点
func (c *CustomConn) buildFullURL() string {
	if strings.HasSuffix(c.config.URL, "/chat/completions") {
		return c.config.URL
	}
	return strings.TrimSuffix(c.config.URL, "/") + "/chat/completions"
}
