package userApp

// 0 未登陆
// 1 操作成功
// 2 请求失败，并返回错误信息
import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"myGinProject/common"
	"myGinProject/contestApp"
	"myGinProject/problemApp"
	"path"
	"sort"
	"strconv"
	"time"
)


//创建用户
func register(c *gin.Context) {
	if IsLogin(c) {
		c.String(403, "请退出当前用户")
		return
	}
	var form RegisterBind
	if err := c.ShouldBind(&form); err != nil {
		c.String(400, err.Error())
		return
	}

	if ok, errInfo := common.Validate(form); !ok {
		c.String(403, errInfo)
		return
	}

	if ok,errInfo := form.LoginBind.CharsValidate(); !ok {
		c.String(403, errInfo)
		return
	}

	var user User
	if !db.Where("username = ?", form.Username).First(&user).RecordNotFound() {
		c.String(403, "用户名已存在")
		return
	}
	if !db.Where("email = ?", form.Email).First(&user).RecordNotFound() {
		c.String(403, "邮箱已被注册")
		return
	}
	//创建表
	user = User {
		LastSeen:    time.Now(),
		Username:    form.Username,
		Password:    form.Password,
		Email:       form.Email,
		School:      form.School,
		Description: form.Description,
	}

	if err:=db.Create(&user).Error; err!=nil {
		c.String(500,"服务器内部错误")
		panic(err)

	}
	SetSession(c, user.Username)
	c.String(200, "注册成功")
}

//登陆
func login(c *gin.Context) {

	if IsLogin(c) {
		DeleteSession(c)
	}

	var form LoginBind
	if err := c.ShouldBind(&form); err != nil {
		c.String(400, "参数无法解析")
		return
	}

	if ok, errInfo := common.Validate(form); !ok {
		c.String(403, errInfo)
		return
	}

	if ok,errInfo := form.CharsValidate(); !ok {
		c.String(403, errInfo)
		return
	}

	var user User
	if db.Where("username = ?", form.Username).First(&user).RecordNotFound() {
		c.String(403, "用户名不存在")
		return
	}
	if user.Password != form.Password {
		c.String(403, "密码错误")
		return
	}
	SetSession(c, user.Username)
	c.String(200, "登陆成功")
}

//退出登陆
func logout(c *gin.Context)  {
	DeleteSession(c)
	c.String(200,"退出成功")
}

//个人信息
func profile(c *gin.Context)  {
	user := GetCurrentUser(c)
	if user == nil {
		c.String(403,"未登录")
		return
	}
	c.JSON(200,user)
}

func getUserInfo(c *gin.Context)  {
	var user User
	if db.Where("username = ?",c.Query("username")).First(&user).RecordNotFound() {
		c.String(403,"不存在该用户")
		return
	}
	info := map[string]interface{} {
		"username": user.Username,
		"school": user.School,
		"email": user.Email,
		"desc": user.Description,
		"create_at": user.CreatedAt.Format("2006-01-02 15:04:05"),
		"last_seen": user.LastSeen.Format("2006-01-02 15:04:05"),
	}
	solved,unsolved,status := user.GetStates()
	c.JSON(200,gin.H{
		"info": info,
		"solved": solved,
		"unsolved": unsolved,
		"status": status,
	})
}
func searchUsers(c *gin.Context)  {
	l,_ := strconv.Atoi(c.DefaultPostForm("l","1"))
	r,_ := strconv.Atoi(c.DefaultPostForm("r","50"))
	var rules map[string]interface{}
	_ = json.Unmarshal([]byte(c.DefaultPostForm("rules","{}")),&rules)
	users,tot,err := SearchUsers(l,r,rules)
	if err!=nil {
		c.String(403,"查询错误")
		return
	}
	var us UserRanks
	for _,item := range users {
		solved,all := item.GetACAndAll()
		var ratio float32 = 0
		if all>0 {
			ratio = float32(solved)/float32(all)
		}
		us = append(us, UsersRank{
			Username: item.Username,
			School:   item.School,
			Solved:   uint(solved),
			All:      uint(all),
			Ratio:    ratio,
		})
	}
	sort.Sort(us)
	c.JSON(200,gin.H{
		"tot": tot,
		"data": us,
	})
}
//更新用户信息
func update(c *gin.Context)  {
	user := GetCurrentUser(c)
	if user == nil {
		c.String(403, "未登陆")
		return
	}
	var form UpdateBind
	if err := c.ShouldBind(&form); err!=nil {
		c.String(400, err.Error())
		return
	}

	if ok, errInfo := common.Validate(form); !ok {
		c.String(403, errInfo)
		return
	}
	if ok, errInfo := form.LoginBind.CharsValidate(); !ok {
		c.String(403, errInfo)
		return
	}

	if user.Password != form.Password {
		c.String(403,"密码错误")
		return
	}

	if user.Username != form.Username {
		if !db.Where("username = ?", form.Username).First(&User{}).RecordNotFound() {
			c.String(403, "用户名已存在")
			return
		}
		user.Username = form.Username
	}
	if user.Email != form.Email {
		if !db.Where("email = ?", form.Email).First(&User{}).RecordNotFound() {
			c.String(403, "邮箱已被注册")
			return
		}
		user.Email = form.Email
	}
	user.School = form.School
	user.Description = form.Description

	if err := db.Model(&user).Updates(user).Error;err !=nil{
			c.String(500,"服务器内部错误")
			panic(err)
		}
	SetSession(c,user.Username)
	c.String(200, "修改成功")
}

func updatePassword(c *gin.Context)  {
	user := GetCurrentUser(c)
	if user == nil {
		c.String(403, "未登陆")
		return
	}
	var form UpdatePasswordBind
	if err := c.ShouldBind(&form); err!=nil {
		c.String(400, err.Error())
		return
	}

	if ok, errInfo := common.Validate(form); !ok {
		c.String(403, errInfo)
		return
	}

	if user.Password != form.OldPassword {
		c.String(403,"密码错误")
		return
	}

	if err:=db.Model(&user).Update("password",form.NewPassword).Error; err!=nil{
		c.String(500,"服务器内部错误")
		panic(err)
	}
	c.String(200,"修改成功")
}

func autologin(c *gin.Context)  {
	user := GetCurrentUser(c)
	if user == nil {
		c.String(403,"")
		return
	}
	c.String(200,user.Username)
}

func getProblems(c *gin.Context)  {
	l,_ := strconv.Atoi(c.Query("l"))
	r,_ := strconv.Atoi(c.Query("r"))
	problems,count,err := problemApp.GetProblems(l,r,true,false)
	if err!=nil {
		c.String(403,err.Error())
		return
	}
	var data [] map[string]interface{}
	for	i:=len(*problems)-1;i>=0;i-- {
		data = append(data, gin.H{
			"index": (*problems)[i].Index,
			"title": (*problems)[i].Title,
			"solved": 0, //0没有尝试过，1AC,2错误
			"ac": 0,
			"all":0,
			"ratio":0,
		})
	}
	c.JSON(200,gin.H{
		"data": data,
		"tot": count,
	})
}
func searchProblems(c *gin.Context)  {
	l,_ := strconv.Atoi(c.DefaultPostForm("l","1"))
	r,_ := strconv.Atoi(c.DefaultPostForm("r","100"))
	var rules map[string]interface{}
	_ = json.Unmarshal([]byte(c.DefaultPostForm("rules","{}")),&rules)
	problems,tot,err := problemApp.SearchProblems(l,r,rules)
	if err!=nil {
		c.String(403,"查找错误")
		return
	}
	user := GetCurrentUser(c)
	var data [] map[string]interface{}
	for _,item := range problems {
		solved := 0
		if user!=nil {
			solved = user.GetProblemState(&item)
		}
		data = append(data, gin.H{
			"index": item.Index,
			"title": item.Title,
			"solved": solved, //0没有尝试过，1AC,2错误
			"ac": item.AC,
			"all":item.All,
			"ratio":item.GetRatio(),
		})
	}
	c.JSON(200,gin.H{
		"data": data,
		"tot": tot,
	})
}

func getOneProblem(c *gin.Context)  {
	index := c.Query("index")
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403,"找不到这道题")
		return
	}
	problem.ChangeStringToExamples()
	filePath := ""
	if problem.FileName!="" {
		filePath = path.Join(problem.FileDIR,problem.FileName)
	}
	code := ""
	lang := "C++11"
	user := GetCurrentUser(c)
	if user!=nil {
		s := user.GetOneSubmission(problem)
		if s!=nil {
			code = s.Code
			lang = s.Lang
		}
	}
	c.JSON(200,gin.H{
		"info": problem,
		"file_path": filePath,
		"code": code,
		"lang": lang,
	})
}

//code,lang,index
func submitCode(c *gin.Context)  {
	user := GetCurrentUser(c)
	if user == nil {
		c.String(401,"未登陆")
		return
	}
	index := c.PostForm("index")
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403,"找不到这道题")
		return
	}
	submission := problemApp.Submission{
		Code:      c.PostForm("code"),
		Length:    uint(len(c.PostForm("code"))),
		Lang:      c.PostForm("lang"),
		ProblemID: problem.ID,
		UserID:    user.ID,
		ProblemIndex:  problem.Index,
		Author: user.Username,
	}
	if err:=submission.Create();err!=nil{
		c.String(403,err.Error())
		return
	}
	go problemApp.Judge(&submission,problem)
	c.String(200,"提交成功")
}

func getHistory(c *gin.Context)  {
	user := GetCurrentUser(c)
	if user == nil {
		c.String(401,"未登陆")
		return
	}
	index := c.DefaultQuery("index","0")
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403,"找不到这道题")
		return
	}
	ss,_ := user.GetSubmissions(map[string]interface{}{"problem_id":problem.ID})
	var data []map[string]interface{}
	for _,item := range *ss {
		data = append(data, map[string]interface{} {
			"run_id":item.ID,
			"created_at": item.CreatedAt.Format("2006-01-02 15:04:05"),
			"lang": item.Lang,
			"status": item.Status,
			"time": item.Time,
			"memory": item.Memory,
		})
	}
	c.JSON(200,gin.H{
		"num": len(data),
		"data" : data,
	})
}

func searchSubmissions(c *gin.Context)  {
	l,_ := strconv.Atoi(c.DefaultPostForm("l","1"))
	r,_ := strconv.Atoi(c.DefaultPostForm("r","50"))
	var rules map[string]interface{}
	_ = json.Unmarshal([]byte(c.DefaultPostForm("rules","{}")),&rules)
	submissions,tot,err:=problemApp.SearchSubmissions(l,r,rules)
	if err!=nil {
		c.String(403,"查询错误")
		return
	}
	var data []map[string]interface{}
	for _,item:= range submissions {
		data=append(data, map[string]interface{}{
			"id": item.ID,
			"author": item.Author,
			"status": item.Status,
			"index": item.ProblemIndex,
			"time": item.Time,
			"memory": item.Memory,
			"length": item.Length,
			"lang": item.Lang,
			"submitTime": item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(200,gin.H{
		"tot": tot,
		"data": data,
	})
}

func getSubmission(c *gin.Context)  {
	id,_ := strconv.Atoi(c.DefaultQuery("id","-1"))
	if id==-1 {
		c.String(403,"错误")
		return
	}
	var submission problemApp.Submission
	if db.First(&submission,id).RecordNotFound() {
		c.String(403,"找不到提交")
		return
	}
	user := GetCurrentUser(c)
	if user==nil || user.ID != submission.UserID {
		c.String(403,"无权查看该提交")
		return
	}
	c.JSON(200,gin.H{
		"author": submission.Author,
		"status": submission.Status,
		"index": submission.ProblemIndex,
		"time": submission.Time,
		"memory": submission.Memory,
		"length": submission.Length,
		"lang": submission.Lang,
		"submitTime": submission.CreatedAt.Format("2006-01-02 15:04:05"),
		"compile_info": submission.CompileInfo,
		"code": submission.Code,
	})
}

func searchContests(c *gin.Context)  {
	l,_ := strconv.Atoi(c.DefaultPostForm("l","1"))
	r,_ := strconv.Atoi(c.DefaultPostForm("r","50"))
	var rules map[string]interface{}
	_ = json.Unmarshal([]byte(c.DefaultPostForm("rules","{}")),&rules)
	contests,tot,err := contestApp.SearchContests(l,r,rules)
	if err!=nil {
		c.String(403,"查询错误")
		return
	}
	var data []map[string]interface{}
	for _,item := range contests {
		data=append(data,map[string]interface{}{
			"id": item.ID,
			"title": item.Title,
			"type": item.Type,
			"format": item.Format,
			"begin": item.Begin.Format("2006-01-02 15:04:05"),
			"length": float64(item.Length)/3600,
			"status": item.Status,
			"num": item.GetTeamsCount(),
		})
	}
	c.JSON(200,gin.H{
		"tot": tot,
		"data": data,
	})
}

func enterContest(c *gin.Context)  {
	user,contest := getUserAndContest(c)
	if contest==nil {
		c.String(403,"不存在该比赛")
		return
	}
	if user==nil && contest.Type =="private" {
		c.String(401,"未登录")
		return
	}
	team := contest.GetOneTeamByUserID(user.ID)
	if team==nil && contest.Type =="private" {
		c.JSON(200,gin.H{
			"ok" : false,
		})
		return
	}
	c.JSON(200,gin.H{
		"ok" : true,
	})
}

func checkContestPwd(c *gin.Context)  {
	user := GetCurrentUser(c)
	if user==nil {
		c.String(401,"未登陆")
		return
	}
	id,_ := strconv.Atoi(c.PostForm("id"))
	contest := contestApp.GetContestByID(uint(id))
	if contest==nil {
		c.String(403,"找不到比赛")
		return
	}
	password := c.PostForm("password")
	if contest.Password != password {
		c.String(403,"错误密码")
		return
	}
	if contest.Status != "Ended" {
		team := &contestApp.Team {
			Name:      user.Username,
			ContestID: contest.ID,
			UserID:    user.ID,
			Order:     contest.Format,
		}
		team.Create()
	}
	c.String(200,"ok")
}

func getContestContent(c *gin.Context)  {
	_,contest := AuthContest(c)
	if contest==nil {
		return
	}
	data := make(map[string]interface{})
	data["contest"] = map[string]interface{}{
		"id": contest.ID,
		"title": contest.Title,
		"type": contest.Type,
		"format": contest.Format,
		"begin": contest.Begin.Format("2006-01-02 15:04:05"),
		"end": contest.End.Format("2006-01-02 15:04:05"),
		"length": float64(contest.Length)/3600,
		"status": contest.Status,
		"author": contest.Author,
		"desc": contest.Desc,
		"num": contest.GetTeamsCount(),
	}
	c.JSON(200,data)
}

func getCproblems(c *gin.Context)  {
	team,contest := AuthContest(c)
	if contest==nil {
		return
	}
	cproblems,err := contest.GetCproblems()
	if err!=nil {
		c.String(403,err.Error())
		return
	}
	var data []map[string]interface{}
	for _,cp := range *cproblems {
		solved := 0
		if team != nil {
			if ps:=team.GetOneProblemStatus(&cp);ps!=nil {
				if ps.Minutes > 0 {
					solved = 1
				} else {
					solved = 2
				}
			}
		}
		var ratio float64 = 0
		if cp.All>0 {
			ratio = float64(cp.AC)/float64(cp.All)
		}
		data = append(data, map[string]interface{}{
			"solved": solved,
			"label": cp.Label,
			"title": cp.Title,
			"ac": cp.AC,
			"all": cp.All,
			"ratio": ratio*100,
		})
	}
	c.JSON(200,data)
}

func getOneContestProblem(c *gin.Context)  {
	team,contest := AuthContest(c)
	if contest==nil {
		return
	}
	label := c.Query("label")
	cproblem := contest.GetOneProblemByLabel(label)
	if cproblem==nil {
		c.String(403,"找不到这道题")
		return
	}
	problem := cproblem.GetProblem()
	problem.ChangeStringToExamples()
	filePath := ""
	if problem.FileName!="" {
		filePath = path.Join(problem.FileDIR,problem.FileName)
	}
	code := ""
	lang := "C++11"
	if team != nil {
		cs := team.GetOneCsubmission(cproblem.ID)
		if cs!=nil {
			code = cs.Code
			lang = cs.Lang
		}
	}
	c.JSON(200,gin.H{
		"info": problem,
		"file_path": filePath,
		"code": code,
		"lang": lang,
		"format": contest.Format,
	})
}

func submitContestCode(c *gin.Context)  {
	team,contest := AuthContestPost(c)
	if contest==nil {
		return
	}
	if contest.Status == "Ended" {
		c.String(403,"比赛已结束，无法提交")
		return
	}

	label := c.PostForm("label")
	cproblem := contest.GetOneProblemByLabel(label)
	if cproblem==nil {
		c.String(403,"找不到这道题")
		return
	}
	if team==nil {
		user := GetCurrentUser(c)
		if contest.Type=="private" || user==nil {
			c.String(401, "未登陆")
			return
		}
		team = &contestApp.Team {
			Name:      user.Username,
			ContestID: contest.ID,
			UserID:    user.ID,
			Order:     contest.Format,
		}
		team.Create()
	}
	lock.Lock()
	cs := contestApp.Csubmission{
		ContestID:   contest.ID,
		RunID: 		 contest.NewCsRunID(),
		TeamID:      team.ID,
		CproblemID:  cproblem.ID,
		Code:        c.PostForm("code"),
		Length:      uint(len(c.PostForm("code"))),
		Lang:        c.PostForm("lang"),
		Author:      team.Name,
		Label: 		 cproblem.Label,
	}
	cs.Create()
	lock.Unlock()
	go contestApp.Judge(&cs,cproblem,team,contest)
	c.JSON(200,gin.H{
		"run_id":cs.RunID,
		"created_at": cs.CreatedAt.Format("2006-01-02 15:04:05"),
		"lang": cs.Lang,
		"status": cs.Status,
		"time": cs.Time,
		"memory": cs.Memory,
		"score" : cs.Score,
	})
}

func searchCsubmissions(c *gin.Context)  {
	_,contest := AuthContestPost(c)
	if contest==nil {
		return
	}
	l,_ := strconv.Atoi(c.DefaultPostForm("l","1"))
	r,_ := strconv.Atoi(c.DefaultPostForm("r","50"))
	var rules map[string]interface{}
	_ = json.Unmarshal([]byte(c.DefaultPostForm("rules","{}")),&rules)
	cs,tot,err:=contestApp.SearchCsubmissions(contest,l,r,rules)
	if err!=nil {
		c.String(403,"查询错误")
		return
	}
	var data []map[string]interface{}
	for _,item:= range cs {
		data=append(data, map[string]interface{}{
			"id": item.RunID,
			"author": item.Author,
			"status": item.Status,
			"label": item.Label,
			"time": item.Time,
			"memory": item.Memory,
			"length": item.Length,
			"lang": item.Lang,
			"score": item.Score,
			"submitTime": item.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	c.JSON(200,gin.H{
		"tot": tot,
		"data": data,
	})
}

func getCsubmission(c *gin.Context)  {
	team,contest := AuthContest(c)
	if contest==nil {
		return
	}
	if team==nil {
		c.String(401,"未登录")
		return
	}
	runID,_ := strconv.Atoi(c.DefaultQuery("run_id","-1"))
	if runID==-1 {
		c.String(403,"错误")
		return
	}
	cs := contest.GetOneCsubmission(uint(runID))
	if cs==nil {
		c.String(403,"找不到提交")
		return
	}
	if team==nil || team.ID != cs.TeamID {
		c.String(403,"无权查看该提交")
		return
	}
	c.JSON(200,gin.H{
		"author": cs.Author,
		"status": cs.Status,
		"index": cs.Label,
		"time": cs.Time,
		"memory": cs.Memory,
		"length": cs.Length,
		"lang": cs.Lang,
		"submitTime": cs.CreatedAt.Format("2006-01-02 15:04:05"),
		"compile_info": cs.CompileInfo,
		"code": cs.Code,
	})
}

func getBaseCS(c *gin.Context)  {
	_,contest := AuthContest(c)
	runID,_ := strconv.Atoi(c.DefaultQuery("run_id","-1"))
	cs := contest.GetOneCsubmission(uint(runID))
	c.JSON(200,gin.H{
		"run_id":cs.RunID,
		"created_at": cs.CreatedAt.Format("2006-01-02 15:04:05"),
		"lang": cs.Lang,
		"status": cs.Status,
		"time": cs.Time,
		"memory": cs.Memory,
		"score" : cs.Score,
	})
}

func getCproblemHistory(c *gin.Context)  {
	team,contest := AuthContest(c)
	if contest==nil {
		return
	}
	if team==nil {
		c.String(401,"未登陆")
		return
	}
	label := c.Query("label")
	cproblem := contest.GetOneProblemByLabel(label)
	if cproblem==nil {
		c.String(403,"找不到这道题")
		return
	}
	cs := team.GetOneProblemCS(cproblem.ID)
	var data []map[string]interface{}
	for _,item := range cs {
		data = append(data, map[string]interface{} {
			"run_id":item.RunID,
			"created_at": item.CreatedAt.Format("2006-01-02 15:04:05"),
			"lang": item.Lang,
			"status": item.Status,
			"time": item.Time,
			"memory": item.Memory,
			"score" : item.Score,
		})
	}
	c.JSON(200,gin.H{
		"num": len(data),
		"data" : data,
	})
}

func getLabelList(c *gin.Context)  {
	_,contest := AuthContest(c)
	if contest==nil {
		return
	}
	cproblems,err := contest.GetCproblems()
	if err!=nil {
		c.String(403,err.Error())
		return
	}
	var lables []string
	for _,item := range *cproblems {
		lables=append(lables,item.Label)
	}
	c.JSON(200,lables)
}

func getRankList(c *gin.Context)  {
	_,contest := AuthContest(c)
	if contest==nil {
		return
	}
	ts,err := contest.GetTeams()
	if err!=nil {
		c.String(403,err.Error())
		return
	}
	firstSolve := make(map[uint]uint)
	cps,_ := contest.GetCproblems()
	for _,item := range *cps {
		firstSolve[item.ID] = item.FirstSolveTime
	}
	var teams contestApp.Teams = *ts
	sort.Sort(teams)
	var data []map[string]interface{}
	for i,t:=range teams {
		ps := t.GetAllProblemStatus()
		mp := make(map[string]interface{})
		mp["rank"] = i+1
		mp["team"] = t.Name
		mp["score"] = t.Scores
		mp["solved"] = t.Solved
		mp["penalty"] = t.Penalty
		for _,item := range ps {
			status := "AC" //AC WA FirstBlood
			if item.Minutes==0 {
				status = "WA"
			} else if item.Minutes <= firstSolve[item.CproblemID] {
				status = "FirstBlood"
			}
			mp[item.Label] = map[string]interface{}{
				"fail_times":item.FailTimes,
				"score": item.Score,
				"minutes": item.Minutes,
				"status": status,
			}
		}
		data = append(data,mp)
	}
	c.JSON(200,data)
}

