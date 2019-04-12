package generate

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
)

//根据modelInfo渲染route_api
func RenderRoute(model *ModelInfo, tableName string) {
	file := createFile(tableName, "./routes/api/")
	defer file.Close()
	render := getRouteRender("./template/routes.tpl")
	// 这里进行映射
	if err := render.Execute(file, model); err != nil {
		log.Fatal(err)
	}
}

// 定义service层我需要Reader的函数--暂未修改
func getRouteRender(modelFile string) *template.Template {
	data, err := ioutil.ReadFile(modelFile)
	if nil != err {
		fmt.Println("read tplFile err:", err)
		return nil
	}
	render := template.Must(template.New("route").
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
