package templates

const ModelTpl = `package models

import (
	"github.com/jinzhu/gorm"
	"gitlab.ronshubao.com/grpc-insure/framework/models"
	"time"
)

type {{UCamelTableName}} struct {
	models.Model
{{FieldStructData}}
}

// TableName sets the insert table name for this struct type
func (p *{{UCamelTableName}}) TableName() string {
	return "{{tableName}}"
}


// Exist{{UCamelTableName}}ById checks if an {{LCamelTableName}} exists based on Id
func Exist{{UCamelTableName}}ById(id int) (bool, error) {
	var {{LCamelTableName}} {{UCamelTableName}}
	err := models.Db.Select("id").Where("id = ? AND deleted_on = ? ", id, time.Time{}).First(&{{LCamelTableName}}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if {{LCamelTableName}}.Id > 0 {
		return true, nil
	}

	return false, nil
}

// Get{{UCamelTableName}}Total gets the total number of {{LCamelTableName}}s based on the constraints
func Get{{UCamelTableName}}Total(maps interface{}) (int, error) {
	// get query
	whereSql, args, err := models.AutoBuildWhere(maps)
	if err != nil {
		return 0, err
	}

	var count int
	if err := models.Db.Model(&{{UCamelTableName}}{}).Where(whereSql, args...).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// Get{{UCamelTableName}}s gets a list of {{LCamelTableName}}s based on paging constraints
func Get{{UCamelTableName}}s(pageNum int, pageSize int, maps interface{}) ([]*{{UCamelTableName}}, error) {
	// get query
	whereSql, args, err := models.AutoBuildWhere(maps)
	if err != nil {
		return nil, err
	}

	var {{LCamelTableName}}s []*{{UCamelTableName}}
	err = models.Db.Where(whereSql, args...).Offset(pageNum).Limit(pageSize).Find(&{{LCamelTableName}}s).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return {{LCamelTableName}}s, nil
}

// Get{{UCamelTableName}} Get a single {{LCamelTableName}} based on Id
func Get{{UCamelTableName}}(id int) (*{{UCamelTableName}}, error) {
	var {{LCamelTableName}} {{UCamelTableName}}
	err := models.Db.Where("id = ? AND deleted_on = ? ", id, time.Time{}).First(&{{LCamelTableName}}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &{{LCamelTableName}}, nil
}

// Get{{UCamelTableName}} Get a single {{LCamelTableName}} based on maps
func Get{{UCamelTableName}}Info(maps interface{}) (*{{UCamelTableName}}, error) {
	whereSql, args, err := models.AutoBuildWhere(maps)
	if err != nil {
		return nil, err
	}

	var {{LCamelTableName}} {{UCamelTableName}}
	err = models.Db.Where(whereSql, args...).First(&{{LCamelTableName}}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &{{LCamelTableName}}, nil
}

// Edit{{UCamelTableName}} modify a single {{LCamelTableName}}
func Edit{{UCamelTableName}}(id int, data interface{}) error {
	if err := models.Db.Model(&{{UCamelTableName}}{}).Where("id = ? AND deleted_on = ? ", id, time.Time{}).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// Add{{UCamelTableName}} add a single {{LCamelTableName}}
func Add{{UCamelTableName}}(data map[string]interface{}) (int, error) {
	{{LCamelTableName}} := {{UCamelTableName}}{
{{ModelAddData}}
	}
	if err := models.Db.Create(&{{LCamelTableName}}).Error; err != nil {
		return 0, err
	}

	return {{LCamelTableName}}.Id, nil
}

// Delete{{UCamelTableName}} delete a single {{LCamelTableName}}
func Delete{{UCamelTableName}}(id int) error {
	if err := models.Db.Where("id = ?", id).Delete({{UCamelTableName}}{}).Error; err != nil {
		return err
	}

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
