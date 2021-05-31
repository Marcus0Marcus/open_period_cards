package main

import (
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/go-chassis/openlog"
	"open_period_cards/db_tools/global_conf"
	"open_period_cards/db_tools/model_generator"
	"open_period_cards/db_tools/service_generator"
	"open_period_cards/db_tools/utils"
	"strings"
)

var (
	defaultSQLPath = "./design/db.sql"
)

func loadSQLContent() []string {
	var ddls []string
	sqlContent := utils.GetFileContent(defaultSQLPath)
	ddls = strings.Split(string(sqlContent), global_conf.DefaultTplSplit)
	return ddls
}
func parseDDL(ddl *sqlparser.CreateTable) {
	fmt.Println("parsing ddl table_name:", ddl.NewName.Name)
	model_generator.GenerateModel(ddl)
	service_generator.GenerateService(ddl)

}
func parseDDLs(strSQL string) {
	stmt, err := sqlparser.Parse(strSQL)
	if err != nil {
		openlog.Fatal("parse SQL statement fail." + err.Error())
		return
	}
	switch stmt.(type) {
	case *sqlparser.CreateTable:
		ddl := stmt.(*sqlparser.CreateTable)
		if ddl.DDL.Action == "create" {
			parseDDL(ddl)
		} else {
			openlog.Fatal("SQL statement not complete DDL.")
		}

	default:
		openlog.Fatal("SQL statement not complete DDL.")
	}
}

func main() {
	ddls := loadSQLContent()
	for i := 0; i < len(ddls); i++ {
		parseDDLs(ddls[i])
	}

}
