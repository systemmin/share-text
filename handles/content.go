/**
 * @Time : 2025/4/22 14:27
 * @File : content.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description:
 */

package handles

import (
	"fmt"
	"share-text/database"
	"share-text/models"
	"sync"
	"time"
)

var dbMutex sync.Mutex

// GetContents 获取当前 ip 内容
func GetContents(ip string) ([]models.Content, error) {
	rows, err := database.DB.Query("SELECT id, content, ip,`password`, expire_time, create_time FROM contents WHERE ip = ? AND password = ? ORDER BY create_time DESC", ip, "")
	var contents []models.Content

	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return contents, err
	}

	for rows.Next() {
		var c models.Content
		rows.Scan(&c.ID, &c.Content, &c.IP, &c.Password, &c.ExpireTime, &c.CreateTime)
		contents = append(contents, c)
	}

	return contents, nil
}

// GetContent pass 内容
func GetContent(pass string) (models.Content, error) {
	row := database.DB.QueryRow("SELECT id, content, ip,`password`, expire_time, create_time FROM contents WHERE password = ? ORDER BY create_time DESC", pass)
	var c models.Content
	err := row.Scan(&c.ID, &c.Content, &c.IP, &c.Password, &c.ExpireTime, &c.CreateTime)
	if err != nil {
		fmt.Println(err)
		return c, err
	}
	// 通过密码获取内容直接删除
	if c.ID != 0 {
		DeleteContent(c.ID)
	}
	return c, nil
}

// CreateContent 创建内容
func CreateContent(content models.Content) int64 {
	dbMutex.Lock()
	defer dbMutex.Unlock()
	db := database.DB
	res, _ := db.Exec("INSERT INTO `contents` (`content`, ip,  `password` , expire_time , create_time) VALUES (?, ?,?,?,?)", content.Content, content.IP, content.Password, content.ExpireTime, content.CreateTime)
	id, _ := res.LastInsertId()
	return id
}

// DeleteContent 通过 id 删除内容
func DeleteContent(id int) {
	_, err := database.DB.Exec("DELETE FROM contents WHERE id = ?", id)
	if err != nil {
		fmt.Println("删除失败", err)
	}
}

// DeleteTimeoutContent 超时内容
func DeleteTimeoutContent() int64 {
	res, err := database.DB.Exec("DELETE FROM contents WHERE expire_time < ?", time.Now().UnixMilli())
	if err != nil {
		fmt.Println("删除失败", err)
	}
	affected, _ := res.RowsAffected()
	return affected
}
