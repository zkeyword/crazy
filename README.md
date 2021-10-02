CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

go build -ldflags "-s -w" main.go


// google authenticator
https://studygolang.com/articles/25748
https://www.jianshu.com/p/3c4832e3c6d5

// 逆向
https://zhuanlan.zhihu.com/p/26733683

// 打包静态文件
go-bindata

// 部署
nohup ./main &1>2 &

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" main.go

nohup 部署 日志 https://www.cnblogs.com/zhangmingcheng/p/11577967.html

// 热更新
go get github.com/pilu/fresh

在环境变量中配置了go bin的目录正常的话就可以正常执行，进入项目目录执行

    fresh

可以代替

    go run main.go
