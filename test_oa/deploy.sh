#交叉编译 $1表示编译结果的项目名称
GOOS=linux GOARCH=amd64 go build -o $1

#上传到指定路径下
scp -P 22 $1 tonnn@10.0.2.122:/home/tonnn/projects/$1

# 执行服务端上的脚本
ssh -tt tonnn@10.0.2.122 "cd /data/script/&&sudo ./deploy.sh $1"