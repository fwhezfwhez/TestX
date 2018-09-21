GOOS=linux GOARCH=amd64 go build -o testoa
scp -P 22 testoa tonnn@10.0.2.122:/home/tonnn/projects/testoa 
ssh -tt tonnn@10.0.2.122 "sudo supervisorctl restart testoa&&sudo supervisorctl restart productsrv"