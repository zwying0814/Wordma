package model

import (
	"crypto/md5"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
	"wordma/server/dto"
)

// User 用户表
type User struct {
	gorm.Model
	Name     string `gorm:"index;size:255"`
	Email    string `gorm:"index;size:255"`
	Url      string
	Password string
	LastIP   string
	LastUA   string
	Role     string
	Notice   bool `gorm:"default:false"`

	Comments []Comment `gorm:"foreignKey:UserID"` // 用户评论
}

// GetUserByID 通过用户id找到用户信息
func GetUserByID(id uint) (*User, error) {
	var user User
	err := DB.Limit(1).Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByName 通过用户Name找到用户信息
func GetUserByName(name string) (*User, error) {
	var user User
	err := DB.Limit(1).Where("name = ?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByEmail 通过用户email找到用户信息
func GetUserByEmail(email string) (*User, error) {
	var user User
	err := DB.Limit(1).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByNameOrEmail 通过name或email查找用户
func GetUserByNameOrEmail(nameOrEmail string) (*User, error) {
	var user User
	err := DB.Limit(1).Where("name = ? OR email = ?", nameOrEmail, nameOrEmail).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// SetPasswordEncrypt 设置密码加密
func (u *User) SetPasswordEncrypt(password string) (err error) {
	var encrypted []byte
	if encrypted, err = bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	); err != nil {
		return err
	}
	u.Password = "(bcrypt)" + string(encrypted)
	return nil
}

// CreateUser 新增用户
func CreateUser(data *User) error {
	err := DB.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

// FindOrCreateUser 查找或创建
func FindOrCreateUser(data dto.CommentDTO) (*User, error) {
	user, err := GetUserByEmail(data.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果没有这个用户，就创建
		user = &User{
			Name:  data.Name,
			Email: data.Email,
		}
		if err := CreateUser(user); err != nil {
			return nil, fmt.Errorf("创建用户失败: %w", err)
		}
		// 确保返回完整的用户信息
		user, err := GetUserByEmail(data.Email)
		if err != nil {
			return nil, fmt.Errorf("创建用户后查询失败: %w", err)
		}
		return user, nil
	}
	// 如果有这个用户，直接返回就好了
	return user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(id uint, data map[string]interface{}) error {
	var user User
	err := DB.Model(&user).Where("id = ? ", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// CreateAdministrator 创建管理员账户
func CreateAdministrator() {
	var err error
	admin := User{
		Name:  "admin",
		Email: "admin@example.com",
		Role:  "admin",
	}
	err = admin.SetPasswordEncrypt("123456")
	if err != nil {
		panic("创建管理员账户时出现错误！")
	}
	err = DB.Create(&admin).Error
	if err != nil {
		panic("创建管理员账户时，数据库出现错误！")
	}
}

func (u *User) CheckPassword(input string) bool {
	if u.ID == 0 {
		return false
	}
	password := strings.TrimSpace(u.Password)
	if password == "" {
		return false
	}

	const BcryptPrefix = "(bcrypt)"
	const MD5Prefix = "(md5)"

	switch {
	case strings.HasPrefix(password, BcryptPrefix):
		if err := bcrypt.CompareHashAndPassword([]byte(password[len(BcryptPrefix):]),
			[]byte(input)); err == nil {
			return true
		}
	case strings.HasPrefix(password, MD5Prefix):
		if strings.EqualFold(password[len(MD5Prefix):],
			fmt.Sprintf("%x", md5.Sum([]byte(input)))) {
			return true
		}
	default:
		if password == input {
			return true
		}
	}

	return false
}
