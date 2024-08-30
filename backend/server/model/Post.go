package model

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"wordma/server/dto"
)

type Post struct {
	Slug     string `gorm:"index;size:255"`
	Up       int
	Down     int
	Read     int
	SiteID   uint      `gorm:"index"` // 外键，指向 Site
	Site     Site      `gorm:"foreignKey:SiteID"`
	Comments []Comment `gorm:"foreignKey:PostID"` // 定义一对多关系
	gorm.Model
}

// GetPostBySlug 根据slug获取post
func GetPostBySlug(slug string) (*Post, error) {
	var post Post
	err := DB.Limit(1).Where("slug = ?", slug).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// CreatePost 创建一条post
func CreatePost(data *Post) error {
	err := DB.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

// FindOrCreatePost 查找或创建
func FindOrCreatePost(data dto.CommentDTO) (*Post, error) {
	post, err := GetPostBySlug(data.PostSlug)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果没有这个文章，就创建
		post = &Post{
			Slug:   data.PostSlug,
			SiteID: data.SiteID,
		}
		if err := CreatePost(post); err != nil {
			return nil, fmt.Errorf("创建文章失败: %w", err)
		}
		// 再获取一次文章，保证数据最新
		post, err = GetPostBySlug(data.PostSlug)
		if err != nil {
			return nil, fmt.Errorf("创建文章后查询失败: %w", err)
		}
		return post, nil
	}
	// 如果有这个文章，直接返回就好了
	return post, nil
}
