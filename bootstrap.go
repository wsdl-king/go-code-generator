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

	//根据主键获得单一文章管理

	engine.GET("/getArticle/:key", routes.GetArticleByPrimaryKey)

	//获取分页文章管理

	engine.POST("/findArticleList", routes.GetAllArticles)

	//修改文章管理

	engine.POST("/updateArticle", routes.EditArticleByPrimaryKey)

	//删除文章管理

	engine.DELETE("/deleteArticle", routes.DeleteArticleByPrimaryKey)

	//新增文章管理

	engine.POST("/AddArticle", routes.AddArticle)

	//根据主键获得auth

	engine.GET("/getAuth/:key", routes.GetAuthByPrimaryKey)

	//获取分页auth

	engine.POST("/findAuthList", routes.GetAllAuths)

	//修改auth

	engine.POST("/updateAuth", routes.EditAuthByPrimaryKey)

	//删除auth

	engine.DELETE("/deleteAuth", routes.DeleteAuthByPrimaryKey)

	//新增auth

	engine.POST("/AddAuth", routes.AddAuth)

	//根据主键获得单一商户表

	engine.GET("/getCustomer/:key", routes.GetCustomerByPrimaryKey)

	//获取分页商户表

	engine.POST("/findCustomerList", routes.GetAllCustomers)

	//修改商户表

	engine.POST("/updateCustomer", routes.EditCustomerByPrimaryKey)

	//删除商户表

	engine.DELETE("/deleteCustomer", routes.DeleteCustomerByPrimaryKey)

	//新增商户表

	engine.POST("/AddCustomer", routes.AddCustomer)

	//根据主键获得msg

	engine.GET("/getMsg/:key", routes.GetMsgByPrimaryKey)

	//获取分页msg

	engine.POST("/findMsgList", routes.GetAllMsgs)

	//修改msg

	engine.POST("/updateMsg", routes.EditMsgByPrimaryKey)

	//删除msg

	engine.DELETE("/deleteMsg", routes.DeleteMsgByPrimaryKey)

	//新增msg

	engine.POST("/AddMsg", routes.AddMsg)

	//根据主键获得单一文章标签管理

	engine.GET("/getTag/:key", routes.GetTagByPrimaryKey)

	//获取分页文章标签管理

	engine.POST("/findTagList", routes.GetAllTags)

	//修改文章标签管理

	engine.POST("/updateTag", routes.EditTagByPrimaryKey)

	//删除文章标签管理

	engine.DELETE("/deleteTag", routes.DeleteTagByPrimaryKey)

	//新增文章标签管理

	engine.POST("/AddTag", routes.AddTag)

	srv.ListenAndServe()
}
