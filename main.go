package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"os"
)

var (
	dingdingToken string	//钉钉告警url
	podId string			//PodId
	folder string			//日期命名文件夹
	postfix string			//文件名的时间后缀
	env string				//部署环境
	projectGroup string		//项目组
	bucketName string		//OSS bucketName
	locaFilename string		//OOM DumpFile
)

func init() {
	flag.StringVar(&podId, "k", "ops", "PodId")
	flag.StringVar(&env, "e", "test", "ENV")
	folder = time.Now().Format("20060102")
	postfix = time.Now().Format("20060102150405")
	locaFilename = "/dumps/oom" //正式
}

// 判断所给路径文件是否存在
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

func main() {
	flag.Parse()
	projectGroup = fmt.Sprintf(strings.Split(podId,"-")[0]) // podId: "ops-demo"
	bucketName = fmt.Sprintf("%s-disaster",projectGroup) //正式的bucketName，如ops-disaster
	// 判断dump文件是否存在
	exist, err := PathExists(locaFilename)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if exist {
		alarm()
		upload()
	}

}
