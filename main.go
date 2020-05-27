package main

import (
	"fmt"
	"github.com/Jurassic-Park/m2s/core"
	"github.com/droundy/goopt"
)

var (
	sqlTable    = goopt.String([]string{"-t", "--table"}, "", "Table to build struct from")
	connString  = goopt.String([]string{"-m", "--mysql"}, "", "mysql config")
	serviceName = goopt.String([]string{"-pn", "--serviceName"}, "", "serviceName config")
)

func init() {
	goopt.Description = func() string {
		return "m2s is tool that can automaticlly generate proto file."
	}
	goopt.Version = "0.1"
	goopt.Summary = `m2s --mysql user:password@tcp\(host:port\)/database\?charset=utf8 --table tableName --protoName`
	goopt.Parse(nil)
}

func main() {
	if *connString == "" {
		fmt.Println("mysql connect can not is empty")
		return
	}
	if *sqlTable == "" {
		fmt.Println("table can not is empty")
		return
	}
	if *serviceName == "" {
		fmt.Println("serviceName can not is empty")
		return
	}
	core.Generator(*connString, *sqlTable, *serviceName)
}
