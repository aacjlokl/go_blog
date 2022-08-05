package controller

import (
	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

//处理get请求
func Index(c *gin.Context) {
	gintemplate.HTML(c, 200, "index", nil)
}

func LoginIndex(c *gin.Context) {
	gintemplate.HTML(c, 200, "login", nil)
}

func SignUpIndex(c *gin.Context) {
	gintemplate.HTML(c, 200, "signup", nil)
}
