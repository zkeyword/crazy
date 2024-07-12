package system

import (
	"CRAZY/model"
	otherService "CRAZY/services/sys/sys_other"
	"CRAZY/utils"
	"encoding/csv"
	"fmt"
	"html"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OtherForm struct {
	Key   string `form:"key" binding:"required"`
	Value string `form:"value" binding:"required"`
}

// PostOther 新增
func PostOther(c *gin.Context) {
	var form OtherForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Other{
			Key:   html.EscapeString(form.Key),
			Value: html.EscapeString(form.Value),
		}
		res, resErr := otherService.Create(Model)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

func DelOtherById(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	resErr := otherService.DeleteById(id)
	if resErr == nil {
		utils.Ok(c)
	} else {
		utils.FailWithMessage(resErr.Error(), c)
	}
}

func PutOtherById(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	var form OtherForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Other{
			Key:   html.EscapeString(form.Key),
			Value: html.EscapeString(form.Value),
		}
		res, resErr := otherService.PutById(id, Model)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

func GetOther(c *gin.Context) {
	res, _ := otherService.Get()
	utils.OkDetailed(res, "success", c)
}

func GetOtherById(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	res, _ := otherService.GetById(id)
	utils.OkDetailed(res, "success", c)
}

func ImportOther(c *gin.Context) {
	// 获取上传的文件
	formFile, err := c.FormFile("file")
	if err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}

	// 创建一个临时文件来保存上传的 CSV 文件
	tempFile, err := os.CreateTemp("", "upload-*.csv")
	if err != nil {
		utils.FailWithMessage(fmt.Sprintf("create temp file err: %s", err.Error()), c)
		return
	}
	defer os.Remove(tempFile.Name()) // 在函数返回后删除临时文件

	// 将上传的文件保存到临时文件中
	err = c.SaveUploadedFile(formFile, tempFile.Name())
	if err != nil {
		utils.FailWithMessage(fmt.Sprintf("save uploaded file err: %s", err.Error()), c)
		return
	}

	// 打开临时文件并创建一个 csv.Reader
	f, err := os.Open(tempFile.Name())
	if err != nil {
		utils.FailWithMessage(fmt.Sprintf("open file err: %s", err.Error()), c)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	var others []*model.Other
	index := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			utils.FailWithMessage(err.Error(), c)
			return
		}

		if index == 0 || len(record) == 0 {
			index++
			continue
		}

		// 创建一个新的 Other 对象
		temp, _ := strconv.Atoi(record[2])
		other := &model.Other{
			Key:   html.EscapeString(record[0]),
			Value: html.EscapeString(record[1]),
			Type:  uint(temp),
		}
		others = append(others, other)
	}
	res, resErr := otherService.BatchCreate(others)
	if resErr != nil {
		utils.FailWithMessage(resErr.Error(), c)
		return
	}
	utils.OkDetailed(res, "success", c)
}

func ExportOther(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment;filename=other.csv")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Disposition")

	wr := csv.NewWriter(c.Writer)
	data := []string{"key", "value", "type"}
	wr.Write(data)
	res, _ := otherService.Get()
	for _, v := range res {
		data := []string{v.Key, v.Value, strconv.Itoa(int(v.Type))}
		wr.Write(data)
	}
	wr.Flush()
}
