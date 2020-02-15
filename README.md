# ipScan

### ui fyne
https://fyne.io/develop/cross-compiling.html


```shell script
go mod download
go mod vendor
go get github.com/lucor/fyne-cross
fyne-cross --targets=linux/amd64,windows/amd64,darwin/amd64 package  
   
```

### docker镜像容器知识扩展
fyne-cross命令打包了对docker镜像的操作，下载，启动容器并打包等操作，了解对已有镜像进行相关明细操作，可以解决一些
场景异常情况。
交叉编译中，对golang.org/x等包访问不到，但是本地已经下载好了，就要把本地包上传到镜像中并重新打出对应版本的镜像。

```shell script
#启动并进入一个镜像的容器
docker run -t -i lucor/fyne-cross:develop /bin/bash

#修改镜像，创建个目录 /go 是该镜像的 GOPATH
mkdir -p /go/pkg_local

#不要exit 另开一个窗口看下当前容器id  例如为 2550ea51c7e6  就会把对应版本的镜像替换
#commit完成前 主窗口中不能exit退出
docker commit 2550ea51c7e6  lucor/fyne-cross:develop

#重新进入新镜像容器 把本地gopath虚拟目录挂载
docker run -t -i  -v /home/gopath/:/go/pkg_local  lucor/fyne-cross:develop /bin/bash
#挂载目录的包复制到容器中
cp -r /go/pkg_local/* /go/pkg/

#不要exit 另开一个窗口看下当前容器id  例如为 123  就会把对应版本的镜像替换，此时镜像就有本地一样的包了
docker commit 123  lucor/fyne-cross:develop

```
