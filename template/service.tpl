{{$exportModelName := .ModelName | FirstCharUpper}}
package  service
import "go-code-generator/model"

type {{$exportModelName}}Service struct {
    {{$exportModelName}} *model.{{$exportModelName}}
    PageNum  int  `form:"pageNum" json:"pageNum"`
    PageSize int  `form:"pageSize" json:"pageSize"`
}


{{if .HavePk}}
    //获得单一实体
    func ({{.ModelName}} *{{$exportModelName}}Service) Get{{$exportModelName}}ByPrimaryKey() (*model.{{$exportModelName}}, error) {
    {{.ModelName}}Res, err := model.Get{{$exportModelName}}ByPrimaryKey({{.ModelName}}.{{$exportModelName}}.{{.GetPkColumn | PkConvert | FirstCharUpper}})
    if err != nil {
    return nil, err
    }
    return  {{.ModelName}}Res, nil
    }
{{end}}


{{if .HavePk}}
    //分页获取实体列表
    func ({{.ModelName}} *{{$exportModelName}}Service) GetAll{{$exportModelName}}s() ([]*model.{{$exportModelName}}, error) {
    {{.ModelName}}s, err := model.Find{{$exportModelName}}s({{.ModelName}}.PageNum, {{.ModelName}}.PageSize, {{.ModelName}}.{{$exportModelName}})
    if err != nil {
    return nil, err
    }
    return  {{.ModelName}}s, nil
    }
{{end}}


{{if .HavePk}}
    //删除单一实体
    func ({{.ModelName}} *{{$exportModelName}}Service) Delete{{$exportModelName}}ByPrimaryKey()  error  {
    return model.Delete{{$exportModelName}}ByPrimaryKey({{.ModelName}}.{{$exportModelName}}.{{.GetPkColumn | PkConvert | FirstCharUpper}})
    }
{{end}}

{{if .HavePk}}
    //编译单一实体
    func ({{.ModelName}} *{{$exportModelName}}Service) Edit{{$exportModelName}}ByPrimaryKey()  error  {
    return model.Edit{{$exportModelName}}ByPrimaryKey({{.ModelName}}.{{$exportModelName}}.{{.GetPkColumn | PkConvert | FirstCharUpper}},{{.ModelName}}.{{$exportModelName}})
    }
{{end}}

{{if .HavePk}}
    //新增单一实体
    func ({{.ModelName}} *{{$exportModelName}}Service) Add{{$exportModelName}}()  error  {
    return model.Add{{$exportModelName}}({{.ModelName}}.{{$exportModelName}})
    }
{{end}}
