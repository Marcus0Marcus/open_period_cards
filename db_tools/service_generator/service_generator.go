package service_generator

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/go-chassis/openlog"
	"open_period_cards/db_tools/global_conf"
	"open_period_cards/db_tools/utils"
	"strings"
)

var (
	defaultTplPath               = "./design/service.tpl"
	defaultPackageName           = "{package_name}"
	defaultServiceModelUpperCase = "{ServiceModelUpper}"
	defaultServiceModelLowerCase = "{ServiceModelLower}"
	defaultModelSuffix           = "{ServiceModelSuffix}"
	defaultCacheKeyPrefix        = "{CacheKeyPrefix}"
)

func GenerateService(ddl *sqlparser.CreateTable) {
	tplContent := utils.GetFileContent(defaultTplPath)

	serviceModelUpperCase := utils.GetUpperCamelStr(ddl.NewName.Name.String())
	serviceModelLowerCase := utils.GetLowerCamelStr(ddl.NewName.Name.String())
	// package name
	content := strings.ReplaceAll(tplContent, defaultPackageName, global_conf.DefaultPackageName)
	// serviceModelUpCase
	content = strings.ReplaceAll(content, defaultServiceModelUpperCase, serviceModelUpperCase)
	// serviceModelLowerCase
	content = strings.ReplaceAll(content, defaultServiceModelLowerCase, serviceModelLowerCase)
	// serviceModelSuffix
	content = strings.ReplaceAll(content, defaultModelSuffix, global_conf.DefaultModelSuffix)
	// CacheKeyPrefix
	content = strings.ReplaceAll(content, defaultCacheKeyPrefix, strings.ReplaceAll(ddl.NewName.Name.String(), "tb_", "")+"_")
	//for i := 0; i < len(ddl.Options); i++ {
	//	if ddl.Options[i].Type == sqlparser.TableOptionComment {
	//		commentArr := strings.Split(strings.ReplaceAll(ddl.Options[i].String(), "COMMENT=", ""), ",")
	//		for j := 0; j < len(commentArr); j++ {
	//			if strings.Index(commentArr[j], "cacheKey") != -1 {
	//				cacheKey := strings.Split(commentArr[j], "|")[1]
	//				fieldName := utils.GetUpperCamelStr(cacheKey)
	//				fmt.Println(fieldName)
	//			}
	//		}
	//	}
	//}
	path := global_conf.GeneratePath + strings.ReplaceAll(ddl.NewName.Name.String(), "tb_", "") + ".go"
	err := utils.SaveFileContent(path, content)
	if err != nil {
		openlog.Fatal("save service " + path + " failed.")
	}
}
