直接使用：

```shell
wget -c https://raw.githubusercontent.com/minroute/lnmp-checker/main/lnmp_checker
chmod +x lnmp_checker
lnmp_checker
```



自己编译： GOOS=linux GOARCH=amd64 go build -o lnmp_checker  main.go

