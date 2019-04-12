package generate

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"html/template"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

type TABLE_SCHEMA struct {
	//列名
	COLUMN_NAME string `db:"COLUMN_NAME" json:"column_name"`
	//数据类型
	DATA_TYPE string `db:"DATA_TYPE" json:"data_type"`
	//主键pri
	COLUMN_KEY string `db:"COLUMN_KEY" json:"column_key"`
	// 注释
	COLUMN_COMMENT string `db:"COLUMN_COMMENT" json:"COLUMN_COMMENT"`
}

type ALLTABLE struct {
	//指定库中对应的表名称
	TABLE_NAME string `db:"TABLE_NAME" json:"table_name"`
	// 指定库中对应的表注释
	TABLE_COMMENT string `db:"TABLE_COMMENT" json:"table_comment"`
}

// 此结构体是用于渲染frameWork
type ModelInfo struct {
	BDName       string
	DBConnection string
	TableName    string
	PackageName  string
	ModelName    string
	// 表属性集合
	TableSchema *[]TABLE_SCHEMA
	// 表集合
	AllTable *[]ALLTABLE
}

func start(db *sqlx.DB, tableName, dbName string) {
	TableSchema := &[]TABLE_SCHEMA{}
	err := db.Select(TableSchema,
		"SELECT COLUMN_NAME, DATA_TYPE,COLUMN_KEY,COLUMN_COMMENT from information_schema.COLUMNS where "+
			"TABLE_NAME"+"='"+tableName+"' and "+"table_schema = '"+dbName+"'")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("成功映射tableSchema", TableSchema)
	if len(*TableSchema) <= 0 {
		fmt.Println(tableName, "tableSchema is null")
		return
	}
	model := &ModelInfo{
		PackageName:  "model",
		BDName:       dbName,
		DBConnection: dbName,
		TableName:    tableName,
		ModelName:    tableName,
		TableSchema:  TableSchema}
	// 根据modelInfo渲染model
	RenderModel(model, tableName)
	// 根据modelInfo渲染Service
	RenderService(model, tableName)
	// 根据modelInfo渲染Route
	RenderRoute(model, tableName)
	// -w指定目录或者文件
	cmd := exec.Command("gofmt", "-w", "./")
	cmd.Run()
}

// 建立路由
func buildRoutes(allTables *[]ALLTABLE) {
	model := &ModelInfo{AllTable: allTables}
	data, err := ioutil.ReadFile("./template/bootstrap.tpl")
	if nil != err {
		fmt.Println("read tplFile err:", err)
		return
	}
	render := template.Must(template.New("bootstrap").
		Funcs(template.FuncMap{
			"FirstCharUpper": FirstCharUpper,
			"TypeConvert":    TypeConvert,
			"ExportColumn":   ExportColumn,
		}).
		Parse(string(data)))
	fileName := "bootstrap.go"
	os.Remove(fileName)
	f, err := os.Create(fileName)
	defer f.Close()
	render.Execute(f, model)
	cmd := exec.Command("gofmt", "-w", fileName)
	cmd.Run()
	fmt.Println("生成成功!")
}

func (m *ModelInfo) TableNames() []string {
	result := make([]string, 0, len(*m.AllTable))
	for _, t := range *m.AllTable {
		result = append(result, t.TABLE_NAME)
	}
	return result
}
func (m *ModelInfo) TableComments() []string {
	result := make([]string, 0, len(*m.AllTable))
	for _, t := range *m.AllTable {
		result = append(result, t.TABLE_COMMENT)
	}
	return result
}

func (m *ModelInfo) ColumnNames() []string {
	result := make([]string, 0, len(*m.TableSchema))
	for _, t := range *m.TableSchema {
		result = append(result, t.COLUMN_NAME)
	}
	return result
}

func (m *ModelInfo) ColumnCount() int {
	return len(*m.TableSchema)
}

func (m *ModelInfo) PkColumnsSchema() []TABLE_SCHEMA {
	result := make([]TABLE_SCHEMA, 0, len(*m.TableSchema))
	for _, t := range *m.TableSchema {
		if t.COLUMN_KEY == "PRI" {
			result = append(result, t)
		}
	}
	return result
}

// 函数记得注册 = =
// a_b_c  aBC
func PkConvert(column string) string {
	columnItems := strings.Split(column, "_")
	for i := 1; i < len(columnItems); i++ {
		item := strings.Title(columnItems[i])
		columnItems[i] = item
	}
	return strings.Join(columnItems, "")
}

func (m *ModelInfo) GetPkColumn() string {
	if m.HavePk() {
		return m.PkColumnsSchema()[0].COLUMN_NAME
	}
	return ""
}

// 得到主键的类型
func (m *ModelInfo) GetPkDataType() string {
	if m.HavePk() {
		return m.PkColumnsSchema()[0].DATA_TYPE
	}
	return ""
}

func (m *ModelInfo) HavePk() bool {
	return len(m.PkColumnsSchema()) > 0
}

func (m *ModelInfo) NoPkColumnsSchema() []TABLE_SCHEMA {
	result := make([]TABLE_SCHEMA, 0, len(*m.TableSchema))
	for _, t := range *m.TableSchema {
		if t.COLUMN_KEY != "PRI" {
			result = append(result, t)
		}
	}
	return result
}

func (m *ModelInfo) NoPkColumns() []string {
	noPkColumnsSchema := m.NoPkColumnsSchema()
	result := make([]string, 0, len(noPkColumnsSchema))
	for _, t := range noPkColumnsSchema {
		result = append(result, t.COLUMN_NAME)
	}
	return result
}

func (m *ModelInfo) PkColumns() []string {
	pkColumnsSchema := m.PkColumnsSchema()
	result := make([]string, 0, len(pkColumnsSchema))
	for _, t := range pkColumnsSchema {
		result = append(result, t.COLUMN_NAME)
	}
	return result
}

func IsUUID(str string) bool {
	return "uuid" == str
}

func FirstCharLower(str string) string {
	if len(str) > 0 {
		return strings.ToLower(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

// 首字母大写函数映射
func FirstCharUpper(str string) string {
	if len(str) > 0 {
		return strings.ToUpper(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

func JsonAndFormTags(columnName string) template.HTML {
	return template.HTML("`form:" + `"` + columnName + `"` +
		" json:" + `"` + columnName + "\"`")
}

// ForExample  test_id--->TestId
func ExportColumn(columnName string) string {
	columnItems := strings.Split(columnName, "_")
	columnItems[0] = FirstCharUpper(columnItems[0])
	for i := 0; i < len(columnItems); i++ {
		item := strings.Title(columnItems[i])
		columnItems[i] = item
	}
	return strings.Join(columnItems, "")
}

func TypeConvert(str string) string {
	switch str {
	case "smallint", "tinyint":
		return "int8"
	case "varchar", "text", "longtext", "char":
		return "string"
	case "date":
		return "string"
	case "int":
		return "int"
	case "timestamp", "datetime":
		return "time.Time"
	case "bigint":
		return "int64"
	case "float", "double", "decimal":
		return "float64"
	case "uuid":
		return "gocql.UUID"
	default:
		return str
	}
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func ColumnAndType(table_schema []TABLE_SCHEMA) string {
	result := make([]string, 0, len(table_schema))
	for _, t := range table_schema {
		result = append(result, t.COLUMN_NAME+" "+TypeConvert(t.DATA_TYPE))
	}
	return strings.Join(result, ",")
}

func ColumnWithPostfix(columns []string, Postfix, sep string) string {
	result := make([]string, 0, len(columns))
	for _, t := range columns {
		result = append(result, t+Postfix)
	}
	return strings.Join(result, sep)
}

func MakeQuestionMarkList(num int) string {
	a := strings.Repeat("?,", num)
	return a[:len(a)-1]
}

// 定义model层我需要Reader的函数

func createFile(fileName string, filePath string) *os.File {
	f := strings.ToLower(fileName) + ".go"
	os.Remove(f)
	file, _ := os.Create(filePath + f)
	return file
}
