package templates

var ServiceTpl = `package {{tableName}}_service

import (
	"zhiyong/insure/{{ServiceName}}/models"
	"time"
)

type {{UCamelTableName}} struct {
	Id       int

{{ServiceStructData}}
	PageNum  int
	PageSize int
	Query    map[string]string
}

// 保存
func (c *{{UCamelTableName}}) Save() (int, error) {
	data := map[string]interface{}{
{{ServiceSaveData}}
	}

	ok, err := c.ExistById()
	if err != nil {
		return 0, err
	}

	if ok {
		// 更新
		err := models.Edit{{UCamelTableName}}(c.Id, data)
		return c.Id, err
	}

	return models.Add{{UCamelTableName}}(data)
}

// Get ...
func (c *{{UCamelTableName}}) Get() (*models.{{UCamelTableName}}, error) {
	return models.Get{{UCamelTableName}}(c.Id)
}

// GetInfo
func (c *{{UCamelTableName}}) GetInfo() (*models.{{UCamelTableName}}, error) {
	return models.Get{{UCamelTableName}}Info(c.getMaps())
}

// GetAll ...
func (c *{{UCamelTableName}}) GetAll() ([]*models.{{UCamelTableName}}, error) {
	var (
		list []*models.{{UCamelTableName}}
	)

	// 缓存

	list, err := models.Get{{UCamelTableName}}s(c.PageNum, c.PageSize, c.getMaps())
	if err != nil {
		return nil, err
	}

	return list, nil
}

// 删除
func (c *{{UCamelTableName}}) Delete() error {
	return models.Delete{{UCamelTableName}}(c.Id)
}

// ExistById ...
func (c *{{UCamelTableName}}) ExistById() (bool, error) {
	return models.Exist{{UCamelTableName}}ById(c.Id)
}

// Count ...
func (c *{{UCamelTableName}}) Count() (int, error) {
	return models.Get{{UCamelTableName}}Total(c.getMaps())
}

// getMaps ....
func (c *{{UCamelTableName}}) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	// 传递到下一层
	for k, v := range c.Query{
		if v != "" {
			maps[k] = v
		}
	}

	maps["deleted_on"] = time.Time{}
	return maps
}
`
