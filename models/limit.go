/**
 * @Time : 2025/4/22 14:26
 * @File : limit.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description: 限流
 */

package models

// Limit 表示限流表的结构体
type Limit struct {
	ID         int    `json:"id"`
	APIAddress string `json:"api_address"`
	IP         string `json:"ip"`
	MethodText string `json:"method_text"`
	CreateTime int64  `json:"create_time"`
}
