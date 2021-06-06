package utils

import (
	"github.com/go-chassis/openlog"
	"io/ioutil"
	"os"
	"strings"
)

func GetFileContent(path string) string {
	file, err := os.Open(path)
	if err != nil {
		openlog.Fatal("open file " + path + " failed. " + err.Error())
		return ""
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		openlog.Fatal("read file " + path + " failed. " + err.Error())
		return ""
	}
	return string(content)

}

func SaveFileContent(path string, content string) error {
	return ioutil.WriteFile(path, []byte(content), os.ModePerm)
}

// model

func GetUpperCamelStr(str string) string {
	strs := strings.Split(str, "_")
	var names []string
	for i := 1; i < len(strs); i++ {
		byteName := []byte(strs[i])
		byteName[0] -= 32
		names = append(names, string(byteName))
	}
	return strings.Join(names, "")
}
func GetLowerCamelStr(str string) string {
	strs := strings.Split(str, "_")
	var names []string
	names = append(names, strs[1])
	for i := 2; i < len(strs); i++ {
		byteName := []byte(strs[i])
		byteName[0] -= 32
		names = append(names, string(byteName))
	}
	return strings.Join(names, "")
}

func GetFieldModel(fieldName string) string {
	fieldNames := strings.Split(fieldName, "_")
	var names []string
	for i := 0; i < len(fieldNames); i++ {
		byteName := []byte(fieldNames[i])
		byteName[0] -= 32
		names = append(names, string(byteName))
	}
	return strings.Join(names, "")
}
func GetFieldType(fieldName string, fieldType string) string {
	if strings.HasSuffix(fieldName, "time") || strings.HasSuffix(fieldName, "times") || strings.HasSuffix(fieldName, "type") {
		return "uint32"
	}
	if strings.Index(fieldName, "id") != -1 {
		return "uint64"
	}
	if strings.Index(strings.ToLower(fieldType), "int") == -1 {
		return "string"
	}
	return "int32"
}
