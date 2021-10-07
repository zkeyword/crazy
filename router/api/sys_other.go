package api

import (
	"CRAZY/model"
	otherService "CRAZY/services/sys_other"
	"CRAZY/utils"
	"html"

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
