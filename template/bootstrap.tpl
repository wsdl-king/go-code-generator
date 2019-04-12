package main

import (
"github.com/gin-gonic/gin"
"go-code-generator/routes/api"
"net/http"
"time"
)
    func main() {
    engine := gin.New()
    //自定义
    srv := &http.Server{
    Addr:              ":8080",
    Handler:           engine,
    ReadTimeout:       10 * time.Second,
    ReadHeaderTimeout: 1 * time.Second,
    WriteTimeout:      10 * time.Second,
    }

{{range .AllTable}}
    {{if .TABLE_COMMENT}}
        //根据主键获得单一{{.TABLE_COMMENT}}
    {{ else }}
        //根据主键获得{{.TABLE_NAME}}
    {{ end }}
    engine.GET("/get{{.TABLE_NAME | FirstCharUpper }}/:customerId", routes.Get{{.TABLE_NAME | FirstCharUpper }}ByPrimaryKey)
    {{if.TABLE_COMMENT}}
        //获取分页{{.TABLE_COMMENT}}
    {{else}}
        //获取分页{{.TABLE_NAME}}
    {{end}}
    engine.POST("/find{{.TABLE_NAME | FirstCharUpper }}List", routes.GetAll{{.TABLE_NAME | FirstCharUpper }}s)
    {{if.TABLE_COMMENT}}
        //修改{{.TABLE_COMMENT}}
    {{else}}
        //修改{{.TABLE_NAME}}
    {{end}}
    engine.POST("/update{{.TABLE_NAME | FirstCharUpper }}", routes.Edit{{.TABLE_NAME | FirstCharUpper }}ByPrimaryKey)
    {{if.TABLE_COMMENT}}
        //删除{{.TABLE_COMMENT}}
    {{else}}
        //删除{{.TABLE_NAME}}
    {{end}}
    engine.DELETE("/delete{{.TABLE_NAME | FirstCharUpper }}", routes.Delete{{.TABLE_NAME | FirstCharUpper }}ByPrimaryKey)
    {{if.TABLE_COMMENT}}
        //新增{{.TABLE_COMMENT}}
    {{else}}
        //新增{{.TABLE_NAME}}
    {{end}}
    engine.POST("/Add{{.TABLE_NAME | FirstCharUpper }}", routes.Add{{.TABLE_NAME | FirstCharUpper }})
{{end}}
    srv.ListenAndServe()
}