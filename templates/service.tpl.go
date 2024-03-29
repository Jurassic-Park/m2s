package templates

var ServiceTpl = `package {{tableName}}_service

import (
	"context"
	models2 "zhiyong/insure/framework/models"
	"zhiyong/insure/{{ServiceName}}/models"
)

type {{UCamelTableName}} struct {
	Ctx context.Context

	Id       int
{{ServiceStructData}}

	models2.TableSearch
}

// 保存
func (c *{{UCamelTableName}}) Save() (int, error) {

	ok, err := c.ExistById()
	if err != nil {
		return 0, err
	}

	if ok {
		// 更新
		data := map[string]interface{}{
				{{ServiceSaveData}}
		}
		err := models.Edit{{UCamelTableName}}(c.Ctx, c.Id, data)
		return c.Id, err
	}

	data := models.{{UCamelTableName}}{
    		{{ServiceSaveAddData}}
	}
	return models.Add{{UCamelTableName}}(c.Ctx, data)
}

// Get ...
func (c *{{UCamelTableName}}) Get() (*models.{{UCamelTableName}}, error) {
	return models.Get{{UCamelTableName}}(c.Ctx, c.TableSearch)
}


// GetAll ...
func (c *{{UCamelTableName}}) GetAll() ([]*models.{{UCamelTableName}}, error) {
	var (
		list []*models.{{UCamelTableName}}
	)

	// 缓存

	list, err := models.Get{{UCamelTableName}}s(c.Ctx, c.TableSearch)
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 删除
func (c *{{UCamelTableName}}) Delete() error {
	return models.Delete{{UCamelTableName}}(c.Ctx, c.TableSearch)
}

// ExistById ...
func (c *{{UCamelTableName}}) ExistById() (bool, error) {	
	if c.Id == 0 {
		return false, nil    
	}
	return models.Exist{{UCamelTableName}}ById(c.Ctx, c.Id, false)
}

// Count ...
func (c *{{UCamelTableName}}) Count() (int, error) {
	return models.Get{{UCamelTableName}}Total(c.Ctx, c.TableSearch)
}
`
