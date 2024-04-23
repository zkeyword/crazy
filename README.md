CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go

go build -ldflags "-s -w" main.go


// google authenticator
https://studygolang.com/articles/25748
https://www.jianshu.com/p/3c4832e3c6d5

// 逆向
https://zhuanlan.zhihu.com/p/26733683

// 打包静态文件
安装 go install github.com/go-bindata/go-bindata/go-bindata@latest

命令 //go:generate go-bindata -fs -nocompress -nomemcopy -o=pkg/assets/assets.go -pkg=assets ./assets

// 部署  // 貌似不能自动重启
nohup ./main &1>2 & 

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" main.go

nohup 部署 日志 https://www.cnblogs.com/zhangmingcheng/p/11577967.html

// 热更新
go get github.com/pilu/fresh

在环境变量中配置了go bin的目录正常的话就可以正常执行，进入项目目录执行

    fresh

可以代替

    go run main.go

// 
go mod tidy的使用
引用项目需要的依赖增加到go.mod文件。
去掉go.mod文件中项目不需要的依赖。
