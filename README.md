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
