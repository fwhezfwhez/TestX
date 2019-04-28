1. 交叉编译
`CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main`

2. 编写docker file
```Dockerfile
FROM centos:latest #使用了镜像大小体积只有5MB的alpine镜像

WORKDIR / #设置工作路径

ADD main / #把上文编译好的main文件添加到镜像里

EXPOSE 3000 #暴露容器内部端口

ENTRYPOINT ["./main"] #入口
```



3. 构建镜像
`docker build -t main:v1 .`  .为DockerFile所在目录

4. 运行容器
`docker run --rm -idt -p 3000:3000 main:v1`

其中，
-d: 守护模式 后台运行
-p: 是容器内部端口绑定到指定的主机端口
-P: 是容器内部端口随机映射到主机的高端口
-f : 让 docker logs 像使用 tail -f 一样来输出容器内部的标准输出
-l : 查询最后一次创建的容器
-t : 选项让Docker分配一个伪终端（pseudo-tty）并绑定到容器的标准输入上
-i : 则让容器的标准输入保持打开
-m:提交的描述信息
-a: 指定镜像作者
--rm: 停止容器后移除容器

拓展:
1. dockerfile暴露多个端口
EXPOSE 8000 8001 8002
2. 容器端口映射多个端口
docker run -p 8001:8001 -p 8002:8002
