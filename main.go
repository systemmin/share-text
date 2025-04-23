/**
 * @Time : 2025/4/22 14:22
 * @File : main.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description: 共享剪切板应用
 */

package main

import (
	"embed"
	"flag"
	"fmt"
	"net/http"
	"share-text/api"
	"share-text/database"
	"share-text/handles"
	"share-text/job"
)

// 嵌入声明

//go:embed static/*
var content embed.FS

var port string

func init() {

	flag.StringVar(&port, "p", "9999", "指定端口 0-65535")
	// 解析命令行参数
	flag.Parse()

	// 初始化数据库
	database.InitDB()

	// 启动定时任务
	job.DeleteTimeoutData()

}

func main() {

	// 创建一个新的 mux 路由器
	mux := http.NewServeMux()

	// 静态文件处理器
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := content.ReadFile("static/index.html")
		if err != nil {
			http.Error(w, "无法加载页面："+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(data)
	})
	mux.HandleFunc("/logo.png", func(w http.ResponseWriter, r *http.Request) {
		data, err := content.ReadFile("static/logo.png")
		if err != nil {
			return
		}
		w.Header().Set("Content-Type", "image/png")
		w.Write(data)
	})
	mux.HandleFunc("/content", func(writer http.ResponseWriter, request *http.Request) {
		api.Content(writer, request)
	})

	fmt.Printf("http://127.0.0.1:%s\n", port)

	// 应用中间件
	handler := handles.Middleware(mux)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), handler); err != nil {
		fmt.Println("无法启动", err)
	}

}
