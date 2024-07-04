package cms

import (
	"CRAZY/model"
	cmsServices "CRAZY/services/cms"
	"CRAZY/utils"
	"html"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostBannerForm struct {
	Status      uint   `form:"status" binding:"required"`
	Index       uint   `form:"index" binding:"required"`
	Title       string `form:"title" binding:"required"`
	Desc        string `form:"desc" binding:"required"`
	Image       string `form:"image" binding:"required"`
	Link        string `form:"link" binding:"required"`
	CategoryIds string `form:"categoryIds" binding:"required"`
}

// BannerBanner 新增文章
func PostBanner(c *gin.Context) {
	var form PostBannerForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Banner{
			Status:      form.Status,
			Index:       form.Index,
			Title:       html.EscapeString(form.Title),
			Desc:        html.EscapeString(form.Desc),
			Image:       html.EscapeString(form.Image),
			Link:        html.EscapeString(form.Link),
			CategoryIds: html.EscapeString(form.CategoryIds),
		}
		res, resErr := cmsServices.NewBannerService.Create(Model, form.CategoryIds)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

type DeleteBannerForm struct {
}

// DelBannerById 删除文章
func DelBannerById(c *gin.Context) {
	var form DeleteBannerForm
	err := c.ShouldBind(&form)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err == nil {
		var resErr = cmsServices.NewBannerService.DeleteById(id)
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
func PutBannerById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form PostBannerForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Banner{
			Status:      form.Status,
			Index:       form.Index,
			Title:       html.EscapeString(form.Title),
			Desc:        html.EscapeString(form.Desc),
			Image:       html.EscapeString(form.Image),
			Link:        html.EscapeString(form.Link),
			CategoryIds: html.EscapeString(form.CategoryIds),
		}
		res, resErr := cmsServices.NewBannerService.PutBannerById(id, Model, form.CategoryIds)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// GetBannerById 获取文章
func GetBannerById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form DeleteBannerForm
	err := c.ShouldBind(&form)
	if err == nil {
		res, resErr := cmsServices.NewBannerService.GetById(id)
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
func GetBanner(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	title := html.EscapeString(c.Query("title"))
	categoryIds := html.EscapeString(c.Query("categoryIds"))
	res, _ := cmsServices.NewBannerService.Get(page, pageSize, title, categoryIds)
	utils.OkDetailed(res, "success", c)
}
