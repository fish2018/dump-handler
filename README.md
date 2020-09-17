编译：
```
GOOS=linux go build
```

使用方法：
```
1、添加jvm参数"-XX:+HeapDumpOnOutOfMemoryError -XX:HeapDumpPath=/dumps/oom"
2、部署应用到k8s时在deployment配置挂载emptyDir的volume
3、在preStop中配置"./devops -k `echo $HOSTNAME` -e `echo $ENV`"
```

说明：
- PODID
k8s pod的id，可以通过`echo $HOSTNAME`获取
- ENV
部署环境，可以通过`echo $ENV`获取(提前通过deployment配置环境变量ENV到容器中)

podid命名规范：
示例 "opu-appserver"，以"-"为分隔符，取一个"opu"值为项目组，后面的"appserver"为app名

OOM Dump文件路径：
/dumps/oom

功能：
判断是否存在jvm dump文件"/dumps/oom"，如果存在就把/dumps/oom文件上传至oss，并发送钉钉告警，如果不存在则忽略。
oss文件分片，断点续传，根据podid、env自动判断上传到对应项目组的bucket
钉钉告警，附带oss链接地址，根据podid自动判断发送到对应项目组的告警群