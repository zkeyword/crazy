package cms

import (
	"CRAZY/model/cms"
	cmsServices "CRAZY/services/cms"
	"CRAZY/utils"
	"html"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PostPostForm struct {
	Title       string `form:"title" binding:"required"`
	Thumbnail   string `form:"thumbnail" binding:"required"`
	Desc        string `form:"desc" binding:"required"`
	Content     string `form:"content" binding:"required"`
	Status      uint   `form:"status" binding:"required"`
	Index       uint   `form:"index" binding:"required"`
	CategoryIds string `form:"categoryIds" binding:"required"`
	PublishAt   int    `form:"publishTime" binding:"required"`
}

// PostPost 新增文章
func PostPost(c *gin.Context) {
	var form PostPostForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &cms.Post{
			Title:       html.EscapeString(form.Title),
			Thumbnail:   html.EscapeString(form.Thumbnail),
			Content:     html.EscapeString(form.Content),
			Desc:        html.EscapeString(form.Desc),
			Status:      form.Status,
			Index:       form.Index,
			PublishAt:   form.PublishAt,
			CategoryIds: html.EscapeString(form.CategoryIds),
		}
		res, resErr := cmsServices.NewPostService.Create(Model, form.CategoryIds)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

type DeletePostForm struct {
}

// DelPostById 删除文章
func DelPostById(c *gin.Context) {
	var form DeletePostForm
	err := c.ShouldBind(&form)
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err == nil {
		var resErr = cmsServices.NewPostService.DeleteById(id)
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
func PutPostById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form PostPostForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &cms.Post{
			Title:       html.EscapeString(form.Title),
			Thumbnail:   html.EscapeString(form.Thumbnail),
			Content:     html.EscapeString(form.Content),
			Desc:        html.EscapeString(form.Desc),
			Status:      form.Status,
			Index:       form.Index,
			CategoryIds: html.EscapeString(form.CategoryIds),
		}
		res, resErr := cmsServices.NewPostService.PutPostById(id, Model, form.CategoryIds)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// GetPostById 获取文章
func GetPostById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form DeletePostForm
	err := c.ShouldBind(&form)
	if err == nil {
		res, resErr := cmsServices.NewPostService.GetById(id)
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
func GetPost(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	title := html.EscapeString(c.Query("title"))
	categoryIds := html.EscapeString(c.Query("categoryIds"))
	res, _ := cmsServices.NewPostService.Get(page, pageSize, title, categoryIds)
	utils.OkDetailed(res, "success", c)
}

func GetPostNews(c *gin.Context) {
	res, _ := cmsServices.NewPostService.GetNews()
	utils.OkDetailed(res, "success", c)
}

type ReqPostNews struct {
	Action string  `json:"action"`
	Ids    []int64 `json:"ids"`
}

const authToken = "tabi-blog-news"

// PutUserById 修改文章
func PostNewsSet(c *gin.Context) {
	token := c.DefaultQuery("token", "")
	if token != authToken {
		utils.FailWithMessage("token error", c)
		return
	}
	var req ReqPostNews
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.FailWithMessage(err.Error(), c)
		return
	}
	err := cmsServices.NewPostService.PostNewsSet(req.Action, req.Ids)
	if err == nil {
		utils.OkDetailed("", "success", c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
