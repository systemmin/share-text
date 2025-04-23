/**
 * @Time : 2025/4/22 16:06
 * @File : utils.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description:
 */

package utils

import (
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

// GenerateRandomNumber 生成指定长度的随机数
func GenerateRandomNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = digits[rand.Intn(len(digits))]
	}
	return string(result)
}

func GetIP(r *http.Request) string {
	header := r.Header
	headerIP := []string{"x-natapp-ip", "remote-host", "x-forwarded-for", "x-real-ip"}
	for k, v := range header {
		for _, key := range headerIP {
			if strings.ToLower(k) == key {
				return v[0]
			}
		}
	}
	addr := r.RemoteAddr
	if strings.Contains(addr, "[::1]") {
		return "127.0.0.1"
	}
	return strings.Split(addr, ":")[0]
}

// GetClientIP 返回客户端 IP，适配反向代理、本地调试等常见场景
func GetClientIP(r *http.Request) string {
	// 优先获取 X-Forwarded-For（可能是多个 IP）
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// 然后尝试 X-Real-IP
	if xrip := r.Header.Get("X-Real-Ip"); xrip != "" {
		return strings.TrimSpace(xrip)
	}

	// fallback 到 RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}

	// IPv6 本地地址处理
	if ip == "::1" {
		return "127.0.0.1"
	}

	return ip
}
