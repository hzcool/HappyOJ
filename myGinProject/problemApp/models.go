package problemApp


import (
	"archive/zip"
	"encoding/json"
	"io/ioutil"
	"myGinProject/common"
	//"myGinProject/userApp"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)


type Problem struct {
	ID uint `json:"-"`
	CreatedAt time.Time

	Author string `gorm:"size:32;INDEX" json:"author" form:"author"` //作者
	Source string `gorm:"size:32" json:"source" form:"source"`  //来源

	Title string `gorm:"size:64;INDEX" json:"title" form:"title"`  //题目
	Background string `gorm:"type:text;" json:"background" form:"background"` //背景说明
	Statement string `gorm:"type:text;" json:"statement" form:"statement"`  //题目描述
	Input string `gorm:"type:text;" json:"input" form:"input"`  //输入说明
	Output string `gorm:"type:text;" json:"output" form:"output"` //输出说明
	ExamplesIn []string `gorm:"-" json:"examples_in" form:"examples_in"` //样例输入
	ExamplesOut []string `gorm:"-" json:"examples_out" form:"examples_out"` //样例输出

	ExamplesInString string `gorm:"type:text" json:"-"`
	ExamplesOutString string `gorm:"type:text" json:"-"`
	Hint string `gorm:"type:text;" json:"hint" form:"hint"`  //提示

	TimeLimit uint `gorm:"required" json:"time_limit" form:"time_limit" validate:"numeric,min=100,max=40000,required" ` //时间限制
	MemoryLimit uint `gorm:"required" json:"memory_limit" form:"memory_limit" validate:"numeric,min=32,max=2048,required" ` //空间限制

	Type bool `gorm:"INDEX" json:"type"` //false私有 true公开
	Index string `gorm:"size:10;UNIQUE;unique_index" json:"index" form:"index"`

	FileDIR string `json:"-"`
	FileName string `json:"file_name"`
	Spj bool `gorm:"DEFAULT:false" json:"spj"`
	SpjType string `gorm:"size:10" json:"spj_type"`

	AC uint
	All uint
}

//样例用这个隔开，题目样例不能出现这个
const JOINString = "1~#^* ^@$l"

func (p *Problem) ChangeExamplesToString()  {
	p.ExamplesInString = strings.Join(p.ExamplesIn,JOINString)
	p.ExamplesOutString = strings.Join(p.ExamplesOut,JOINString)
}

func (p *Problem) ChangeStringToExamples()  {
	p.ExamplesIn = strings.Split(p.ExamplesInString,JOINString)
	p.ExamplesOut = strings.Split(p.ExamplesOutString,JOINString)
}

func (p *Problem) GetSpjFilePath() string {
	suf := ".c"
	if p.SpjType=="C++" {
		suf = ".cpp"
	} else if p.SpjType=="Python2" || p.SpjType=="Python3" {
		suf = ".py"
	}
	return path.Join(p.FileDIR,"spj"+suf)
}

func (p *Problem) Create() error {
	return db.Create(p).Error
}
func (p *Problem) Save() error{
	return db.Save(p).Error
}
func (p *Problem) Update() error{
	if p.Type == false {
		db.Model(p).Update("type", false)
	}
	if p.Spj == false {
		db.Model(p).Update("spj", false)
	}
	return db.Model(p).Updates(p).Error
}
func (p *Problem) UpdateMap(mp map[string]interface{}) error{
	return db.Model(p).Updates(mp).Error
}

func (p *Problem) Delete() error{
	if err:=os.RemoveAll(p.FileDIR);err!=nil{
		return err
	}
	return db.Delete(p).Error
}

func (p *Problem)GetTestdatas() (*[]Testdata,error){
	var testdatas []Testdata
	if err:=db.Model(p).Order("case_id").Related(&testdatas).Error;err!=nil{
		return nil,err
	}
	return &testdatas,nil
}

func (p *Problem)GetOneTestdata(case_id int) (*Testdata,error) {
	var testdata Testdata
	if err:=db.Model(p).Where("case_id = ?",case_id).Related(&testdata).Error;err!=nil{
		return nil,err
	}
	return  &testdata,nil
}

func (p *Problem)DeleteTestdatas() error {
	testdatas,err := p.GetTestdatas()
	if err!=nil {
		return  err
	}
	for _,item := range *testdatas {
		if err:=item.Delete(p);err!=nil{
			return  err
		}
	}
	infoPath := path.Join(p.FileDIR,"info")
	if exist,_:=common.PathExists(infoPath);exist {
		if err:=os.Remove(infoPath);err!=nil {
			return  err
		}
	}
	return nil
}

func (p *Problem)HandleZipData(zipPath string) error {  //导入的zip文件进行处理
	zipReader,err := zip.OpenReader(zipPath)
	if err !=nil {
		return err
	}
	defer zipReader.Close()
	input := make(map[int]*zip.File)
	output := make(map[int]*zip.File)
	for _,f := range zipReader.File {
		var id,base = 0,1
		for i:=strings.LastIndex(f.Name,".")-1;i>=0;i-- {
			ch := int(f.Name[i]) - int('0')
			if ch<0 || ch>9 {
				break
			}
			id += base * ch
			base *= 10
		}
		if path.Ext(f.Name) == ".in" {
			input[id] = f
		} else if path.Ext(f.Name)==".out" || path.Ext(f.Name)==".ans" {
			output[id] = f
		}
	}
	if p.Spj {
		for k, in := range input {
			inPath := path.Join(p.FileDIR,in.Name)
			if err := common.StoreZipFile(in,inPath);err!=nil {
				return err
			}
			testdata := Testdata{
				CaseID:            k,
				Input:             in.Name,
				InputSize:         common.FileSize(inPath),
				ProblemID:         p.ID,
			}
			if err := testdata.Create();err!=nil {
				return err
			}
		}
	} else {
		for k,in :=  range input {
			if out,ok := output[k];ok {
				inPath := path.Join(p.FileDIR,in.Name)
				outPath := path.Join(p.FileDIR,out.Name)
				if err := common.StoreZipFile(in,inPath);err!=nil {
					return err
				}
				if err := common.StoreZipFile(out,outPath);err!=nil {
					return err
				}
				md5,err := common.MD5(outPath)
				if err!=nil {
					return err
				}
				testdata := Testdata{
					CaseID:            k,
					Input:             in.Name,
					Output:            out.Name,
					InputSize:         common.FileSize(inPath),
					OutputSize:        common.FileSize(outPath),
					StrippedOutputMd5: md5,
					ProblemID:         p.ID,
				}
				if err := testdata.Create();err!=nil {
					return err
				}
			}
		}
	}
	return  p.MakeInfoFile()
}

func (p *Problem)MakeInfoFile() error {  //生成info文件
	testdatas,err := p.GetTestdatas()
	if err!=nil {
		return  err
	}
	info,cases := make(map[string]interface{}),make(map[string]interface{})
	info["spj"] = p.Spj
	if p.Spj {
		info["spj_lang"] = p.SpjType
	}
	for _,item := range *testdatas {
		tmp := map[string]interface{}{
			"input_name" : item.Input,
		}
		if !p.Spj {
			tmp["output_name"] = item.Output
			tmp["output_size"] = item.OutputSize
			tmp["stripped_output_md5"] = item.StrippedOutputMd5
		}
		cases[strconv.Itoa(item.CaseID)] = tmp
	}
	info["test_cases"] = cases
	js,_ := json.Marshal(info)
	if err:= ioutil.WriteFile(path.Join(p.FileDIR,"info"),js,os.ModePerm); err!=nil{
		return  err
	}
	return nil
}

func (p *Problem)GetSubmissions() (*[]Submission,error){
	var ss []Submission
	if err := db.Model(p).Order("id desc").Related(&ss).Error;err!=nil{
		return  nil,err
	}
	return &ss,nil
}
func (p *Problem)GetRatio() float32 {
	if p.All == 0 {
		return  0
	}
	return float32(p.AC)/float32(p.All)*100
}


type Testdata struct {
	ID uint `json:"-"`
	CaseID int `gorm:"size:10" json:"case_id"`  //1
	Input string `gorm:"size:32" json:"input_name"`  //输入文件名 1.in
	Output string `gorm:"size:32" json:"output_name"` //输出文件名 1.out
	InputSize uint `json:"input_size" gorm:"column:input_size"`
	OutputSize uint `json:"output_size" gorm:"column:output_size"`
	StrippedOutputMd5 string `gorm:"size:32;column:stripped_output_md5" json:"stripped_output_md5"`
	ProblemID uint  `gorm:"INDEX" json:"-"`
}

func (t *Testdata)GetProblem() (*Problem,error) {
	var problem Problem
	if err:=db.Model(t).Related(&problem).Error;err!=nil{
		return nil,err
	}
	return &problem,nil
}

func (t *Testdata) Create() error {
	return  db.Create(t).Error
}

func (t *Testdata) Delete(p *Problem) error {
	if err:=os.Remove(path.Join(p.FileDIR,t.Input));err!=nil {
		return  err
	}
	if p.Spj==false || t.Output!="" {
		if err := os.Remove(path.Join(p.FileDIR, t.Output)); err != nil {
			return err
		}
	}
	return db.Delete(t).Error
}

type Submission struct {
	ID uint `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Code string `gorm:"type:text" json:"-"`
	Time uint `json:"time"`
	Memory uint `json:"memory"`
	Length uint `json:"length"`
	Status string `gorm:"size:20;default:'Queueing';INDEX" json:"status"`
	CompileInfo string `gorm:"type:text" json:"info"`
	Lang string `gorm:"size:10;INDEX" json:"lang"`
	ProblemID uint `gorm:"INDEX" json:"-"`
	UserID uint `gorm:"INDEX" json:"-"`
	Score uint

	ProblemIndex string `gorm:"size:10;INDEX"`
	Author string
}


func (s *Submission)Create() error {
	return db.Create(s).Error
}
func (s *Submission) UpdateMap(mp map[string]interface{}) error{
	return db.Model(s).Updates(mp).Error
}
func (s *Submission)GetProblem() *Problem {
	var problem Problem
	_ = db.Model(s).Related(&problem)
	return &problem
}


func models_init()  {
	db.AutoMigrate(&Problem{})
	db.AutoMigrate(&Testdata{})
	db.AutoMigrate(&Submission{})
	db.Model(&Testdata{}).AddForeignKey("problem_id","problem(id)","CASCADE", "CASCADE")
	db.Model(&Submission{}).AddForeignKey("problem_id","problem(id)","CASCADE", "CASCADE")
	db.Model(&Submission{}).AddForeignKey("user_id","user(id)","CASCADE", "CASCADE")
}



