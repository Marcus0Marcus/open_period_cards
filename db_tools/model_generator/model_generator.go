package model_generator

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/go-chassis/openlog"
	"open_period_cards/db_tools/global_conf"
	"open_period_cards/db_tools/utils"
	"strings"
)

var (
	defaultTplPath         = "./design/model.tpl"
	defaultTableModel      = "TBName"
	defaultTableName       = "tb_name"
	defaultFieldModel      = "FieldName"
	defaultFieldName       = "field_name"
	defaultFieldNameSplit  = "|"
	defaultModelReplaceNum = 3
	defaultMapFields       = map[string]string{
		"id":      "\tId           uint64  `json:\"id\" gorm:\"column:id\"`\n",
		"mtime":   "\tMtime   uint32  `json:\"mtime\" gorm:\"autoUpdateTime <-:update\"`\n",
		"ctime":   "\tCtime   uint32  `json:\"ctime\" gorm:\"autoCreateTime <-:create\"`\n",
		"deleted": "\tDeleted int32  `json:\"deleted\" gorm:\"column:deleted\"`\n"}
)

func generateFields(tpl string, columns []*sqlparser.ColumnDef) string {
	var fields []string
	for i := 0; i < len(columns); i++ {
		field, ok := defaultMapFields[columns[i].Name]
		if ok {
			fields = append(fields, field)
		} else {
			fieldType := utils.GetFieldType(columns[i].Name, columns[i].Type)
			tpls := strings.Split(tpl, defaultFieldNameSplit)
			tpls[0] = "\t" + strings.ReplaceAll(tpls[0], defaultFieldModel, utils.GetFieldModel(columns[i].Name))
			tpls[1] = fieldType
			tpls[2] = strings.ReplaceAll(tpls[2], defaultFieldName, columns[i].Name)
			fields = append(fields, strings.Join(tpls, "\t"))
		}
	}
	return strings.Join(fields, "")
}
func GenerateModel(ddl *sqlparser.CreateTable) {
	tplContent := utils.GetFileContent(defaultTplPath)
	tplContents := strings.Split(tplContent, global_conf.DefaultTplSplit)
	if len(tplContents) != defaultModelReplaceNum {
		openlog.Fatal("model Tpl replace number wrong.")
	}
	tableModel := utils.GetUpperCamelStr(ddl.NewName.Name.String()) + global_conf.DefaultModelSuffix

	tplContents[0] = strings.ReplaceAll(tplContents[0], defaultTableModel, tableModel)
	// generate fields

	tplContents[1] = generateFields(tplContents[1], ddl.Columns)
	tplContents[2] = strings.ReplaceAll(tplContents[2], defaultTableModel, tableModel)
	tplContents[2] = strings.ReplaceAll(tplContents[2], defaultTableName, ddl.NewName.Name.String())
	modelContent := strings.Join(tplContents, "\n")
	path := global_conf.GeneratePath + strings.ReplaceAll(ddl.NewName.Name.String()+"_m.go", "tb_", "")
	err := utils.SaveFileContent(path, modelContent)
	if err != nil {
		openlog.Fatal("save model " + path + " failed.")
	}
}
