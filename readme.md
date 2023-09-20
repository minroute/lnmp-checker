直接使用：

```shell
# 下载
wget -c https://github.com/minroute/lnmp-checker/releases/download/v1/lnmp_checker

# 给执行权限
chmod +x lnmp_checker

# 检测
./lnmp_checker
```



自己编译： GOOS=linux GOARCH=amd64 go build -o lnmp_checker  main.go

