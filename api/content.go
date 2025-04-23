/**
 * @Time : 2025/4/22 14:59
 * @File : content.go
 * @Software: share-text
 * @Author : Mr.Fang
 * @Description: 内容接口
 */

package api

import (
	"encoding/json"
	"io"
	"net/http"
	"share-text/handles"
	"share-text/models"
	"share-text/utils"
	"time"
)

func Content(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddContent(w, r)
	case http.MethodGet:
		GetContent(w, r)
	}
}

// AddContent 添加内容
func AddContent(w http.ResponseWriter, r *http.Request) {

	ip := utils.GetClientIP(r)
	var content models.ContentParams
	all, _ := io.ReadAll(r.Body)
	if len(all) > 0 {
		err := json.Unmarshal(all, &content)
		if err != nil {
			utils.Failure("参数错误", "").Write(w)
			return
		}
	}
	if len(content.Content) == 0 || len([]rune(content.Content)) > 500 {
		utils.Failure("内容不能为空或超过 500 字符", "").Write(w)
		return
	}

	pass := ""
	if content.Encrypt {
		pass = utils.GenerateRandomNumber(4)
	}
	now := time.Now()
	add := now.Add(time.Minute * 10)
	c := models.Content{
		Password:   pass,
		Content:    content.Content,
		IP:         ip,
		CreateTime: now.UnixMilli(),
		ExpireTime: add.UnixMilli(),
	}

	id := handles.CreateContent(c)
	if id > 0 {
		utils.Success("添加成功", struct {
			Pass string `json:"pass"`
		}{pass}).Write(w)
	} else {
		utils.Failure("失败", "").Write(w)
	}
}

// GetContent 获取内容
func GetContent(w http.ResponseWriter, r *http.Request) {
	ip := utils.GetIP(r)
	query := r.URL.Query()
	pass := query.Get("pass")
	if len(pass) > 0 {
		if content, err := handles.GetContent(pass); err != nil {
			utils.Success("查询结果", nil).Write(w)
		} else {
			utils.Success("查询结果", content).Write(w)
		}
	} else {
		if contents, err := handles.GetContents(ip); err != nil {
			utils.Success("查询结果", nil).Write(w)
		} else {
			utils.Success("查询结果", contents).Write(w)
		}
	}

}
