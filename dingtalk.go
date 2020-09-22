package main

import (
	"fmt"
	"github.com/braumye/grobot"
	"time"
)


func msgSender(dingdingToken string,msg string) {
	robot, _ := grobot.New("dingtalk", dingdingToken)
	err := robot.SendMarkdownMessage("通知",msg)
	fmt.Println("通知发送完毕",err)
}

func alarm() {
	//key=项目组，value=钉钉token [需要修改]
	tokenMap := map[string]string{
		"ops":"xxx",
	}
	dingdingToken = tokenMap[projectGroup]
	ossurl := fmt.Sprintf("http://%s.oss-cn-shanghai-internal.aliyuncs.com/%s/k8s/jvm/%s/%s-%s", bucketName,env,folder,podId,postfix) //建议OSS内网，"oss-cn-shanghai-internal"更换成自己的endpoint[需要修改]
	alarmMsg := fmt.Sprintf("<font color=#FF0000 size=5 face='黑体'>事故警告: JVM OOM</font>\n### 服务名: %s\n### 时间: %s\n### Dump文件: %s",podId,time.Now().Format("2006/01/02 15:04:05"),ossurl)
	msgSender(dingdingToken,alarmMsg)
}