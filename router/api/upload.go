package api

import (
	"CRAZY/utils"
	"CRAZY/utils/folder"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ResponseImg struct {
	Url  string `json:"url"`
	Alt  string `json:"alt"`
	Href string `json:"href"`
}

type ResponseUpload struct {
	Code int           `json:"errno"`
	Data []ResponseImg `json:"data"`
}

// Upload 文件上传 // TODO: 添加OSS方式上传
func Upload(c *gin.Context) {
	header, formErr := c.FormFile("file")
	if formErr == nil {
		arr := strings.Split(header.Filename, ".")
		str := arr[0] + strconv.FormatInt(time.Now().UnixNano(), 10) + "." + arr[1]
		if !folder.IsDir("./public/tmp/") {
			folder.CreateDir("./public/tmp/")
		}
		err := c.SaveUploadedFile(header, "./public/tmp/"+str)
		if err == nil {
			var res ResponseUpload
			res.Data = append(res.Data, ResponseImg{Url: str, Alt: "", Href: ""})
			c.JSON(0, res)
		} else {
			utils.FailWithMessage(err.Error(), c)
		}
	} else {
		utils.FailWithMessage("file 为空", c)
	}
}
