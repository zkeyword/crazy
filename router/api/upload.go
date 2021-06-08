package api

import (
	"CRAZY/utils"

	"github.com/gin-gonic/gin"
)

// Upload 文件上传 // TODO: 添加OSS方式上传
func Upload(c *gin.Context) {
	header, formErr := c.FormFile("file")
	if formErr == nil {
		err := c.SaveUploadedFile(header, "./public/tmp/"+header.Filename) // TODO:确保文件名唯一
		if err == nil {
			utils.OkDetailed(gin.H{
				"fileName": header.Filename,
			}, "success", c)
		} else {
			utils.FailWithMessage("上传失败", c)
		}
	} else {
		utils.FailWithMessage("file 为空", c)
	}
}
