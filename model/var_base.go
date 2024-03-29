package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	// Middleware 查询中间操作
	Middleware func(db *gorm.DB) *gorm.DB

	// Objects List
	Objects struct {
		Obj        interface{}
		Model      *gorm.DB
		Pagination *Pagination
		Result     *ObjResult
	}
	// ObjResult Result
	ObjResult struct {
		*Pagination
		List interface{} `json:"list"`
	}
	// Pagination 分页数据
	Pagination struct {
		Page      int    `json:"page"`
		Size      int    `json:"size"`
		Total     int    `json:"total"`
		URLFormat string `json:"-"`
	}
	// Page 分页
	Page struct {
		Page int
		URL  string
	}

	// GormDB 自定义方法
	GormDB struct {
		*gorm.DB
	}
)

// ConnectMysql 链接mysql
func (g *GormDB) ConnectMysql(dsn string) (err error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Info),
		SkipDefaultTransaction: true,
	})
	g.DB = db
	return
}

// SetURLFormat 设置url
func (p *Pagination) SetURLFormat(url string) string {
	p.URLFormat = url
	return ""
}

// URL 获取url
func (p *Pagination) URL(page int, size int) string {
	if p.URLFormat == "" {
		p.URLFormat = "?page=%d&size=%d"
	}
	return fmt.Sprintf(p.URLFormat, page, size)
}

// MarshalJSON MarshalJSON
func (o *ObjResult) MarshalJSON() ([]byte, error) {
	if _json, err := json.Marshal(map[string]interface{}{
		"page":       o.Page,
		"size":       o.Size,
		"page_count": o.Pages(),
		"total":      o.Total,
		"list":       o.List,
	}); err == nil {
		return _json, nil
	}
	return []byte("{}"), nil
}

// Pages 总页数
func (p *Pagination) Pages() int {
	f := float64(p.Total) / float64(p.Size)
	return int(math.Ceil(f))
}

// Range 生成数组
func (p *Pagination) Range() []Page {
	_page := make([]Page, p.Pages())
	for i := range _page {
		_page[i].Page = i + 1
		_page[i].URL = p.URL(i+1, p.Size)
	}
	return _page
}

// All 全部数据
func (o *Objects) All() (err error) {
	err = o.Model.Find(o.Obj).Error
	return
}

// Paging 分页数据
func (o *Objects) Paging(page int, size int, args ...Middleware) (err error) {
	offset := size * (page - 1)
	var count int64
	row := o.Model.Count(&count)

	//中间操作，count之后会丢失select，暂不清楚是否有其他异常
	for _, midd := range args {
		if midd != nil {
			row = midd(row)
		}
	}

	row.Offset(offset).Limit(size).Find(o.Obj)
	if row.Error != nil {
		err = row.Error
	}
	o.Pagination = &Pagination{
		Size:  size,
		Page:  page,
		Total: int(count),
	}
	o.Result = &ObjResult{
		List:       o.Obj,
		Pagination: o.Pagination,
	}
	return
}

//GetObjectOrNotFound 获取某一条数据
//gorm  查询接收空条件，在某些情况下会操作到错误到数据
func (g *GormDB) GetObjectOrNotFound(obj interface{}, query map[string]interface{}, midd ...Middleware) (err error) {
	row := g.Model(obj)
	if query != nil {
		row = row.Where(query)
	} else {
		return errors.New("query 不可为nil")
	}
	for _, mid := range midd {
		row = mid(row)
	}
	err = row.First(obj).Error
	return
}

// GetObjectsOrEmpty 获取列表 \n
// 可选参数 middleware models.middleware 接收一个 *gorm.DB 返回 *gorm.DB
func (g *GormDB) GetObjectsOrEmpty(obj interface{}, query interface{}, args ...Middleware) *Objects {
	row := g.Model(obj)
	if query != nil {
		row = row.Where(query)
	}
	// 可选参数
	for _, midd := range args {
		if midd != nil {
			row = midd(row)
		}
	}
	return &Objects{
		Model: row,
		Obj:   obj,
	}
}

// GetParamsTryInt 字符串转数字
func GetParamsTryInt(val string, defaults int) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		num = defaults
	}
	return num
}

// GetPagingParams 获取分页参数
func GetPagingParams(c *gin.Context) (page int, size int) {
	page = GetParamsTryInt(c.Query("page"), 1)
	size = GetParamsTryInt(c.Query("size"), 10)
	return
}
