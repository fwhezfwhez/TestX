#交叉编译 $1表示编译结果的项目名称
GOOS=linux GOARCH=amd64 go build -o $1

#上传到指定路径下
scp -P 22 $1 tonnn@10.0.2.122:/home/tonnn/projects/$1

#远程登陆并依次执行命令
ssh -tt tonnn@10.0.2.122 "sudo touch /etc/supervisor/config.d/$1.ini&&
                          sudo chmod 777 /etc/supervisor/config.d/$1.ini&&
						  sudo echo '[program:$1]
directory = /home/tonnn/projects/$1
command = /home/tonnn/projects/$1/$1 -port=:8088 
autostart = true
autorestart = true
redirect_stderr = true
stdout_logfile = /data/log/$1.log'>/etc/supervisor/config.d/$1.ini&&
sudo touch /data/log/$1.log&&sudo chmod 777 /data/log/$1.log&&
sudo echo '[log for $1]' >/data/log/$1.log&&
sudo supervisorctl reload&&
sudo supervisorctl restart $1&&
sudo supervisorctl restart productsrv"