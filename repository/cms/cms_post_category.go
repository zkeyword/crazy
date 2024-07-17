package cms

import (
	"CRAZY/model/cms"
	"CRAZY/utils/db"
)

type PostCategoryRepository struct {
}

// PostCategory 类型
type PostCategory struct {
	ID         uint
	PostID     string
	CategoryID string
}

// NewUserRepository 实例化 DAO
func NewPostCategoryRepository() *PostCategoryRepository {
	return &PostCategoryRepository{}
}

func (r *PostCategoryRepository) Create(postID uint, categoryID uint) (*cms.PostCategory, error) {
	var ret = new(cms.PostCategory)
	ret.PostID = postID
	ret.CategoryID = categoryID
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(ret).Error
	return ret, err
}

func (r *PostCategoryRepository) DeleteByPostId(id int64) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	return _db.Where("post_id = ?", id).Delete(PostCategory{}).Error
}
