package adminApp

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"myGinProject/common"
	"myGinProject/contestApp"
	"myGinProject/problemApp"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)


func autologin(c *gin.Context)  {
	if admin:=GetAdminBySession(c);admin!=nil {
		c.String(200,admin.Name)
		return
	}
	c.String(403, "")
}

func login(c *gin.Context) {
	DeleteSession(c)
	var form Admin
	if err := c.ShouldBind(&form); err != nil {
		c.String(400,err.Error())
		return
	}
	var admin Admin
	if db.Where("name = ?",form.Name).First(&admin).RecordNotFound() {
		c.String(403,"登陆失败，用户名或密码错误")
		return
	}
	if admin.Password != form.Password {
		c.String(403,"登陆失败，用户名或密码错误")
		return
	}
	SetSession(c,admin.Name)
	c.String(200,"登陆成功")
}
func logout(c *gin.Context)  {
	DeleteSession(c)
	c.String(200,"退出成功")
}


//problem, file
func problemHandle(c *gin.Context)  {
	var problem problemApp.Problem
	if err := json.Unmarshal([]byte(c.PostForm("problem")),&problem);err!=nil{
		c.String(403, err.Error())
		return
	}
	if ok, errInfo := common.Validate(problem); !ok {
		c.String(403, errInfo)
		return
	}
	problem.Index = strings.TrimSpace(problem.Index)
	problem.Title = strings.TrimSpace(problem.Title)
	problem.Input = strings.TrimSpace(problem.Input)
	problem.Output = strings.TrimSpace(problem.Output)
	problem.Author = strings.TrimSpace(problem.Author)
	problem.Hint = strings.TrimSpace(problem.Hint)
	var tmp *problemApp.Problem
	if problem.Index!="" {
		tmp=problemApp.GetProblemByIndex(problem.Index)
		problem.ID = tmp.ID
	}
	if problem.Index == "" || (problem.Index[0]=='U' && problem.Type) || (problem.Index[0]=='P' && !problem.Type){
		problem.Index = problemApp.NxtIndex(problem.Type)
	}
	problem.FileDIR = path.Join(problemApp.BASE_DIR,problem.Index)
	if(tmp==nil) {
		_ = os.MkdirAll(problem.FileDIR, os.ModePerm)
	}else if tmp!=nil && problem.FileDIR!=tmp.FileDIR {
		os.Rename(tmp.FileDIR,problem.FileDIR)
	}
	problem.ChangeExamplesToString()

	if file, err := c.FormFile("file");err==nil{
		if tmp!=nil {
			_ = os.Remove(path.Join(tmp.FileDIR, tmp.FileName))
		}

		problem.FileName = file.Filename
		if err2:=c.SaveUploadedFile(file,path.Join(problem.FileDIR,problem.FileName));err2!=nil {
			c.String(500,err2.Error())
			return
		}
	}

	if spj_file,err := c.FormFile("spj_file");err==nil{
		if tmp!=nil {
			_ = os.Remove(tmp.GetSpjFilePath())
		}
		if err2:=c.SaveUploadedFile(spj_file,problem.GetSpjFilePath());err2!=nil {
			c.String(500,err2.Error())
			return
		}
	}

	if tmp==nil {
		if err := problem.Create(); err != nil {
			c.String(403, err.Error())
			return
		}
	} else {
		if err := problem.Update();err!=nil {
			c.String(403,err.Error())
		}
	}
	c.JSON(200,gin.H{
		"index" : problem.Index,
	})
}

func delFile(c *gin.Context)  {
	if problem := problemApp.GetProblemByIndex(c.Query("index"));problem==nil{
		c.String(403,"找不到题目")
		return
	} else {
		if err:=os.Remove(path.Join(problem.FileDIR,problem.FileName));err!=nil {
			c.String(403,err.Error())
			return
		}
		_ = problem.UpdateMap(map[string]interface{}{"file_name":""})
	}
	c.String(200,"删除成功")
}

func getProblems(c *gin.Context)  {
	l,_ := strconv.Atoi(c.Query("l"))
	r,_ := strconv.Atoi(c.Query("r"))
	Type := false
	order := false
	if c.DefaultQuery("type","false")== "true" {
		Type = true
	}
	if c.DefaultQuery("order","false") == "true" {
		order = true
	}
	problems,count,err:= problemApp.GetProblems(l,r,Type,order)
	if err!=nil {
		c.String(403,err.Error())
		return
	}
	var ret []Base
	for _,item := range *problems{
		ret = append(ret, Base{
			Index:  item.Index,
			Title:  item.Title,
			Author: item.Author,
			Source: item.Source,
		})
	}
	c.JSON(200,gin.H{
		"data": ret,
		"total": count,
	})
}

func getOneProblem(c *gin.Context) {
	index := c.Query("index")
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403,"找不到这道题")
		return
	}
	problem.ChangeStringToExamples()
	c.JSON(200,problem)
}

func deleteProblem(c *gin.Context)  {
	index := c.Query("index")
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403,"找不到这道题")
		return
	}
	if err:=problem.Delete();err!=nil{
		c.String(403,err.Error())
		return
	}
	c.String(200,"删除成功")
}


func addTestdatas(c *gin.Context) {
	var zipFile *multipart.FileHeader
	var err error
	if zipFile, err = c.FormFile("zip");err!=nil {
		c.String(403,err.Error())
		return
	}
	index := c.PostForm("index")
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403,"找不到题目")
		return
	}
	if err:=problem.DeleteTestdatas();err!=nil{
		c.String(500,err.Error())
		return
	}
	zipPath := path.Join(problem.FileDIR,zipFile.Filename)
	if err:=c.SaveUploadedFile(zipFile,zipPath);err!=nil {
		c.String(500,err.Error())
		return
	}
	if err:=problem.HandleZipData(zipPath);err!=nil {
		c.String(403,err.Error())
		return
	}
	_ = os.Remove(zipPath)
	c.String(200,"上传成功")
}

func getTestdatas(c *gin.Context) {
	index := c.Query("index")
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403, "找不到这题")
		return
	}
	testdatas,err := problem.GetTestdatas()
	if err!=nil {
		c.String(403,err.Error())
		return
	}
	var info []problemApp.Testdata
	for _,item := range *testdatas {
		info = append(info, item)
	}
	c.JSON(200,gin.H{"data":info})
}

func getOneTestdata(c *gin.Context)  {
	index := c.Query("index")
	case_id,_ := strconv.Atoi(c.Query("case_id"))
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403, "找不到这题")
		return
	}
	testdata,err:=problem.GetOneTestdata(case_id)
	if err!=nil {
		c.String(403,err.Error())
		return
	}
	input,err := common.GetContent(path.Join(problem.FileDIR,testdata.Input))
	if err!=nil {
		c.String(403,err.Error())
	}
	output := ""
	if !problem.Spj {
		output, err = common.GetContent(path.Join(problem.FileDIR, testdata.Output))
		if err != nil {
			c.String(403, err.Error())
		}
	}
	c.JSON(200,gin.H{
		"input": input,
		"output": output,
	})
}

func removeOneTestdata(c *gin.Context)  {
	index := c.Query("index")
	case_id,_ := strconv.Atoi(c.Query("case_id"))
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403, "找不到这题")
		return
	}
	testdata,err:=problem.GetOneTestdata(case_id)
	if err!=nil {
		c.String(403,err.Error())
		return
	}
	if err:=testdata.Delete(problem);err!=nil {
		c.String(403,err.Error())
		return
	}
	if err:=problem.MakeInfoFile();err!=nil {
		c.String(403,err.Error())
		return
	}
	c.String(200,"删除成功")
}

func removeTestdatas(c *gin.Context)  {
	index := c.Query("index")
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(403, "找不到这题")
		return
	}
	if err:=problem.DeleteTestdatas();err!=nil {
		c.String(403,err.Error())
		return
	}
	c.String(200,"删除成功")
}

func searchProbelmTitle(c *gin.Context)  {
	index := c.Query("index")
	problem := problemApp.GetProblemByIndex(index)
	if problem==nil {
		c.String(200, "找不到该题")
		return
	}
	c.String(200,problem.Title)
}

func contestHandle(c *gin.Context)  {
	var contest contestApp.Contest
	if err := json.Unmarshal([]byte(c.PostForm("contest")),&contest);err!=nil{
		c.String(403, err.Error())
		return
	}
	duration := contest.End.Sub(contest.Begin)
	contest.Length = uint(duration/time.Second)
	if contest.ID != 0 {
		if err:=contest.Update();err!=nil {
			c.String(500,err.Error())
			return
		}
	} else {
		if err:=contest.Create();err!=nil {
			c.String(500,err.Error())
			return
		}
	}
	c.JSON(200,gin.H{
		"id": contest.ID,
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
		})
	}
	c.JSON(200,gin.H{
		"tot": tot,
		"data": data,
	})
}

func getContest(c *gin.Context)  {
	id,_ := strconv.Atoi(c.DefaultQuery("id","0"))
	contest := contestApp.GetContestByID(uint(id))
	if contest==nil {
		c.String(403,"不存在该场比赛")
		return
	}
	c.JSON(200,contest)
}

func addCproblems(c *gin.Context)  {
	tmpid,_ := strconv.Atoi(c.PostForm("id"))
	id := uint(tmpid)
	var problems,labels []string
	if err:= json.Unmarshal([]byte(c.PostForm("problems")),&problems);err!=nil{
		c.String(403,err.Error())
		return
	}
	if err:= json.Unmarshal([]byte(c.PostForm("labels")),&labels);err!=nil{
		c.String(403,err.Error())
		return
	}
	contest := contestApp.GetContestByID(id)
	if contest==nil {
		c.String(403,"不存在该比赛")
		return
	}
	if cp,err:= contest.GetCproblems();err!=nil {
		c.String(403,err.Error())
		return
	} else {
		for _,item := range *cp {
			item.Delete()
		}
	}
	for idx,_ := range(problems) {
		problem := problemApp.GetProblemByIndex(problems[idx])
		cproblem := &contestApp.Cproblem{
			ProblemID:problem.ID,
			Title:problem.Title,
			ContestID:id,
			Label:labels[idx],
		}
		cproblem.Create()
	}
	c.String(200,"ok")
}

func getCproblems(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Query("id"))
	contest := contestApp.GetContestByID(uint(id))
	if contest==nil {
		c.String(403,"不存在该比赛")
		return
	}
	cproblems,err := contest.GetCproblems()
	if err!=nil {
		c.String(403,err.Error())
		return
	}
	var problems,titles []string
	for _,cp := range *cproblems {
		problem := cp.GetProblem()
		problems = append(problems,problem.Index)
		titles = append(titles,problem.Title)
	}
	c.JSON(200,gin.H{
		"num": len(problems),
		"problems": problems,
		"titles": titles,
	})
}

func deleteContest(c *gin.Context)  {
	id,_ := strconv.Atoi(c.Query("id"))
	contest := contestApp.GetContestByID(uint(id))
	if contest==nil {
		c.String(403,"不存在该比赛")
		return
	}
	if err:=contest.Delete();err!=nil {
		c.String(403,err.Error())
		return
	}
	c.String(200,"ok")
}

func getContestCode(c *gin.Context)  {
	id,_ := strconv.Atoi(c.DefaultQuery("id","0"))
	run_id,_ := strconv.Atoi(c.DefaultQuery("run_id","0"))
	contest := contestApp.GetContestByID(uint(id))
	if contest==nil {
		c.String(403,"不存在该比赛")
		return
	}
	cs := contest.GetOneCsubmission(uint(run_id))
	if(cs==nil) {
		c.String(403,"找不到提交")
		return
	}
	c.String(200,cs.Code)
}

func searchCsubmissions(c *gin.Context)  {
	id,_ := strconv.Atoi(c.DefaultPostForm("id","0"))
	contest := contestApp.GetContestByID(uint(id))
	if contest==nil {
		c.String(403,"不存在该比赛")
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

func rejudge(c *gin.Context)  {
	labels := c.QueryArray("labels[]")
	id,_ := strconv.Atoi(c.DefaultQuery("id","0"))
	contest := contestApp.GetContestByID(uint(id))
	if contest==nil {
		c.String(403,"不存在该比赛")
		return
	}
	for idx,item := range labels {
		labels[idx] = item[0:strings.Index(item,".")]
	}
	go contestApp.Rejudge(contest,labels)
	c.String(200,"ok")
}