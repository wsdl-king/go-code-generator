{{$exportModelName := .ModelName | FirstCharUpper}}

package {{.PackageName}}
import "github.com/jinzhu/gorm"
type {{$exportModelName}} struct {
{{range .TableSchema}}
    {{.COLUMN_NAME | ExportColumn}} {{.DATA_TYPE | TypeConvert}} {{.COLUMN_NAME | JsonAndFormTags}} // {{.COLUMN_COMMENT}}
{{end}}}

{{if .HavePk}}
    //根据主键得到单一实体
    func Get{{$exportModelName}}ByPrimaryKey({{.GetPkColumn | PkConvert}} {{.GetPkDataType | TypeConvert}}) (*{{$exportModelName}}, error) {
    var {{.ModelName}}  {{$exportModelName}}
    err := db.Where("{{.GetPkColumn}} = ?", {{.GetPkColumn | PkConvert}} ).First(&{{.ModelName}}).Error
    if err != nil && err != gorm.ErrRecordNotFound {
    return nil, err
    }
    return &{{.ModelName}}, nil
    }
{{end}}

{{if .HavePk}}
    //根据主键通过条件编辑实体
    func Edit{{$exportModelName}}ByPrimaryKey({{.GetPkColumn | PkConvert}} {{.GetPkDataType | TypeConvert}} , maps interface{})  error {
    if err := db.Model(&{{$exportModelName}}{}).Where("{{.GetPkColumn}} = ?", {{.GetPkColumn | PkConvert}} ).Updates(maps).Error; err != nil {
    return err
    }
    return nil
    }
{{end}}

{{if .HavePk}}
    //根据主键删除实体
    func Delete{{$exportModelName}}ByPrimaryKey({{.GetPkColumn | PkConvert}} {{.GetPkDataType | TypeConvert}})  error {
    if err := db.Where("{{.GetPkColumn}} = ?", {{.GetPkColumn | PkConvert}}).Delete(&{{$exportModelName}}{}).Error; err != nil {
    return err
    }
    return nil
    }
{{end}}
{{if .HavePk}}
    //插入实体
    func Add{{$exportModelName}}({{.ModelName}}  *{{$exportModelName}})  error {
    if err := db.Create(&{{.ModelName}}).Error; err != nil {
    return err
    }
    return nil
    }
{{end}}


{{if .HavePk}}
    //根据条件获得分页实体集合
    func Find{{$exportModelName}}s(pageNum int, pageSize int, maps interface{}) ([]*{{$exportModelName}}, error) {
    var (
    {{.ModelName}}s  []*{{$exportModelName}}
    err  error
    )
    if pageSize > 0 && pageNum > 0 {
    err = db.Where( maps).Offset(pageSize*(pageNum-1)).Limit(pageSize).Find(&{{.ModelName}}s).Error
    } else {
    err = db.Where( maps).Find(&{{.ModelName}}s).Error
    }
    if err != nil && err != gorm.ErrRecordNotFound {
    return nil, err
    }
    return {{.ModelName}}s, nil
    }
{{end}}
