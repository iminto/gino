# gino
gin框架学习使用笔记，使用gin+gorm

#安装运行
```
# 启用 Go Modules 功能(go 1.13后不需要)
$env:GO111MODULE="on" #Windows 
export GO111MODULE=on #linux
# 配置 GOPROXY 环境变量
$env:GOPROXY="https://goproxy.io" #Windows 
export GOPROXY=https://goproxy.io #linux
#修改config/config.go中数据库配置
#编译
go build
#运行
./gindemo
```