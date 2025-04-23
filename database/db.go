/**
 * @Time : 2025/4/22 14:22
 * @File : db.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description: 数据库连接
 */

package database

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

// InitDB 初始化数据库
func InitDB() {

	var err error
	DB, err = sql.Open("sqlite", "./share.db")
	if err != nil {
		log.Fatal(err)
	}

	// 配置连接池参数
	//DB.SetMaxOpenConns(100)                // 设置最大打开连接数
	//DB.SetMaxIdleConns(20)                 // 设置最大空闲连接数
	//DB.SetConnMaxLifetime(5 * time.Minute) // 设置连接的最大生命周期

	// 测试连接
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 启用 WAL 模式
	_, err = DB.Exec("PRAGMA journal_mode=WAL")
	if err != nil {
		log.Fatal(err)
	}

	sql := `CREATE TABLE IF NOT EXISTS contents (
	id INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键自增
	content TEXT, -- 内容，字符串，长度无限制
	ip TEXT, -- IP 地址
	password TEXT CHECK(length(password) <= 10), -- 密码，字符串，长度不超过 10
	expire_time INTEGER, -- 失效时间
	create_time INTEGER -- 创建时间
	);`
	_, err = DB.Exec(sql)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}
	sql = `CREATE TABLE IF NOT EXISTS limits (
    id INTEGER PRIMARY KEY AUTOINCREMENT, -- 主键自增
    api_address TEXT, -- 接口地址，字符串，长度不超过 200
    ip TEXT, -- IP 地址
    method_text TEXT, 
    create_time INTEGER -- 创建时间
	);`
	_, err = DB.Exec(sql)
	if err != nil {
		log.Fatal("创建表失败:", err)
	}

}
