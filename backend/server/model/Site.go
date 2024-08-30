package model

import (
	"gorm.io/gorm"
)

type Site struct {
	Url   string `gorm:"uniqueIndex;size:255"` // 站点URL
	Name  string
	Posts []Post `gorm:"foreignKey:SiteID"` // 站点下的文章
	gorm.Model
}

// CreateSite 新增站点
func CreateSite(data *Site) error {
	err := DB.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

// FindSiteByID 根据id查找站点
func FindSiteByID(id uint) (*Site, error) {
	var site Site
	err := DB.Limit(1).Where("id = ?", id).First(&site).Error
	if err != nil {
		return nil, err
	}
	return &site, nil
}

// FindSiteByURL 根据url查找站点
func FindSiteByURL(url string) (*Site, error) {
	var site Site
	err := DB.Limit(1).Where("url = ?", url).First(&site).Error
	if err != nil {
		return nil, err
	}
	return &site, nil
}

// DeleteSiteById 根据Id删除站点
func DeleteSiteById(id uint) error {
	return DB.Delete(&Site{}, id).Error
}

// UpdateSiteById 根据ID修改站点信息
func UpdateSiteById(id uint, site *Site) error {
	return DB.Model(&Site{}).Where("id = ?", id).Updates(site).Error
}
