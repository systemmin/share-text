/**
 * @Time : 2025/4/23 10:20
 * @File : timer.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description: 定时任务
 */

package job

import (
	"log"
	"share-text/handles"
	"time"
)

func DeleteTimeoutData() {
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop() // 确保在程序退出时停止 Ticker
		for {
			select {
			case _ = <-ticker.C: // 接收时间戳
				content := handles.DeleteTimeoutContent()
				if content > 0 {
					log.Println("删除行数", content)
				}
			}
		}
	}()
}
