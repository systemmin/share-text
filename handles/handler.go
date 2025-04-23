/**
 * @Time : 2025/4/23 10:19
 * @File : handler.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description:
 */

package handles

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"share-text/models"
	"share-text/utils"
	"time"
)

// Middleware 中间件，请求拦截处理
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		ip := utils.GetClientIP(r)
		path := r.URL.Path
		if path != "/" {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.Header().Set("Content-Type", "application/json")
		}

		log.Printf("IP :%s ", ip)
		log.Printf("Started %s %s", r.Method, path)
		header := r.Header
		data := make(map[string]string)
		for k, v := range header {
			data[k] = v[0]
		}
		data["url"] = path
		data["ip"] = ip
		marshal, _ := json.Marshal(data)
		log.Printf("headers: %s", marshal)

		milli := start.UnixMilli()

		limit := models.Limit{
			IP:         ip,
			APIAddress: path,
			MethodText: r.Method,
			CreateTime: milli,
		}

		if r.Method == http.MethodOptions {
			fmt.Println("options")
			// options 直接返回 200
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
			return
		} else { // 限流 2 秒
			lastLimit := LastLimit(limit)
			if lastLimit.ID != 0 {
				t := milli - lastLimit.CreateTime
				milliseconds := 2 * 1000
				if t < int64(milliseconds) && path != "/" && r.Method != http.MethodGet {
					utils.Failure("触发2秒限流", "").Write(w)
					return
				}
			}
			CreateLimit(limit)
		}
		next.ServeHTTP(w, r)
	})
}
