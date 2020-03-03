package userApp

import (
	"myGinProject/problemApp"
	"sort"
	"time"
)


type User struct {
	ID       uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	LastSeen time.Time `json:"last_seen"`
	Username string `gorm:"size:32;UNIQUE;unique_index;not null" json:"username" `
	Password string `gorm:"size:16;not null" json:"-"`
	Email    string	`gorm:"UNIQUE;not null;" json:"email"`
	School 	 string `gorm:"not null;size:20;INDEX" json:"school"`
	Description string `gorm:"DEFAULT:''" json:"description"`
}


//按照规则查询关于自己的提交
func (u *User) GetSubmissions(rules map[string]interface{}) (*[]problemApp.Submission,error) {
	var ss []problemApp.Submission
	if err := db.Model(u).Where(rules).Order("id desc").Related(&ss).Error;err!=nil{
		return nil,err
	}
	return &ss,nil
}


func (u *User) GetOneSubmission(problem *problemApp.Problem) *problemApp.Submission {
	ss,_ := u.GetSubmissions(map[string]interface{}{"problem_id":problem.ID})
	if ss==nil || len(*ss)==0{
		return nil
	}
	return &(*ss)[0]
}

func (u *User) GetProblemState(problem *problemApp.Problem) int{ //查看是否AC,0未尝试,1AC,2错误
	ss,_ := u.GetSubmissions(map[string]interface{}{"problem_id":problem.ID})
	if ss==nil || len(*ss)==0{
		return 0
	}
	for _,item := range *ss {
		if item.Status == "Accepted" {
			return 1
		}
	}
	return 2
}
func (u *User)GetACAndAll() (int,int) {
	var ss []problemApp.Submission
	db.Model(u).Related(&ss)
	mp := make(map[string]int)
	for _,item := range ss {
		p := item.GetProblem()
		if item.Status=="Accepted" {
			mp[p.Index] += 10000
		}  else {
			mp[p.Index] --
		}
	}
	solved,all :=0,len(ss)
	for _,v := range mp {
		if(v>0) {
			solved++
		}
	}
	return  solved,all
}
func (u *User)GetStates() ([]string,[]string,map[string]uint) {
	var ss []problemApp.Submission
	db.Model(u).Related(&ss)
	status := map[string]uint {
		"Accepted": 0,
		"WrongAnswer": 0,
		"TimeLimitExceeded": 0,
		"MemoryLimitExceeded": 0,
		"OutputLimitExceeded": 0,
		"RuntimeError": 0,
		"SystemError": 0,
		"CompileError": 0,
	}
	var solved,unsolved []string
	mp := make(map[string]int)
	for _,item := range ss {
		if(item.Status!="Running" && item.Status!="Queueing") {
			status[item.Status]++
		}
		p := item.GetProblem()
		if item.Status=="Accepted" {
			mp[p.Index] += 10000
		}  else {
			mp[p.Index] --
		}
	}
	for k,v := range mp {
		if v<0 {
			unsolved = append(unsolved, k)
		} else {
			solved = append(solved,k)
		}
	}
	sort.Strings(solved)
	sort.Strings(unsolved)
	return solved,unsolved,status
}

type UsersRank struct {
	Username string `json:"username"`
	School string `json:"school"`
	Solved uint `json:"solved"`
	All uint `json:"all"`
	Ratio float32 `json:"ratio"`
}
type UserRanks []UsersRank
func (u UserRanks) Len() int { return len(u) }
func (u UserRanks) Swap(i, j int) { u[i], u[j] = u[j], u[i] }
func (u UserRanks) Less(i, j int)  bool {
	if u[i].Solved == u[j].Solved {
		if u[i].All == u[j].All {
			return u[i].Username < u[j].Username
		}
		return u[i].All < u[j].All
	}
	return u[i].Solved > u[j].Solved
}
func models_init() {
	db.AutoMigrate(&User{})
}








