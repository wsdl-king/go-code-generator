{{$exportModelName := .ModelName | FirstCharUpper}}
{{$context := "c *gin.Context"}}

package  routes
import ("go-code-generator/model"
"github.com/Unknwon/com"
"github.com/gin-gonic/gin"
"go-code-generator/service"
"go-code-generator/result"
)

{{if .HavePk}}
    //获得单一实体
    func Get{{$exportModelName}}ByPrimaryKey({{$context}}) {
    id := com.StrTo(c.Param("key")).MustInt()
    {{.ModelName}} := &model.{{$exportModelName}}{ {{.GetPkColumn | PkConvert | FirstCharUpper}}: id}
    {{.ModelName}}Service :=  service.{{$exportModelName}}Service{ {{$exportModelName}}:{{.ModelName}}}
    {{.ModelName}}Res,e:={{.ModelName}}Service.Get{{$exportModelName}}ByPrimaryKey()
    if e != nil {
    c.JSON(500, result.ErrorResult({{.ModelName}}Res,50000))
    }
    c.JSON(200, result.SuccessResult({{.ModelName}}Res))
    }
{{end}}

{{if .HavePk}}
    //分页获取实体列表
    func  GetAll{{$exportModelName}}s({{$context}}) {
    {{.ModelName}}:= &model.{{$exportModelName}}{}
    pageNum:= com.StrTo(c.Query("pageNum")).MustInt()
    pageSize:= com.StrTo(c.Query("pageSize")).MustInt()
    c.Bind({{.ModelName}})
    {{.ModelName}}Service :=  service.{{$exportModelName}}Service{
    PageNum:pageNum,
    PageSize:pageSize,
    {{$exportModelName}}:{{.ModelName}},
    }
    {{.ModelName}}Res,e:={{.ModelName}}Service.GetAll{{$exportModelName}}s()
    if e != nil {
    c.JSON(500, result.ErrorResult({{.ModelName}}Res,50000))
    }
    c.JSON(200, result.SuccessResult({{.ModelName}}Res))
    }
{{end}}

{{if .HavePk}}
    //获得单一实体
    func  Edit{{$exportModelName}}ByPrimaryKey({{$context}}){
    {{.ModelName}}:= &model.{{$exportModelName}}{}
    c.Bind({{.ModelName}})
    {{.ModelName}}Service :=  service.{{$exportModelName}}Service{
    {{$exportModelName}}:{{.ModelName}},
    }
    e:={{.ModelName}}Service.Edit{{$exportModelName}}ByPrimaryKey()
    if e != nil {
    c.JSON(500, result.ErrorResult("",50000))
    }
    c.JSON(200, result.SuccessResult(""))
    }
{{end}}

{{if .HavePk}}
    //获得单一实体
    func  Delete{{$exportModelName}}ByPrimaryKey({{$context}}){
    id := com.StrTo(c.Param("{{.GetPkColumn | PkConvert}}")).MustInt()
    {{.ModelName}} := &model.{{$exportModelName}}{ {{.GetPkColumn | PkConvert | FirstCharUpper}}: id}
    {{.ModelName}}Service :=  service.{{$exportModelName}}Service{ {{$exportModelName}}:{{.ModelName}}}
    e:={{.ModelName}}Service.Delete{{$exportModelName}}ByPrimaryKey()
    if e != nil {
    c.JSON(500, result.ErrorResult("",50000))
    }
    c.JSON(200, result.SuccessResult(""))
    }
{{end}}
{{if .HavePk}}
    //新增实体
    func  Add{{$exportModelName}}({{$context}}){
    {{.ModelName}}:= &model.{{$exportModelName}}{}
    c.Bind({{.ModelName}})
    {{.ModelName}}Service :=  service.{{$exportModelName}}Service{
    {{$exportModelName}}:{{.ModelName}},
    }
    e:={{.ModelName}}Service.Add{{$exportModelName}}()
    if e != nil {
    c.JSON(500, result.ErrorResult("",50000))
    }
    c.JSON(200, result.SuccessResult(""))
    }
{{end}}