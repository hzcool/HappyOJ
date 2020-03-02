package common

import (
	"archive/zip"
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	zh_translations "gopkg.in/go-playground/validator.v9/translations/zh"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	DB *gorm.DB
	R *gin.Engine
)

//安装绑定验证
func Validate(s interface{}) (bool,string) {
	Validate := validator.New()
	zh_ch := zh.New()
	uni := ut.New(zh_ch)
	trans, _ := uni.GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(Validate,trans)
	errs := Validate.Struct(s)
	if errs != nil {
		var msg string
		for _,err := range errs.(validator.ValidationErrors) {
			msg += err.Translate(trans) + "\n"
		}
		return false,msg
	}
	return true,""
}

//移走一个文件夹下的所有文件
func RemoveContents(dir string) error {
	d,err := os.Open(dir)
	if err != nil {
		return err
	}
	defer d.Close()
	names,err := d.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _,name := range names {
		err = os.RemoveAll(filepath.Join(dir,name))
		if err != nil {
			return err
		}
	}
	return nil
}

//读取文件内容
func GetContent(path string)  (string,error) {
	file,err := os.Open(path)
	if err!=nil {
		return "",nil
	}
	defer file.Close()
	bt,err := ioutil.ReadAll(file)
	if err !=nil {
		return  "",nil
	}
	return string(bt),err
}

//判断文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//保存zip文件到指定文件夹
func StoreZipFile(zipFIle *zip.File,dest string) error {
	inFile,err := zipFIle.Open()
	if err!=nil {
		return  err
	}
	defer inFile.Close()
	outFile,err2 := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, zipFIle.Mode())
	if err2!=nil {
		return  err2
	}
	defer outFile.Close()
	_,err = io.Copy(outFile,inFile)
	if err!=nil {
		return  err
	}
	return  nil
}

//求文件大小
func FileSize(path string) uint  {
	fileInfo, _ := os.Stat(path)
	return uint(fileInfo.Size())
}

//求文件的md5(每行去掉首尾的空格)
func MD5(path string) (string,error) {
	f,err:=os.Open(path)
	if err != nil {
		return "",err
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	var content string
	for {
		line,err:=rd.ReadString('\n')
		content += strings.TrimSpace(line)
		if err!=nil || io.EOF==err{
			break
		}
	}
	h := md5.New()
	h.Write([]byte(content))
	return hex.EncodeToString(h.Sum(nil)),nil
}