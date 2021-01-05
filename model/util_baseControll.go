package model

import (
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Controller ModelInterface
type Controller interface {
	// add update 验证
	Validator() error
	//添加数据生成uuid
	SetCode() error
	//添加数据绑定用户id
	SetUser(c *gin.Context, data interface{}) error

	// 表名
	TableName() string
	// 自身实例 用于found one
	Object() interface{}
	// 自身实例 用于found list
	Objects() interface{}

	// 以下三个方法用于辅助默认方法实现curd，过于复杂的直接override
	// Where 搜索条件
	// Search(db *gorm.DB, key string) *gorm.DB
	// 查询的补充条件
	Joins(db *gorm.DB) *gorm.DB
	// 处理列表返回结果
	Result(data interface{}) interface{}
}

// Model 空方法用户数据模型继承方法
type Model struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	Code      string         `json:"code" gorm:"primary_key;index;unique;not null"` //
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}
