package templates

const ModelTpl = `package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"zhiyong/insure/framework/cache"
	log "zhiyong/insure/framework/glog"
	"zhiyong/insure/framework/models"
)

type {{UCamelTableName}} struct {
	models.Model
{{FieldStructData}}
}

// TableName sets the insert table name for this struct type
func (p *{{UCamelTableName}}) TableName() string {
	return "{{tableName}}"
}

const {{LCamelTableName}}CachePrefix = "{{LCamelTableName}}"

// Exist{{UCamelTableName}}ById checks if an {{LCamelTableName}} exists based on Id
func Exist{{UCamelTableName}}ById(id int, unscoped bool) (bool, error) {
	var {{LCamelTableName}} {{UCamelTableName}}

	cacheKey := {{LCamelTableName}}CachePrefix + fmt.Sprintf(":Exist{{UCamelTableName}}ById-%v", unscoped)

	var cc cache.Cache
	var err error
	if cc, err = cache.NewCache(cacheKey, id).GetRedisCache(&{{LCamelTableName}}); err == nil {
		if {{LCamelTableName}}.Id > 0 {
			log.Info("--命中cache--")
			return true, nil
		}
	}

	db := models.Db.Select("id")
	if !unscoped {
		db = db.Where("id = ? AND deleted_on = ? ", id, time.Time{})
	} else {
		db = db.Where("id = ? ", id)
	}
	err = db.First(&{{LCamelTableName}}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if {{LCamelTableName}}.Id > 0 {
		cc.SaveRedisCache({{LCamelTableName}})
		return true, nil
	}

	return false, nil
}

// Get{{UCamelTableName}}Total gets the total number of {{LCamelTableName}}s based on the constraints
func Get{{UCamelTableName}}Total(maps models.TableSearch) (int, error) {
	var count int
	cacheKey := {{LCamelTableName}}CachePrefix + ":Get{{UCamelTableName}}Total"
	var cc cache.Cache
	var err error
	if cc, err = cache.NewCache(cacheKey, maps).GetRedisCache(&count); err == nil {
		return count, nil
	}

	db := models.SearchConditionBuild(maps)
	if err := db.Model(&{{UCamelTableName}}{}).Count(&count).Error; err != nil {
		return 0, err
	}
	cc.SaveRedisCache(count)
	return count, nil
}

// Get{{UCamelTableName}}s gets a list of {{LCamelTableName}}s based on paging constraints
func Get{{UCamelTableName}}s(search models.TableSearch) ([]*{{UCamelTableName}}, error) {
	var {{LCamelTableName}}s []*{{UCamelTableName}}
	cacheKey := {{LCamelTableName}}CachePrefix + ":Get{{UCamelTableName}}s"

	var cc cache.Cache
	var err error
	if cc, err = cache.NewCache(cacheKey, search).GetRedisCache(&{{LCamelTableName}}s); err == nil {
		log.Info("--命中cache--")
		return {{LCamelTableName}}s, err
	}

	// get query
	db, err := models.SearchBuild(search)
	if err != nil {
		return nil, err
	}

	err = db.Find(&{{LCamelTableName}}s).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	cc.SaveRedisCache({{LCamelTableName}}s)
	return {{LCamelTableName}}s, nil
}

// Get{{UCamelTableName}} Get a single {{LCamelTableName}} based on Id
func Get{{UCamelTableName}}(search models.TableSearch) (*{{UCamelTableName}}, error) {
	var {{LCamelTableName}} {{UCamelTableName}}
	cacheKey := {{LCamelTableName}}CachePrefix + ":Get{{UCamelTableName}}"

	var cc cache.Cache
	var err error
	if cc, err = cache.NewCache(cacheKey, search).GetRedisCache(&{{LCamelTableName}}); err == nil {
		log.Info("--命中cache--")
		return &{{LCamelTableName}}, err
	}
	db, err := models.SearchBuild(search)
	if err != nil {
		return nil, err
	}

	err = db.First(&{{LCamelTableName}}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	cc.SaveRedisCache({{LCamelTableName}})
	return &{{LCamelTableName}}, nil
}

// Edit{{UCamelTableName}} modify a single {{LCamelTableName}}
func Edit{{UCamelTableName}}(id int, data interface{}) error {
	if err := models.Db.Model(&{{UCamelTableName}}{}).Where("id = ? AND deleted_on = ? ", id, time.Time{}).Updates(data).Error; err != nil {
		return err
	}

	cache.DeleteRedisCache({{LCamelTableName}}CachePrefix)
	return nil
}

// Add{{UCamelTableName}} add a single {{LCamelTableName}}
func Add{{UCamelTableName}}({{LCamelTableName}} {{UCamelTableName}}) (int, error) {
	if err := models.Db.Create(&{{LCamelTableName}}).Error; err != nil {
		return 0, err
	}

	cache.DeleteRedisCache({{LCamelTableName}}CachePrefix)
	return {{LCamelTableName}}.Id, nil
}

// Delete{{UCamelTableName}} delete a single {{LCamelTableName}}
func Delete{{UCamelTableName}}(search models.TableSearch) error {
	if len(search.WhereMaps) == 0 {
		return errors.New("删除异常[999]")
	}
	db := models.SearchConditionBuild(search)
	if err := db.Delete({{UCamelTableName}}{}).Error; err != nil {
		return err
	}

	cache.DeleteRedisCache({{LCamelTableName}}CachePrefix)
	return nil
}

// CleanAll{{UCamelTableName}} clear all {{LCamelTableName}}
//func CleanAll{{UCamelTableName}}() error {
//	if err := models.Db.Unscoped().Where("deleted_on != ? ", 0).Delete(&{{UCamelTableName}}{}).Error; err != nil {
//		return err
//	}
//
//	return nil
//}

`
