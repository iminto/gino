# gino
gin框架学习使用笔记，使用gin+gorm

#安装运行
```
# 启用 Go Modules 功能
$env:GO111MODULE="on" #Windows 
export GO111MODULE=on #linux
# 配置 GOPROXY 环境变量
$env:GOPROXY="https://goproxy.io" #Windows 
export GOPROXY=https://goproxy.io #linux
#开启MYSQL
systemctl start mariadb
#端口映射包本地（如果有必要）
ssh -fN -L3306:localhost:3306 -p5555 baicai@148.70.1.1
#修改config/config.go中配置
#编译
go build
#运行
./gindemo
```