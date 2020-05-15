package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func GetBetweenStr(str, start, end string) string {
	n := strings.Index(str, start)
	n += len(start)
	if n == -1 {
		n = 0
	}
	str = string([]byte(str)[n:])
	m := strings.Index(str, end)
	if m == -1 {
		m = len(str)
	}
	str = string([]byte(str)[:m])
	return str
}

//首字母大写其他小写
func FirstToUpper(s string) string {
	return strings.ToUpper(s[0:1]) + strings.ToLower(s[1:])
}

// 生成驼峰
// tag 大驼峰 1， 小驼峰 0
func GeneratorCamelName(str string, tag int) (name string) {
	parts := strings.Split(str, "_")
	if tag == 1 {
		for _, v := range parts {
			name += FirstToUpper(v)
		}
	} else {
		for k, v := range parts {
			if k != 0 {
				name += FirstToUpper(v)
			} else {
				name += v
			}
		}
	}
	return
}

// 创建目录
func CreateDir(dir string) {
	if ok, err := PathExists(dir); err == nil && ok == false {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
	}
}

//写入文件
func WriteFile(fileDir string, fileName string, file string, mode os.FileMode) error {
	_, err := os.Stat(fileDir)
	if err != nil {
		err = os.MkdirAll(fileDir, mode)
		if err != nil {
			log.Fatalln(err.Error() + ": " + fileDir)
		}
	}
	fn := filepath.Join(fileDir, fileName)
	err = ioutil.WriteFile(fn, []byte(file), mode)
	if err != nil {
		log.Fatalln(err.Error() + ": " + fn)
	}
	fmt.Println("success create :" + fn)
	return err
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// DbTypeToGoType mysql 类型转换为go类型
func DbTypeToGoType(columnType string) string {
	var goType string
	if strings.Index(columnType, "bigint") > -1 {
		goType = "int64"
	} else if strings.Index(columnType, "int") > -1 {
		goType = "int"
	} else if strings.Index(columnType, "text") > -1 {
		goType = "string"
	} else if strings.Index(columnType, "char") > -1 {
		goType = "string"
	} else if strings.Index(columnType, "enum") > -1 {
		goType = "string"
	} else if strings.Index(columnType, "blob") > -1 {
		goType = "string"
	} else if strings.Index(columnType, "float") > -1 {
		goType = "float32"
	} else if strings.Index(columnType, "double") > -1 {
		goType = "float64"
	} else if strings.Index(columnType, "date") > -1 {
		goType = "time.Time"
	} else if strings.Index(columnType, "time") > -1 {
		goType = "time.Time"
	} else {
		goType = "string"
	}

	return goType
}
