package router

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"net/http"
)
import "blog/controller"

func Router() {
	router := gin.Default()

	//配置静态映射
	router.Static("/assets", "./assets")
	router.StaticFS("/static", http.Dir("./static"))

	//new template engine
	router.HTMLRender = gintemplate.New(gintemplate.TemplateConfig{
		Root:         "template/fontend",
		Extension:    ".html",
		DisableCache: true,
	})

	router.GET("/", controller.Index)
	router.GET("/loginindex", controller.LoginIndex)
	router.GET("/signupindex", controller.SignUpIndex)

	router.Run() // listen and serve on 0.0.0.0:8080
}
