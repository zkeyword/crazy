package api

import (
	"CRAZY/utils"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

type PersonInfo struct {
	Name    string   `json:"name"`
	Age     int32    `json:"age,string"`
	Sex     bool     `json:"sex"`
	Hobbies []string `json:"hobbies"`
}

// GetConfig 测试
func GetConfig(c *gin.Context) {
	personInfo := []PersonInfo{{"David", 30, true, []string{"跑步", "读书", "看电影"}}, {"Lee", 27, false, []string{"工作", "读书", "看电影"}}}
	schoolStr := `{"student":{"name":"jack","age":"20"},"teacher":"lucy"}`
	var result map[string]json.RawMessage
	if err := json.Unmarshal([]byte(schoolStr), &result); err != nil {
		fmt.Printf("json.Unmarshal schoolStr err: %v\n", err)
		return
	}
	utils.Result(200, gin.H{
		"message": "pong",
		"xx":      []string{"跑步", "读书", "看电影"},
		"xx2":     personInfo,
		"xx3":     result,
	}, "操作成功", c)
}
