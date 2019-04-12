package generate

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

// 根据modelInfo渲染model
func RenderService(model *ModelInfo, tableName string) {
	file := createFile(tableName, "./service/")
	defer file.Close()
	render := getServiceRender("./template/service.tpl")
	// 这里进行映射
	if err := render.Execute(file, model); err != nil {
		log.Fatal(err)
	}
}

func getServiceRender(modelFile string) *template.Template {
	data, err := ioutil.ReadFile(modelFile)
	if nil != err {
		fmt.Println("read tplFile err:", err)
		return nil
	}
	render := template.Must(template.New("service").
		Funcs(template.FuncMap{
			"FirstCharUpper":       FirstCharUpper,
			"TypeConvert":          TypeConvert,
			"PkConvert":            PkConvert,
			"ExportColumn":         ExportColumn,
			"Join":                 Join,
			"MakeQuestionMarkList": MakeQuestionMarkList,
			"ColumnAndType":        ColumnAndType,
			"ColumnWithPostfix":    ColumnWithPostfix,
		}).
		Parse(string(data)))
	return render
}
