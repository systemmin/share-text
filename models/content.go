/**
 * @Time : 2025/4/22 14:26
 * @File : content.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description: 内容
 */

package models

// Content 表示分享内容表的结构体
type Content struct {
	ID         int    `json:"id"`
	Content    string `json:"content"`
	IP         string `json:"ip"`
	Password   string `json:"password"`
	ExpireTime int64  `json:"expire_time"`
	CreateTime int64  `json:"create_time"`
}
