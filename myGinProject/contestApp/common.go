package contestApp

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"myGinProject/problemApp"
	"time"
)

var (
	db *gorm.DB
	pending,running map[uint]*time.Timer
)
func Init(r *gin.Engine,_db *gorm.DB)  {
	db = _db
	models_init()
	urls_init(r)
	pending,running = make(map[uint]*time.Timer),make(map[uint]*time.Timer)
	var contests []Contest
	db.Where("status = 'Pending'").Or("status = 'Running'").Find(&contests)
	for idx,_ := range contests {
		PutContest(&contests[idx])
	}
}
func RemoveContest(c *Contest)  {
	if _,ok := pending[c.ID]; ok {
		pending[c.ID].Stop()
		delete(pending,c.ID)
	}
	if _,ok := running[c.ID]; ok {
		running[c.ID].Stop()
		delete(running,c.ID)
	}
}
func PutContest(c *Contest)  {
	RemoveContest(c)
	now := time.Now()
	if now.After(c.End) {
		if c.Status!="Ended" {
			c.UpdateMap(map[string]interface{}{"status": "Ended"})
		}
		return
	}
	if now.Before(c.Begin) {
		if c.Status!="Pending" {
			c.UpdateMap(map[string]interface{}{"status": "Pending"})
		}
		pending[c.ID] = time.AfterFunc(c.Begin.Sub(now), func() {
			fmt.Println("比赛",c.ID,"已经开始")
			PutContest(c)
		})
		return
	}
	if c.Status!="Running" {
		c.UpdateMap(map[string]interface{}{"status": "Running"})
	}
	running[c.ID] = time.AfterFunc(c.End.Sub(c.Begin), func() {
		fmt.Println("比赛",c.ID,"已经结束")
		PutContest(c)
	})
}

func GetContestByID(id uint) *Contest {
	var c Contest
	db.First(&c,id)
	return &c
}

func SearchContests(l,r int,rules map[string]interface{}) ([]Contest,int,error) {
	var cs []Contest
	err := db.Where(rules).Order("id desc").Find(&cs).Error
	if err!=nil {
		return  []Contest{},0,err
	}
	cnt := len(cs)
	if l>cnt {
		return []Contest{},cnt,nil
	}
	if r>cnt {
		r = cnt
	}
	return  cs[l-1:r],cnt,nil
}

func SearchCsubmissions(contest *Contest,l,r int,rules map[string]interface{}) ([]Csubmission,int,error) {
	var cs []Csubmission
	err := db.Model(contest).Where(rules).Order("id desc").Related(&cs).Error
	if err!=nil {
		return  []Csubmission{},0,err
	}
	cnt := len(cs)
	if l>cnt {
		return []Csubmission{},cnt,nil
	}
	if r>cnt {
		r = cnt
	}
	return  cs[l-1:r],cnt,nil
}

func Judge(cs *Csubmission,cproblem *Cproblem,team *Team,contest *Contest)  {
	chID := <- problemApp.CH
	defer func() {
		problemApp.CH <- chID
	}()
	cs.UpdateMap(map[string]interface{}{"status":"Running"})
	problem := cproblem.GetProblem()
	update := problemApp.ToJudge(map[string]interface{}{
		"lang": cs.Lang,
		"max_cpu_time": problem.TimeLimit,
		"max_memory": problem.MemoryLimit*1024*1024,
		"test_case": problem.Index,
		"src": cs.Code,
		"run_id": cs.RunID,
	})
	cs.UpdateMap(update)
	ps := team.GetOneProblemStatus(cproblem)
	if ps==nil {
		ps = &ProblemStatus{
			TeamID:     team.ID,
			CproblemID: cproblem.ID,
			Label: cproblem.Label,
		}
		ps.Create()
	}
	if(ps.Minutes==0) {
		if(cs.Status=="Accepted") {
			team.Scores += 100 - ps.Score
			team.Solved ++
			ps.Score = 100
			ps.Minutes = uint(cs.CreatedAt.Sub(contest.Begin).Minutes()) + 1
			team.Penalty += ps.FailTimes*20 + ps.Minutes
			cproblem.AC++
			cproblem.All++
			if cproblem.FirstSolveTime == 0 || cproblem.FirstSolveTime > ps.Minutes{
				cproblem.FirstSolveTime = ps.Minutes
			}
		} else {
			cproblem.All++
			ps.FailTimes++
			if(ps.Score < cs.Score) {
				team.Scores += cs.Score - ps.Score
				ps.Score = cs.Score
			}
		}
		team.Update()
		cproblem.Update()
		ps.Update()
	} else {
		if cs.Status=="Accepted" {
			cproblem.AC++
		}
		cproblem.All++
		cproblem.Update()
	}
}

func Rejudge(contest *Contest,labels []string)  {
	var cproblems []Cproblem
	if err:=db.Model(contest).Where("label in (?)",labels).Related(&cproblems).Error;err!=nil{
		return
	}
	var teams []Team
	if err:=db.Model(contest).Related(&teams).Error;err!=nil {
		return
	}

	for _,t := range teams {
		for _,cp := range cproblems {
			if ps := t.GetOneProblemStatus(&cp);ps != nil {
				t.Scores -= ps.Score
				if ps.Minutes>0 {
					t.Solved--
					t.Penalty -= ps.Minutes
				}
				ps.Delete()
			}
		}
		t.Update()
	}
	for i,_ := range cproblems {
		cp := &cproblems[i]
		cp.AC,cp.All,cp.FirstSolveTime = 0,0,0
		css := contest.GetOneProblemCS(cp.ProblemID)
		for idx,_ := range *css {
			cs := &(*css)[idx]
			cs.UpdateMap(map[string]interface{}{"status":"Queueing"})
			go Judge(cs,cp,cs.GetTeam(),contest)
		}
	}
}