package generate

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

// 根据modelInfo渲染model
func RenderModel(model *ModelInfo, tableName string) {
	file := createFile(tableName, "./model/")
	defer file.Close()
	render := getModelRender("./template/model.tpl")
	// 这里进行映射
	if err := render.Execute(file, model); err != nil {
		log.Fatal(err)
	}
}

// 定义model层我需要Reader的函数
func getModelRender(modelFile string) *template.Template {
	data, err := ioutil.ReadFile(modelFile)
	if nil != err {
		fmt.Println("read tplFile err:", err)
		return nil
	}
	render := template.Must(template.New("model").
		Funcs(template.FuncMap{
			"FirstCharUpper":       FirstCharUpper,
			"TypeConvert":          TypeConvert,
			"PkConvert":            PkConvert,
			"JsonAndFormTags":      JsonAndFormTags,
			"ExportColumn":         ExportColumn,
			"Join":                 Join,
			"MakeQuestionMarkList": MakeQuestionMarkList,
			"ColumnAndType":        ColumnAndType,
			"ColumnWithPostfix":    ColumnWithPostfix,
		}).
		Parse(string(data)))
	return render
}
