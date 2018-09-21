GOOS=linux GOARCH=amd64 go build -o $1
scp -P 22 $1 tonnn@10.0.2.122:/home/tonnn/projects/$1
ssh -tt tonnn@10.0.2.122 "sudo supervisorctl restart $1&&sudo supervisorctl restart productsrv"