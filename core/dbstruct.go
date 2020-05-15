package core

import (
	"database/sql"
	"fmt"
	"github.com/Jurassic-Park/m2s/util"
)

// 数据库字段
type SqlFieldDesc struct {
	COLUMN_NAME    string
	COLUMN_COMMENT string
	COLUMN_TYPE    string
}

type Mysql struct {
	ConnString string
	TableName  string
	db         *sql.DB
}

func (m *Mysql) getDb() {
	var err error
	db, err := sql.Open("mysql", m.ConnString)
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(50)
	if err != nil {
		fmt.Printf("connect mysql fail ! [%s]", err)
		panic("connect mysql fail")
	}

	fmt.Println("connect to mysql success")

	m.db = db
}

// GetMysqlStruct 获取mysql结构的slic
func (m *Mysql) GetMysqlStruct() ([]SqlFieldDesc, error) {
	var slic = make([]SqlFieldDesc, 0)
	var sqlFieldDesc = new(SqlFieldDesc)
	m.getDb()
	defer m.db.Close()

	tableSchema := util.GetBetweenStr(m.ConnString, ")/", "?")
	fmt.Println("table schema is " + tableSchema)

	querySql := "select COLUMN_NAME, COLUMN_COMMENT, COLUMN_TYPE from information_schema.columns where table_schema ='" + tableSchema + "' and table_name = '" + TableName + "' ;"
	rows, err := m.db.Query(querySql)
	defer rows.Close()
	if err != nil {
		return slic, err
	}
	if err != nil {
		return slic, err
	}
	for rows.Next() {
		//定义变量接收查询数据
		err := rows.Scan(&sqlFieldDesc.COLUMN_NAME, &sqlFieldDesc.COLUMN_COMMENT, &sqlFieldDesc.COLUMN_TYPE)
		if err != nil {
			return slic, err
		}
		slic = append(slic, *sqlFieldDesc)
	}
	return slic, err
}
