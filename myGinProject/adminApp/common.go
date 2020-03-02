package adminApp

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/gin-contrib/sessions"
	"os"
)


var (
	db *gorm.DB
)

func Init(r *gin.Engine)  {
	var err error
	db,err = gorm.Open("mysql", "root:root@(localhost:3306)/goadmin?charset=utf8&parseTime=True&loc=Local")
	if err!=nil {
		fmt.Println("打开admin数据库失败")
		os.Exit(-1)
	}
	db.SingularTable(true)
	urls_init(r)
	models_init()
}


const SESSSION_ADMINNAME_KEY  = "ADMIN_NAME"

func GetAdminByName(name string)  *Admin{
	var admin Admin
	if db.Where("name = ?",name).First(&admin).RecordNotFound() {
		return nil
	}
	return &admin
}

//通过session获取当前登陆的用户
func GetAdminBySession(c *gin.Context) *Admin {
	session := sessions.Default(c)
	name := session.Get(SESSSION_ADMINNAME_KEY)
	if name == nil {
		return nil
	}
	var admin Admin
	if db.Where("name = ?",name.(string)).First(&admin).RecordNotFound() {
		return nil
	}
	return &admin
}


//判断是否登陆
func IsLogin(c *gin.Context) bool {
	if admin:=GetAdminBySession(c); admin==nil{
		return  false
	}
	return  true
}


func SetSession(c *gin.Context,val string)  {
	session := sessions.Default(c)
	session.Set(SESSSION_ADMINNAME_KEY, val)
	session.Save()
}

func DeleteSession(c *gin.Context)  {
	session := sessions.Default(c)
	session.Delete(SESSSION_ADMINNAME_KEY)
	session.Save()
}


func Authorrize(c *gin.Context) {
	if !IsLogin(c) {
		c.String(401,"")
		c.Abort()
	}
}


