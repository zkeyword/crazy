package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
	"errors"
	"strings"
)

type PostRepository struct {
}

// Post 类型
type Post struct {
	ID        uint
	Title     string
	Desc      string
	Thumbnail string
	Content   string
	Index     uint
}

// NewUserRepository 实例化 DAO
func NewPostRepository() *PostRepository {
	return &PostRepository{}
}

// Create 创建文章
func (r *PostRepository) Create(t *model.Post) (*model.Post, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

// DeleteById 删除删除
func (r *PostRepository) DeleteById(id int64) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	return _db.Where("id = ?", id).Delete(Post{}).Error
}

// UpdateById 修改文章
func (r *PostRepository) UpdateById(id int64, t *model.Post) (*model.Post, error) {
	var ret = new(model.Post)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// GetById 获取文章
func (r *PostRepository) GetById(id int64) (*model.Post, error) {
	var ret = &model.Post{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.First(ret, "id = ?", id).Error
	return ret, err
}

// Get 获取文章
func (r *PostRepository) Get(page int, pageSize int, title string, categoryIds string) ([]model.Post, error) {
	var Post []model.Post
	var PostCategory []model.PostCategory
	var err error
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	if title != "" {
		err = _db.Order("`index` DESC, `updated_at` ASC").Where("title like ?", "%"+title+"%").Limit(pageSize).Offset((page - 1) * 10).Find(&Post).Error
	} else {
		ids := strings.Split(categoryIds, ",")
		if categoryIds != "" && len(ids) > 0 {
			err = _db.Where("post_id IN (?)", ids).Find(&PostCategory).Error
			if err == nil {
				var postIDs []uint
				for _, v := range PostCategory {
					postIDs = append(postIDs, v.PostID)
				}
				err = _db.Order("`index` DESC, `updated_at` ASC").Limit(pageSize).Offset((page-1)*10).Where("id IN (?)", postIDs).Find(&Post).Error
			}
		} else {
			err = _db.Order("`index` DESC, `updated_at` ASC").Limit(pageSize).Offset((page - 1) * 10).Find(&Post).Error
		}
	}
	return Post, err
}

func (r *PostRepository) GetPostCount(title string) int64 {
	var users []model.Post
	var count int64
	_db, err := db.GetMysql()
	if err != nil {
		return 0
	}
	if title != "" {
		_db.Where("title like ?", "%"+title+"%").Find(&users).Select("count(id)").Count(&count)
	} else {
		_db.Find(&users).Select("count(id)").Count(&count)
	}
	return count
}

// GetNews 获取文章
func (r *PostRepository) GetNews() ([]model.Post, error) {
	var Post []model.Post
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Select("id,title,`desc`,thumbnail,`status`,`index`,category_ids,news_rank,publish_at").Where("news_rank >?", 0).Order("news_rank ASC").Find(&Post).Error
	return Post, err
}

func (r *PostRepository) PostNewsSet(action string, ids []int64) error {
	var err error
	if len(ids) == 0 {
		return errors.New("no news change")
	}
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	if action == "add" {
		err = _db.Model(&model.Post{}).Where("id=?", ids[0]).Update("news_rank", 10).Error
	} else {
		err = _db.Model(&model.Post{}).Where("news_rank>?", 0).Update("news_rank", 0).Error
		for k, v := range ids {
			err = _db.Model(&model.Post{}).Where("id=?", v).Update("news_rank", k+1).Error
		}
	}
	return err
}
