package main

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"myGinProject/adminApp"
	"myGinProject/common"
	"myGinProject/contestApp"
	"myGinProject/problemApp"
	"myGinProject/userApp"
	"net/http"
	"os"
)

func init()  {

	var err error

	common.DB , err = gorm.Open("mysql", "root:root@(localhost:3306)/online_judge?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	common.DB.SingularTable(true)
	defer common.DB.Close()

	//启用session
	common.R = gin.Default()


	store := cookie.NewStore([]byte("secert"))

	common.R.Use(sessions.Sessions("ginSession",store))


	//初始化所以app
	userApp.Init(common.R,common.DB)
	problemApp.Init(common.R,common.DB)
	adminApp.Init(common.R)
	contestApp.Init(common.R,common.DB)


	common.R.LoadHTMLGlob("./templates/**/*")
	//前端
	common.R.StaticFS("/js",http.Dir("./frontend/js"))
	common.R.StaticFS("/css",http.Dir("./frontend/css"))
	common.R.StaticFS("/fonts",http.Dir("./frontend/fonts"))
	common.R.StaticFS("/img",http.Dir("./frontend/img"))
	common.R.StaticFile("favicon.ico", "./frontend/favicon.ico")
	common.R.GET("/", func(c *gin.Context) {
		c.HTML(200,"index.html",nil)
	})

	//后台
	common.R.StaticFS("admin/js",http.Dir("./vueAdmin/js"))
	common.R.StaticFS("admin/css",http.Dir("./vueAdmin/css"))
	common.R.StaticFS("admin/fonts",http.Dir("./vueAdmin/fonts"))
	common.R.StaticFS("admin/img",http.Dir("./vueAdmin/img"))
	common.R.StaticFile("admin/favicon.ico", "./vueAdmin/favicon.ico")
	common.R.GET("/admin", func(c *gin.Context) {
		c.HTML(200,"index2.html",nil)
	})
	problemApp.Ping()
	common.R.Run(":8888")
}

func main() {

	fmt.Println("系统关闭...")
}
