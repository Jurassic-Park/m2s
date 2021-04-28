package core

import (
	"fmt"
	"github.com/Jurassic-Park/m2s/templates"
	"github.com/Jurassic-Park/m2s/util"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
)

var UCamelTableName string
var LCamelTableName string
var fieldSlic []SqlFieldDesc
var TableName string
var ServiceName string

// 开始生成服务文件
func Generator(connString string, tableName string, serviceName string) {
	// 大驼峰表名
	UCamelTableName = util.GeneratorCamelName(tableName, 1)
	LCamelTableName = util.GeneratorCamelName(tableName, 0)
	TableName = tableName
	ServiceName = serviceName

	mysql := Mysql{
		ConnString: connString,
		TableName:  tableName,
	}
	var err error
	fieldSlic, err = mysql.GetMysqlStruct()
	if err != nil {
		panic(err)
	}

	GeneratorModel()
	GeneratorService()
	GeneratorApi()
}

// 生成model
func GeneratorModel() {
	fmt.Println("------开始生成模型数据------")
	var fileString = templates.ModelTpl
	//整理参数
	format := map[string]string{
		"{{UCamelTableName}}": UCamelTableName,
		"{{LCamelTableName}}": LCamelTableName,
		"{{FieldStructData}}": GeneralModelStruct(fieldSlic),
		"{{tableName}}":       TableName,
		"{{ModelAddData}}":    GeneralModelAddData(fieldSlic),
	}
	//替换关键字
	for k, v := range format {
		fileString = strings.ReplaceAll(fileString, k, v)
	}
	// 当前有相同文件不更新
	fileDir := "./models/"
	fileName := TableName + ".go"
	filePath := fileDir + TableName + ".go"
	if ok, err := util.PathExists(filePath); err == nil && ok {
		fmt.Println("目录下存在相同文件:" + filePath)
		return
	}
	util.WriteFile(fileDir, fileName, fileString, 0755)
	fmt.Println("------生成模型成功-----")
}

// 生成service
func GeneratorService() {
	fmt.Println("------开始生成服务数据------")
	var fileString = templates.ServiceTpl
	//整理参数
	format := map[string]string{
		"{{UCamelTableName}}":    UCamelTableName,
		"{{LCamelTableName}}":    LCamelTableName,
		"{{tableName}}":          TableName,
		"{{ServiceStructData}}":  GeneralServiceStructData(fieldSlic),
		"{{ServiceSaveData}}":    GeneralServiceSaveData(fieldSlic),
		"{{ServiceSaveAddData}}": GeneralServiceSaveAddData(fieldSlic),
		"{{ServiceName}}":        ServiceName,
	}
	//替换关键字
	for k, v := range format {
		fileString = strings.ReplaceAll(fileString, k, v)
	}
	// 当前有相同文件不更新
	fileDir := "./service/" + TableName + "_service/"
	util.CreateDir(fileDir)

	fileName := TableName + ".go"
	filePath := fileDir + TableName + ".go"
	if ok, err := util.PathExists(filePath); err == nil && ok {
		fmt.Println("目录下存在相同文件:" + filePath)
		return
	}
	util.WriteFile(fileDir, fileName, fileString, 0755)
	fmt.Println("------生成模型成功-----")
}

// 生成api
func GeneratorApi() {
	fmt.Println("------开始生成API数据------")
	var apiString = templates.ApiTpl
	var partnerString = templates.PartnerTpl
	var adminString = templates.AdminTpl

	//整理参数
	format := map[string]string{
		"{{UCamelTableName}}":    UCamelTableName,
		"{{LCamelTableName}}":    LCamelTableName,
		"{{tableName}}":          TableName,
		"{{ApiValidData}}":       GeneralApiValidData(fieldSlic, false),
		"{{ApiSaveData}}":        GeneralApiSaveData(fieldSlic, true),
		"{{ApiUpdateValidData}}": GeneralApiValidData(fieldSlic, true),
		"{{ApiUpdateSaveData}}":  GeneralApiSaveData(fieldSlic, true),
		"{{ApiAllBackData}}":     GeneralAllBackData(fieldSlic),
		"{{ApiViewBackData}}":    GeneralViewBackData(fieldSlic),
		"{{ServiceName}}":        ServiceName,
	}
	//替换关键字
	for k, v := range format {
		apiString = strings.ReplaceAll(apiString, k, v)
		partnerString = strings.ReplaceAll(partnerString, k, v)
		adminString = strings.ReplaceAll(adminString, k, v)
	}

	// 检查routers文件夹
	routersPath := "./routers"
	if ok, err := util.PathExists(routersPath); err != nil || !ok {
		fmt.Println("dir not exist:" + routersPath)
		return
	}
	// 检查是否有 admin  partner api
	adminPath := routersPath + "/admin"
	partnerPath := routersPath + "/partner"
	apiPath := routersPath + "/api"
	if ok, err := util.PathExists(adminPath); err != nil && !ok {
		// 新建
		err := os.Mkdir(adminPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed (%s)![%v]\n", adminPath, err)
			return
		}
		fmt.Printf("mkdir " + adminPath + " success!\n")
	}
	if ok, err := util.PathExists(partnerPath); err != nil && !ok {
		// 新建
		err := os.Mkdir(partnerPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed (%s)![%v]\n", partnerPath, err)
			return
		}
		fmt.Printf("mkdir " + partnerPath + " success!\n")
	}
	if ok, err := util.PathExists(apiPath); err != nil && !ok {
		// 新建
		err := os.Mkdir(apiPath, os.ModePerm)
		if err != nil {
			fmt.Printf("mkdir failed (%s)![%v]\n", apiPath, err)
			return
		}
		fmt.Printf("mkdir " + apiPath + " success!\n")
	}

	// 生成三个相同的api接口

	// 当前有相同文件不更新
	fileName := TableName + ".go"

	apiFilePath := apiPath + "/" + fileName
	partnerFilePath := partnerPath + "/" + fileName
	adminFilePath := adminPath + "/" + fileName

	if ok, err := util.PathExists(apiFilePath); err == nil && ok {
		fmt.Println("目录下存在相同文件:" + apiFilePath)
	} else {
		util.WriteFile(apiPath, fileName, apiString, 0755)
	}
	if ok, err := util.PathExists(partnerFilePath); err == nil && ok {
		fmt.Println("目录下存在相同文件:" + partnerFilePath)
	} else {
		util.WriteFile(partnerPath, fileName, partnerString, 0755)
	}
	if ok, err := util.PathExists(adminFilePath); err == nil && ok {
		fmt.Println("目录下存在相同文件:" + adminFilePath)
	} else {
		util.WriteFile(adminPath, fileName, adminString, 0755)
	}
	fmt.Println("------生成API成功-----")
}
