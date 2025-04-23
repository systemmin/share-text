/**
 * @Time : 2025/4/22 15:23
 * @File : content_params.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description:
 */

package models

type ContentParams struct {
	Encrypt bool   `json:"encrypt"`
	Content string `json:"content"`
}
