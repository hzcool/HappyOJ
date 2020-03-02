package userApp

import (
	"myGinProject/problemApp"
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

func models_init() {
	db.AutoMigrate(&User{})
}








