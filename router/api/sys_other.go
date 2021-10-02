package api

import (
	"CRAZY/model"
	otherService "CRAZY/services/sys_other"
	"CRAZY/utils"
	"html"
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
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	resErr := otherService.DeleteById(id)
	if resErr == nil {
		utils.Ok(c)
	} else {
		utils.FailWithMessage(resErr.Error(), c)
	}
}

func PutOtherById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
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
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res, _ := otherService.GetById(id)
	utils.OkDetailed(res, "success", c)
}
