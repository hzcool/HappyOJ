package contestApp

import (
	"myGinProject/problemApp"
	"time"
)


type Contest struct {
	ID uint `json:"id"`
	Title string `gorm:"INDEX;size:32" json:"title"`
	Begin time.Time `json:"begin"`
	End time.Time `json:"end"`
	Length uint   `json:"length"`//单位秒 
	Desc string `json:"desc"`
	Author string `json:"author"`
	Type string  `gorm:"INDEX;size:10" json:"type"`//public或private
	Password string `json:"password"`
	Format string `json:"format"`//ACM 或 OI
	Status string `gorm:"DEFAULT:'Pending';INDEX" json:"status"` //Pending,Running,Ended
}

func (c *Contest)Create() error {
	err := db.Create(c).Error
	if err!=nil {
		return  err
	}
	PutContest(c)
	return nil
}
func (c *Contest)Update() error {
	err := db.Model(c).Updates(c).Error
	if err!=nil {
		return  err
	}
	PutContest(c)
	return nil
}
func (c *Contest)UpdateMap(mp map[string]interface{}) error {
	return db.Model(c).Updates(mp).Error
}
func (c *Contest)GetCproblems() (*[]Cproblem,error){
	var cproblems []Cproblem
	if err:=db.Model(c).Related(&cproblems).Error;err!=nil {
		return nil,err
	}
	return &cproblems,nil
}
func (c *Contest)GetTeams() (*[]Team,error){
	var teams []Team
	if err:=db.Model(c).Related(&teams).Error;err!=nil {
		return nil,err
	}
	return &teams,nil
}
func (c *Contest)GetTeamsCount() int {
	var teams []Team
	db.Model(c).Related(&teams)
	return len(teams)
}
func (c *Contest)Delete() error {
	RemoveContest(c)
	return db.Delete(c).Error
}
func (c *Contest)GetOneTeamByUserID(userId uint) *Team  {
	var team Team
	if db.Model(c).Where("user_id = ?",userId).Related(&team).RecordNotFound() {
		return nil
	}
	return &team
}
func (c *Contest)GetOneProblemByLabel(label string) *Cproblem  {
	var cp Cproblem
	if db.Model(c).Where("label = ?",label).Related(&cp).RecordNotFound() {
		return nil
	}
	return &cp
}

func (c *Contest)CountCsubmissions() uint {
	var cnt uint
	db.Model(&Csubmission{}).Where("contest_id = ?",c.ID).Count(&cnt)
	return cnt
}
func (c *Contest)NewCsRunID() uint {
	var cs Csubmission
	if db.Model(&Csubmission{}).Where("contest_id = ?",c.ID).Last(&cs).RecordNotFound() {
		return 1
	}
	return cs.RunID + 1
}
func (c *Contest)GetOneCsubmission(runID uint) *Csubmission {
	var cs Csubmission
	if db.Model(&Csubmission{}).Where("contest_id = ?",c.ID).Where("run_id = ?",runID).First(&cs).RecordNotFound() {
		return  nil
	}
	return  &cs
}
func  (c *Contest)GetOneProblemCS(cproblem_id uint) *[]Csubmission {
	var cs []Csubmission
	if err:=db.Model(c).Where("cproblem_id = ?",cproblem_id).Related(&cs).Error;err!=nil{
		return  nil
	}
	return &cs
}

type Team struct {
	ID uint
	Name string `gorm:"size:32;UNIQUE_INDEX"`
	ContestID uint `gorm:"INDEX"`
	UserID uint    `gorm:"INDEX"`
	Solved uint
	Penalty uint
	Scores uint
	Order string `gorm:"size:5"` //ACM或OI
}

func (t *Team)Create() error {
	return db.Create(t).Error
}
func (t *Team)Update() error {
	return  db.Model(t).Updates(t).Error
}
func (t *Team)GetAllProblemStatus() []ProblemStatus {
	var ps []ProblemStatus
	db.Model(t).Related(&ps)
	return ps
}
func (t *Team)GetOneProblemStatus(cp *Cproblem) *ProblemStatus {
	var ps ProblemStatus
	if db.Where("team_id = ?",t.ID).Where("cproblem_id = ?",cp.ID).First(&ps).RecordNotFound() {
		return nil
	}
	return &ps
}
func (t *Team)GetOneCsubmission(cproblemID uint) *Csubmission {
	var cs Csubmission
	if err:=db.Where("team_id = ?",t.ID).Where("cproblem_id = ?",cproblemID).Last(&cs).Error;err!=nil {
		return  nil
	}
	return &cs
}
func (t *Team)GetOneProblemCS(cproblemID uint) []Csubmission {
	var cs []Csubmission
	db.Model(t).Where("cproblem_id = ?",cproblemID).Order("id desc").Related(&cs)
	return cs
}

type Teams  []Team
func (t Teams) Len() int { return len(t) }
// 根据元素的年龄降序排序 （此处按照自己的业务逻辑写）
func (t Teams) Less(i, j int) bool {
	if t[i].Order!="ACM" {
		if(t[i].Scores == t[j].Scores) {
			return  t[i].ID < t[j].ID
		}
		return t[i].Scores > t[j].Scores
	}
	if t[i].Solved == t[j].Solved {
		if t[i].Penalty == t[j].Penalty {
			return t[i].ID < t[j].ID
		}
		return t[i].Penalty < t[j].Penalty
	}
	return t[i].Solved > t[j].Solved
}
func (t Teams) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

type Cproblem struct {
	ID uint
	ProblemID uint `gorm:"INDEX"`
	ContestID uint `gorm:"INDEX"`
	Label string `gorm:"size:10;INDEX"`
	FirstSolveTime uint
	Title string `gorm:"INDEX"`
	AC uint
	All uint
}
func (cp *Cproblem)Create() error {
	return db.Create(cp).Error
}
func (cp *Cproblem)Update() error {
	return  db.Model(cp).Updates(cp).Error
}
func (cp *Cproblem)UpdateMap(mp map[string]interface{}) error {
	return db.Model(cp).Updates(mp).Error
}
func (cp *Cproblem)Delete() error {
	return db.Delete(cp).Error
}

func (cp *Cproblem)GetProblem() *problemApp.Problem {
	var p problemApp.Problem
	db.Model(cp).Related(&p)
	return &p
}


type Csubmission struct {
	ID uint `json:"id"`
	RunID uint `gorm:"INDEX" json:"run_id"`
	ContestID uint `gorm:"INDEX" json:"-"`
	TeamID uint `gorm:"INDEX" json:"-"`
	CproblemID uint `gorm:"size:10;INDEX" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Code string `gorm:"type:text" json:"-"`
	Time uint `json:"time"`
	Memory uint `json:"memory"`
	Length uint `json:"length"`
	Status string `gorm:"size:20;INDEX;default:'Queueing'" json:"status"`
	CompileInfo string `gorm:"type:text" json:"info"`
	Lang string `gorm:"size:10;INDEX" json:"lang"`
	Score uint `json:"score"`

	Label string `gorm:"size:10;INDEX"`
	Author string
}
func (cs *Csubmission)Create() error {
	return db.Create(cs).Error
}
func (cs *Csubmission)Update() error {
	return  db.Model(cs).Updates(cs).Error
}
func (cs *Csubmission)UpdateMap(mp map[string]interface{}) error   {
	return db.Model(cs).Updates(mp).Error
}
func (cs *Csubmission)GetCproblem() *Cproblem {
	var cproblem Cproblem
	if err:=db.Model(cs).Related(&cproblem).Error;err!=nil{
		return nil
	}
	return &cproblem
}
func (cs *Csubmission)GetTeam() *Team {
	var t Team
	if err:=db.Model(cs).Related(&t).Error;err!=nil{
		return nil
	}
	return &t
}


type ProblemStatus struct {
	ID uint
	CproblemID uint `gorm:"INDEX"`
	Label string `gorm:"INDEX"`
	TeamID uint `gorm:"INDEX"`
	FailTimes uint
	Minutes uint
	Score uint
}

func (ps *ProblemStatus)Create() error {
	return db.Create(ps).Error
}
func (ps *ProblemStatus)Update() error {
	return db.Model(ps).Updates(ps).Error
}
func (ps *ProblemStatus)Delete() error {
	return db.Delete(ps).Error
}


func models_init()  {
	db.AutoMigrate(&Contest{})
	db.AutoMigrate(&Team{})
	db.AutoMigrate(&Cproblem{})
	db.AutoMigrate(&Csubmission{})
	db.AutoMigrate(&ProblemStatus{})
	db.Model(&Team{}).AddForeignKey("contest_id","contest(id)","CASCADE", "CASCADE")
	db.Model(&Team{}).AddForeignKey("user_id","user(id)","CASCADE", "CASCADE")
	db.Model(&Cproblem{}).AddForeignKey("problem_id","problem(id)","CASCADE", "CASCADE")
	db.Model(&Cproblem{}).AddForeignKey("contest_id","contest(id)","CASCADE", "CASCADE")
	db.Model(&Csubmission{}).AddForeignKey("contest_id","contest(id)","CASCADE", "CASCADE")
	db.Model(&Csubmission{}).AddForeignKey("team_id","team(id)","CASCADE", "CASCADE")
	db.Model(&Csubmission{}).AddForeignKey("cproblem_id","cproblem(id)","CASCADE", "CASCADE")
	db.Model(&ProblemStatus{}).AddForeignKey("cproblem_id","cproblem(id)","CASCADE", "CASCADE")
	db.Model(&ProblemStatus{}).AddForeignKey("team_id","team(id)","CASCADE", "CASCADE")
}













