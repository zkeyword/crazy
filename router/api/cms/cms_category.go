package cms

import (
	"CRAZY/model"
	cmsServices "CRAZY/services/cms"
	"CRAZY/utils"
	"html"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostCategoryForm struct {
	Status uint   `form:"status" binding:"required"`
	Title  string `form:"title" binding:"required"`
	Index  uint   `form:"index" binding:"required"`
}

// CategoryCategory 新增文章
func PostCategory(c *gin.Context) {
	var form PostCategoryForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Category{
			Status: form.Status,
			Index:  form.Index,
			Title:  html.EscapeString(form.Title),
		}
		res, resErr := cmsServices.NewCategoryService.Create(Model)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

type DeleteCategoryForm struct {
}

// DelCategoryById 删除文章
func DelCategoryById(c *gin.Context) {
	var form DeleteCategoryForm
	err := c.ShouldBind(&form)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err == nil {
		var resErr = cmsServices.NewCategoryService.DeleteById(id)
		if resErr == nil {
			utils.Ok(c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// PutUserById 修改文章
func PutCategoryById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form PostCategoryForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Category{
			Status: form.Status,
			Index:  form.Index,
			Title:  html.EscapeString(form.Title),
		}
		res, resErr := cmsServices.NewCategoryService.PutCategoryById(id, Model)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// GetCategoryById 获取文章
func GetCategoryById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form DeleteCategoryForm
	err := c.ShouldBind(&form)
	if err == nil {
		res, resErr := cmsServices.NewCategoryService.GetById(id)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// GetUser 获取用户列表
func GetCategory(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	title := html.EscapeString(c.Query("title"))
	res, _ := cmsServices.NewCategoryService.Get(page, pageSize, title)
	utils.OkDetailed(res, "success", c)
}
