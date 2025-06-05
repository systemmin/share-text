/**
 * @Time : 2025/4/22 16:14
 * @File : limit.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description:
 */

package handles

import (
	"fmt"
	"share-text/database"
	"share-text/models"
)

// LastLimit 获取当前 ip 最新一条内容
func LastLimit(limit models.Limit) models.Limit {
	db := database.DB
	row := db.QueryRow("SELECT id,api_address,ip,method_text,create_time FROM limits  WHERE ip = ? AND   api_address = ? AND method_text = ? ORDER BY create_time DESC limit 1", limit.IP, limit.APIAddress, limit.MethodText)
	var c models.Limit
	err := row.Scan(&c.ID, &c.APIAddress, &c.IP, &c.MethodText, &c.CreateTime)
	if err != nil {
		fmt.Println(err)
	}
	return c
}

// CreateLimit 添加限制
func CreateLimit(limit models.Limit) int64 {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	res, err := database.DB.Exec("INSERT INTO `limits` (api_address, ip,  method_text , create_time ) VALUES (?, ?,?,?)", limit.APIAddress, limit.IP, limit.MethodText, limit.CreateTime)
	if err != nil {
		fmt.Println(err)
	}
	id, _ := res.LastInsertId()
	return id
}
