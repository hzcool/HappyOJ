package problemApp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Ping()  {
	url := SERVICE_URL+"/ping"
	req,_ := http.NewRequest("POST",url,nil)

	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Access-Token",os.Getenv("ACCESS_TOKEN"))

	res,_ := client.Do(req)
	defer res.Body.Close()

	var mp  map[string]interface{}
	json.NewDecoder(res.Body).Decode(&mp)
	fmt.Println("ping result : ")
	fmt.Println(mp["info"].(string))
}

func ToJudge(task map[string]interface{}) map[string]interface{} {
	post,_ := json.Marshal(task)
	url := SERVICE_URL+"/judge"
	req,_ := http.NewRequest("POST",url,bytes.NewBuffer(post))
	req.Header.Add("Content-Type","application/json")
	req.Header.Add("Access-Token",ACCESS_TOKEN)
	res,_ := client.Do(req)
	defer res.Body.Close()

	//结果处理
	resBody, _ := ioutil.ReadAll(res.Body)
	var result map[string]interface{}
	_ = json.Unmarshal(resBody,&result)
	var update map[string]interface{}
	err := ErrType(result["err"].(float64))
	if err == OK {
		update = map[string]interface{}{
			"status": resultMap[int(result["result"].(float64))],
			"time": uint(result["cpu_time"].(float64)),
			"memory": uint(result["memory"].(float64))/1024/1024,
			"score": uint(result["pass"].(float64) / result["total"].(float64)*100),
			"compile_info": result["compile_info"].(string),
		}
	} else if err == COMPILE_ERROR{
		update = map[string]interface{}{
			"status": "CompileError",
			"compile_info": result["info"].(string),
		}
	} else {
		update = map[string]interface{} {
			"status": "SystemError",
		}
	}
	return  update
}

func Judge(s *Submission,problem *Problem)  {
	chID := <-CH
	defer func() {
		CH <- chID
	}()
	//测评
	s.UpdateMap(map[string]interface{}{"status":"Running"})
	update := ToJudge(map[string]interface{}{
		"lang": s.Lang,
		"max_cpu_time": problem.TimeLimit,
		"max_memory": problem.MemoryLimit*1024*1024,
		"test_case": problem.Index,
		"src": s.Code,
	})
	s.UpdateMap(update)
	if s.Status=="Accepted" {
		problem.UpdateMap(map[string]interface{}{"ac":problem.AC+1,"all":problem.All+1})
	} else {
		problem.UpdateMap(map[string]interface{}{"all":problem.All+1})
	}
}




