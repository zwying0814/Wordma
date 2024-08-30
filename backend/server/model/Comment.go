package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	Content string
	UA      string
	IP      string
	Region  string
	Type    string `gorm:"default:published"`
	Up      int
	Down    int

	PostID uint `gorm:"index"` // 外键，指向 Post
	Post   Post `gorm:"foreignKey:PostID"`
	UserID uint `gorm:"index"` // 外键，指向 User
	User   User `gorm:"foreignKey:UserID"`

	Parent uint `gorm:"index;default:0"`
	gorm.Model
}

// CreateComment 新增评论
func CreateComment(data *Comment) error {
	err := DB.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

// GetCommentByID 根据ID查找评论
func GetCommentByID(id uint) (*Comment, error) {
	var comment Comment
	err := DB.First(&comment, id).Error
	if err != nil {
		return nil, err
	}
	return &comment, nil
}

// UpdateCommentByID 根据ID更新评论
func UpdateCommentByID(id uint, data *Comment) error {
	err := DB.Model(&Comment{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCommentByID 根据ID删除评论及其所有子评论
func DeleteCommentByID(id uint) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		// 删除所有子评论
		if err := tx.Where("parent = ?", id).Delete(&Comment{}).Error; err != nil {
			return err
		}

		// 删除当前评论
		if err := tx.Delete(&Comment{}, id).Error; err != nil {
			return err
		}

		return nil
	})
}
