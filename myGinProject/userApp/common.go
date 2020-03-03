package userApp

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/gin-contrib/sessions"
	"myGinProject/contestApp"
	"strconv"
	"sync"
	"time"
)

var (
	db *gorm.DB
	lock sync.Mutex
)

func Init(r *gin.Engine,_db *gorm.DB)  {
	db = _db
	urls_init(r)
	models_init()
}

const SESSSION_USERNAME_KEY  = "who"

//通过session获取当前登陆的用户
func GetCurrentUser(c *gin.Context) *User {
	session := sessions.Default(c)
	username := session.Get(SESSSION_USERNAME_KEY)
	if username == nil {
		return nil
	}
	var user User
	if db.Where("username = ?",username.(string)).First(&user).RecordNotFound() {
		return nil
	}
	db.Model(&user).Update("last_seen",time.Now())
	return &user
}

//判断是否登陆
func IsLogin(c *gin.Context) bool {
	if user:=GetCurrentUser(c); user==nil{
		return  false
	}
	return  true
}


func SetSessionKV(c *gin.Context,key string,val interface{})  {
	session := sessions.Default(c)
	session.Set(key,val)
	session.Save()
}

func DeleteSessionK(c *gin.Context, key string)  {
	session := sessions.Default(c)
	session.Delete(key)
	session.Save()
}

func SetSession(c *gin.Context,val string)  {
	session := sessions.Default(c)
	session.Set(SESSSION_USERNAME_KEY, val)
	session.Save()
}

func DeleteSession(c *gin.Context)  {
	session := sessions.Default(c)
	session.Delete(SESSSION_USERNAME_KEY)
	session.Save()
}
//认证登陆
func AuthLogin(c *gin.Context) {
	if !IsLogin(c) {
		c.String(401,"未登陆")
		c.Abort()
	}
}

func getUserAndContest(c *gin.Context) (*User,*contestApp.Contest) {
	user := GetCurrentUser(c)
	id,_ := strconv.Atoi(c.DefaultQuery("id","0"))
	contest := contestApp.GetContestByID(uint(id))
	return user,contest
}

func AuthContest(c *gin.Context) (*contestApp.Team,*contestApp.Contest)  {
	user,contest := getUserAndContest(c)
	if contest==nil {
		c.String(403,"不存在该比赛")
		return nil,nil
	}
	if user==nil && contest.Type =="private" {
		c.String(401,"未登录")
		return  nil,nil
	}
 	team := contest.GetOneTeamByUserID(user.ID)
	if team==nil && contest.Type =="private" {
		c.String(403,"没有注册比赛")
		return  nil,nil
	}
	return team,contest
}

func AuthContestPost(c *gin.Context) (*contestApp.Team,*contestApp.Contest)  {
	user := GetCurrentUser(c)
	id,_ := strconv.Atoi(c.DefaultPostForm("id","0"))
	contest := contestApp.GetContestByID(uint(id))
	if contest==nil {
		c.String(403,"不存在该比赛")
		return nil,nil
	}
	if user==nil && contest.Type =="private" {
		c.String(401,"未登录")
		return  nil,nil
	}
	team := contest.GetOneTeamByUserID(user.ID)
	if team==nil && contest.Type =="private" {
		c.String(403,"没有注册比赛")
		return  nil,nil
	}
	return team,contest
}

func SearchUsers(l,r int,rules map[string]interface{}) ([]User,int,error) {
	var users []User
	err := db.Where(rules).Find(&users).Error
	if err!=nil {
		return []User{},0,err
	}
	cnt := len(users)
	if l>cnt {
		return []User{},cnt,nil
	}
	if r > cnt {
		r = cnt
	}
	return users[l-1:r],cnt,nil
}

