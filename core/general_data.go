package core

import "github.com/Jurassic-Park/m2s/util"

func GeneralModelStruct(fields []SqlFieldDesc) string {
	schema := ""
	for _, v := range fields {
		goType := util.DbTypeToGoType(v.COLUMN_TYPE)

		if v.COLUMN_NAME == "id" || v.COLUMN_NAME == "created_on" || v.COLUMN_NAME == "modified_on" || v.COLUMN_NAME == "deleted_on" {
			continue
		}
		schema += "    " + util.GeneratorCamelName(v.COLUMN_NAME, 1) + " " + goType + " `gorm:\"column:" + v.COLUMN_NAME + "\" json:\"" + v.COLUMN_NAME + "\" description:\"" + v.COLUMN_COMMENT + "\"`"
		schema += "\n"
	}

	return schema
}

// Name: data["name"].(string),
func GeneralModelAddData(fields []SqlFieldDesc) string {
	schema := ""
	for _, v := range fields {
		goType := util.DbTypeToGoType(v.COLUMN_TYPE)

		if v.COLUMN_NAME == "id" || v.COLUMN_NAME == "created_on" || v.COLUMN_NAME == "modified_on" || v.COLUMN_NAME == "deleted_on" {
			continue
		}
		schema += "        " + util.GeneratorCamelName(v.COLUMN_NAME, 1) + ": data[\"" + v.COLUMN_NAME + "\"].(" + goType + "),"
		schema += "\n"
	}

	return schema
}

// valid.Required(r.Name, "name")   //不能为空
func GeneralApiValidData(fields []SqlFieldDesc) string {
	schema := ""
	for _, v := range fields {
		// goType := util.DbTypeToGoType(v.COLUMN_TYPE)

		if v.COLUMN_NAME == "id" || v.COLUMN_NAME == "created_on" || v.COLUMN_NAME == "modified_on" || v.COLUMN_NAME == "deleted_on" {
			continue
		}
		schema += "    valid.Required(r." + util.GeneratorCamelName(v.COLUMN_NAME, 0) + ", \"" + v.COLUMN_NAME + "\")"
		schema += "\n"
	}

	return schema
}

// GeneralApiSaveData 组装服务时数据 		Id:       int(r.Id),
func GeneralApiSaveData(fields []SqlFieldDesc) string {
	schema := ""
	for _, v := range fields {
		goType := util.DbTypeToGoType(v.COLUMN_TYPE)

		if v.COLUMN_NAME == "id" || v.COLUMN_NAME == "created_on" || v.COLUMN_NAME == "modified_on" || v.COLUMN_NAME == "deleted_on" {
			continue
		}
		temp := "r." + util.GeneratorCamelName(v.COLUMN_NAME, 1)
		if goType == "int" {
			temp = "int(" + temp + ")"
		}

		schema += "        " + util.GeneratorCamelName(v.COLUMN_NAME, 1) + ": " + temp + ","
		schema += "\n"
	}

	return schema
}

// GeneralServiceStructData 服务里struct 数据		Name string
func GeneralServiceStructData(fields []SqlFieldDesc) string {
	schema := ""
	for _, v := range fields {
		goType := util.DbTypeToGoType(v.COLUMN_TYPE)

		if v.COLUMN_NAME == "id" || v.COLUMN_NAME == "created_on" || v.COLUMN_NAME == "modified_on" || v.COLUMN_NAME == "deleted_on" {
			continue
		}

		schema += "    " + util.GeneratorCamelName(v.COLUMN_NAME, 1) + " " + goType
		schema += "\n"
	}

	return schema
}

// GeneralServiceSaveData 服务里save 数据				"name":      c.Name,
func GeneralServiceSaveData(fields []SqlFieldDesc) string {
	schema := ""
	for _, v := range fields {
		// goType := util.DbTypeToGoType(v.COLUMN_TYPE)

		if v.COLUMN_NAME == "id" || v.COLUMN_NAME == "created_on" || v.COLUMN_NAME == "modified_on" || v.COLUMN_NAME == "deleted_on" {
			continue
		}

		schema += "        \"" + v.COLUMN_NAME + "\": c." + util.GeneratorCamelName(v.COLUMN_NAME, 1) + ","
		schema += "\n"
	}

	return schema
}

// GeneralAllBackData api返回时的数据 				Id:       int32(v.ID),	Name:     v.Name,
func GeneralAllBackData(fields []SqlFieldDesc) string {
	schema := ""
	for _, v := range fields {
		goType := util.DbTypeToGoType(v.COLUMN_TYPE)

		if v.COLUMN_NAME == "id" || v.COLUMN_NAME == "created_on" || v.COLUMN_NAME == "modified_on" || v.COLUMN_NAME == "deleted_on" {
			continue
		}
		temp := "v." + util.GeneratorCamelName(v.COLUMN_NAME, 1)
		if goType == "int" {
			temp = "int32(" + temp + ")"
		}
		schema += "        " + util.GeneratorCamelName(v.COLUMN_NAME, 1) + ": " + temp + ","
		schema += "\n"
	}

	return schema
}

// GeneralViewBackData api返回时的数据 				Name:     {{LCamelTableName}}.Name,
func GeneralViewBackData(fields []SqlFieldDesc) string {
	schema := ""
	for _, v := range fields {
		goType := util.DbTypeToGoType(v.COLUMN_TYPE)

		if v.COLUMN_NAME == "id" || v.COLUMN_NAME == "created_on" || v.COLUMN_NAME == "modified_on" || v.COLUMN_NAME == "deleted_on" {
			continue
		}
		temp := LCamelTableName + "." + util.GeneratorCamelName(v.COLUMN_NAME, 1)
		if goType == "int" {
			temp = "int32(" + temp + ")"
		}

		schema += "        " + util.GeneratorCamelName(v.COLUMN_NAME, 1) + ": " + temp + ","
		schema += "\n"
	}

	return schema
}
