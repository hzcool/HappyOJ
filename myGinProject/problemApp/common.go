package problemApp

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type ErrType int
const (
	OK ErrType = 0   //没有错误
	INVALID_HEADER ErrType = 1 // 请求头错误
	INVALID_JSON_REQUEST ErrType = 2 // 请求的json数据不对
	STORE_SRC_FAILED ErrType = 3 // 系统保存上传的测试代码出错
	COMPILE_ERROR ErrType = 4  // 编译错误
	INVALID_CONFIG ErrType = 5 // 时间、内存的设置不在合理范围内
	INVAILD_TESTCASE_INFO_FILE ErrType = 6 //测试数据的info描述文件不对
	SYSTEM_EXCEPTION ErrType = 7 //系统异常
)

const (
	ACCEPTED int = 0
	WRONG_ANSWER int = 1
	TIME_LIMIT_EXCEEDED int = 2
	MEMORY_LIMIT_EXCEEDED int = 3
	OUTPUT_LIMIT_EXCEEDED int = 4
	RUNTIME_ERROR int = 5
	SYSTEM_ERROR int = 6
)
var resultMap = map[int]string {
	0: "Accepted",
	1: "WrongAnswer",
	2: "TimeLimitExceeded",
	3: "MemoryLimitExceeded",
	4: "OutputLimitExceeded",
	5: "RuntimeError",
	6: "SystemError",
}



var (
	db *gorm.DB
	BASE_DIR = filepath.Join(os.Getenv("BASE_PATH"),"test_case")
	MAX_TASKS,_ = strconv.Atoi(os.Getenv("MAX_TASKS"))
	SERVICE_URL = os.Getenv("SERVICE_URL")
	ACCESS_TOKEN = os.Getenv("ACCESS_TOKEN")
	client = &http.Client{}
	CH chan int	//控制最大数量的同时处理
)

func Init(r *gin.Engine,_db *gorm.DB)  {
	db = _db
	CH = make(chan int,MAX_TASKS)
	for i:=1;i<=MAX_TASKS;i++ {
		CH <- i
	}
	urls_init(r)
	models_init()
}


func NxtIndex(Type bool) string {
	var problem Problem
	index := "1000"
	if !db.Where("Type = ?",Type).Order("problem.index desc").First(&problem).RecordNotFound() {
		x,_ := strconv.Atoi(problem.Index[1:])
		index = strconv.Itoa(x+1)
	}
	if(Type) {
		return "P" + index
	}
	return  "U" + index
}

func GetProblemByIndex(index string) *Problem {
	var problem Problem
	if	db.Where(map[string]interface{}{"index":index}).First(&problem).RecordNotFound(){
		return  nil
	}
	return &problem
}

//order true为从小到大,否则从大到小
func GetProblems(l,r int,Type bool,order bool) (*[]Problem,int,error) {
	var problems []Problem
	orderStr := "problem.index"
	if !order {
		orderStr += " desc"
	}
	err := db.Where("type = ?",Type).Order(orderStr).Offset(l-1).Limit(r-l+1).Find(&problems).Error
	if err!=nil {
		return nil,0,err
	}
	cnt := 0
	err = db.Model(&Problem{}).Where("type = ?",Type).Count(&cnt).Error
	return &problems,cnt,err
}
func SearchProblems(l,r int,rules map[string]interface{}) ([]Problem,int,error) {
	var problems []Problem
	err := db.Where(rules).Order("problem.index").Find(&problems).Error
	if err!=nil {
		return []Problem{},0,err
	}
	cnt := len(problems)
	if l>cnt {
		return []Problem{},cnt,nil
	}
	if r > cnt {
		r = cnt
	}
	return problems[l-1:r],cnt,nil
}

func GetSubmissionById(id uint) *Submission {
	var s Submission
	if db.First(&s,id).RecordNotFound() {
		return nil
	}
	return &s
}

func SearchSubmissions(l,r int,rules map[string]interface{}) ([]Submission,int,error) {
	var ss []Submission
	err := db.Where(rules).Order("id desc").Find(&ss).Error
	if err!=nil {
		return []Submission{},0,err
	}
	cnt := len(ss)
	if l>cnt {
		return []Submission{},cnt,nil
	}
	if r > cnt {
		r = cnt
	}
	return ss[l-1:r],cnt,nil
}

